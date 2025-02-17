package oracle

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"sort"
	"time"

	"morph-l2/bindings/bindings"
	"morph-l2/node/derivation"
	"morph-l2/oracle/backoff"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmtypes "github.com/tendermint/tendermint/types"
)

func (o *Oracle) getBlockTimeAndNumber(isFinalized bool) (uint64, *big.Int, error) {
	var lastBlockNumber *big.Int
	if isFinalized {
		latestFinalized, err := o.rollup.LastFinalizedBatchIndex(nil)
		if err != nil {
			return 0, nil, err
		}
		batch, err := o.l2Client.GetRollupBatchByIndex(context.Background(), latestFinalized.Uint64())
		if err != nil {
			return 0, nil, err
		}
		if batch == nil {
			return 0, nil, fmt.Errorf("batch not found")
		}
		var batchData derivation.BatchInfo
		if err = batchData.ParseBatch(*batch); err != nil {
			return 0, nil, fmt.Errorf("parse batch error:%v", err)
		}
		lastBlockNumber = big.NewInt(int64(batchData.LastBlockNumber()))
	}

	header, err := o.l2Client.HeaderByNumber(o.ctx, lastBlockNumber)
	if err != nil {
		return 0, nil, err
	}
	return header.Time, header.Number, nil
}

func (o *Oracle) syncRewardEpoch() error {
	_, finalizedBlock, err := o.getBlockTimeAndNumber(o.isFinalized)
	if err != nil {
		return fmt.Errorf("get block time and number error:%v", err)
	}
	startRewardEpochIndex, err := o.record.NextRewardEpochIndex(nil)
	if err != nil {
		return err
	}
	startHeight, err := o.getNextHeight()
	if err != nil {
		return err
	}
	if startHeight.Cmp(finalizedBlock) > 0 {
		time.Sleep(30 * time.Second)
		return nil
	}
	recordRewardEpochInfo, err := o.getRewardEpochs(startRewardEpochIndex, startHeight)
	if err != nil {
		return err
	}
	chainId, err := o.l2Client.ChainID(o.ctx)
	if err != nil {
		return err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(o.privKey, chainId)
	if err != nil {
		return err
	}
	tx, err := o.record.RecordRewardEpochs(opts, []bindings.IRecordRewardEpochInfo{*recordRewardEpochInfo})
	if err != nil {
		return fmt.Errorf("record reward epochs error:%v", err)
	}
	log.Info("send record reward tx success", "txHash", tx.Hash().Hex(), "nonce", tx.Nonce())
	receipt, err := o.l2Client.TransactionReceipt(o.ctx, tx.Hash())
	if err != nil {
		return fmt.Errorf("receipt record reward epochs error:%v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("record reward epochs not success")
	}
	o.metrics.SetRewardEpoch(recordRewardEpochInfo.Index.Uint64())
	return nil
}

func (o *Oracle) getRewardEpochs(startRewardEpochIndex, startHeight *big.Int) (*bindings.IRecordRewardEpochInfo, error) {
	endTime, err := o.getEndTime(startHeight, startRewardEpochIndex)
	if err != nil {
		return nil, err
	}
	log.Info("new epoch fetching...", "startHeight", startHeight, "startRewardEpochIndex", startRewardEpochIndex, "endTime", endTime)
	height := startHeight
	sequencersBlockCount := make(map[common.Address]int64)
	for {
		_, finalizedBlock, err := o.getBlockTimeAndNumber(o.isFinalized)
		if err != nil {
			continue
		}
		if height.Cmp(finalizedBlock) > 0 {
			log.Info("finalized block small than syncing block,wait...", "finalizedBlock", finalizedBlock, "syncingBlock", height)
			time.Sleep(time.Second * 30)
			continue
		}
		tmHeader, err := o.L2HeaderByNumberWithRetry(height.Int64())
		if err != nil {
			return nil, fmt.Errorf("get l2 header error:%v", err)
		}
		if tmHeader.Time.Unix() > endTime.Int64() {
			break
		}
		log.Info("get new header", "headerNumber", tmHeader.Height, "headerTime", tmHeader.Time)
		sequencer, err := o.getSequencer(tmHeader.ProposerAddress, height)
		if err != nil {
			return nil, fmt.Errorf("get sequencer error:%v", err)
		}
		sequencersBlockCount[sequencer] += 1

		height = new(big.Int).Add(height, big.NewInt(1))
	}

	var sequencers []common.Address
	var seqBlockCounts, sequencerRatios, sequencerCommissions []*big.Int
	var blC int64
	for seq, count := range sequencersBlockCount {
		blC += count
		sequencers = append(sequencers, seq)
		seqBlockCounts = append(seqBlockCounts, big.NewInt(count))
	}
	blockCount := new(big.Int).Sub(height, startHeight)
	precision := big.NewInt(defaultPrecision)
	residue := big.NewInt(defaultPrecision)
	maxRatio := big.NewInt(0)
	var maxRatioIndex int
	for i := 0; i < len(sequencers); i++ {
		ratio := new(big.Int).Div(new(big.Int).Mul(seqBlockCounts[i], precision), blockCount)
		sequencerRatios = append(sequencerRatios, ratio)
		residue = new(big.Int).Sub(residue, ratio)
		if ratio.Cmp(maxRatio) > 0 {
			maxRatio = ratio
			maxRatioIndex = i
		}
		commission, err := o.getSequencerCommission(new(big.Int).Sub(startHeight, big.NewInt(1)), sequencers[i])
		if err != nil {
			return nil, fmt.Errorf("get sequencer commission error:%v", err)
		}
		sequencerCommissions = append(sequencerCommissions, commission)
	}
	sequencerRatios[maxRatioIndex] = new(big.Int).Add(sequencerRatios[maxRatioIndex], residue)
	rewardEpochInfo := bindings.IRecordRewardEpochInfo{
		Index:                startRewardEpochIndex,
		BlockCount:           blockCount,
		Sequencers:           sequencers,
		SequencerBlocks:      seqBlockCounts,
		SequencerRatios:      sequencerRatios,
		SequencerCommissions: sequencerCommissions,
	}
	return &rewardEpochInfo, nil
}

func (o *Oracle) getSequencerCommission(blockNumber *big.Int, address common.Address) (*big.Int, error) {
	if blockNumber.Uint64() < o.cfg.StartBlock {
		return big.NewInt(0), nil
	}
	return o.l2Staking.Commissions(&bind.CallOpts{
		BlockNumber: blockNumber,
	}, address)
}

// L2HeaderByNumberWithRetry retries getting headers.
func (o *Oracle) L2HeaderByNumberWithRetry(height int64) (*tmtypes.Header, error) {
	var res *tmtypes.Header
	err := backoff.DoCtx(o.ctx, 3, backoff.Exponential(), func() error {
		var err error
		headerResp, err := o.TmClient.Header(context.Background(), &height)
		if err != nil {
			return err
		}
		res = headerResp.Header
		return nil
	})
	return res, err
}

func (o *Oracle) getSequencer(proposerAddress tmtypes.Address, blockNumber *big.Int) (common.Address, error) {
	stakers, err := o.l2Staking.GetStakers(&bind.CallOpts{
		BlockNumber: new(big.Int).Sub(blockNumber, big.NewInt(1)),
	})
	if err != nil {
		return common.Address{}, err
	}
	for _, staker := range stakers {
		if bytes.Equal(proposerAddress, ed25519.PubKey(staker.TmKey[:]).Address().Bytes()) {
			return staker.Addr, nil
		}
	}
	return common.Address{}, fmt.Errorf("sequencer not found")
}

func (o *Oracle) getNextHeight() (*big.Int, error) {
	latest, err := o.record.LatestRewardEpochBlock(nil)
	if err != nil {
		return latest, err
	}
	return new(big.Int).Add(latest, big.NewInt(1)), nil
}

func (o *Oracle) getEndTime(blockNumber *big.Int, nextRewardEpochIndex *big.Int) (*big.Int, error) {
	startTime, err := o.l2Staking.RewardStartTime(&bind.CallOpts{
		BlockNumber: blockNumber,
	})
	if err != nil {
		return nil, err
	}
	internal := new(big.Int).Mul(nextRewardEpochIndex, big.NewInt(int64(o.rewardEpoch)))
	epochStart := new(big.Int).Add(startTime, internal)
	epochEnd := new(big.Int).Add(epochStart, big.NewInt(int64(o.rewardEpoch)))
	return epochEnd, nil
}

func (o *Oracle) findStartBlock(start, end uint64, timeStamp int64) (int64, error) {
	headerStart, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(start)))
	if err != nil {
		return 0, err
	}
	headerEnd, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(end)))
	if err != nil {
		return 0, err
	}
	if end < start {
		return 0, fmt.Errorf("invalid start or end,start:%v,end:%v", start, end)
	}
	if int64(headerStart.Time) > timeStamp || int64(headerEnd.Time) < timeStamp {
		return 0, fmt.Errorf("this timestamp is not within the given block range")
	}

	s := sort.Search(int(end)+1-int(start), func(i int) bool {
		header, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(i)+int64(start)))
		if err != nil {
			log.Error("get header by number failed", "error", err)
			return false
		}
		return int64(header.Time) >= timeStamp
	})
	if s == int(end)+1-int(start) {
		log.Error("start block not found")
	}
	target, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(start)+int64(s)))
	if err != nil {
		return 0, err
	}
	preHeader, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(start)+int64(s)-1))
	if err != nil {
		return 0, err
	}
	if !(int64(preHeader.Time) < timeStamp && int64(target.Time) >= timeStamp) {
		return 0, fmt.Errorf("invalid start block")
	}
	log.Info("find start block success", "preHeader_time", preHeader.Time, "timestamp", timeStamp, "target_time", target.Time)
	return int64(start) + int64(s), nil
}

func (o *Oracle) setStartBlock() {
	start := o.cfg.StartBlock
	for {
		header, err := o.l2Client.HeaderByNumber(o.ctx, nil)
		if err != nil {
			panic(err)
		}
		rewardStarted, err := o.l2Staking.RewardStarted(&bind.CallOpts{
			BlockNumber: header.Number,
		})
		if err != nil {
			panic(err)
		}
		if rewardStarted {
			log.Info("reward start")
			break
		}
		start = header.Number.Uint64()
		log.Info("wait reward start...")
		time.Sleep(defaultSleepTime)
		continue
	}

	for {
		header, err := o.l2Client.HeaderByNumber(o.ctx, nil)
		if err != nil {
			log.Error("query header by number failed", "error", err)
		}
		startTime, err := o.l2Staking.RewardStartTime(&bind.CallOpts{
			BlockNumber: header.Number,
		})
		if err != nil {
			log.Error("query reward start time failed", "error", err)
		}
		latestRewardEpochBlock, err := o.record.LatestRewardEpochBlock(nil)
		if err != nil {
			log.Error("query latest reward epoch block failed", "error", err)
		}
		if latestRewardEpochBlock.Uint64() != 0 {
			break
		}
		if header.Time < startTime.Uint64() {
			start = header.Number.Uint64()
			time.Sleep(defaultSleepTime)
			continue
		}
		log.Info("start find start block", "start_block", start, "end_block", header.Number.Uint64())
		startBlock, err := o.findStartBlock(start, header.Number.Uint64(), startTime.Int64())
		if err != nil {
			log.Error("find start block failed", "error", err)
			time.Sleep(defaultSleepTime)
			continue
		}

		chainID, err := o.l2Client.ChainID(o.ctx)
		if err != nil {
			log.Error("get chain id failed", "error", err)
			time.Sleep(defaultSleepTime)
			continue
		}
		opts, err := bind.NewKeyedTransactorWithChainID(o.privKey, chainID)
		if err != nil {
			log.Error(fmt.Sprintf("new opts error:%v", err))
		}
		nonce, err := o.l2Client.NonceAt(context.Background(), crypto.PubkeyToAddress(o.privKey.PublicKey), nil)
		if err != nil {
			return
		}
		opts.NoSend = true
		opts.Nonce = big.NewInt(int64(nonce))
		tx, err := o.record.SetLatestRewardEpochBlock(opts, big.NewInt(startBlock))
		if err != nil {
			log.Error("set latestReward epoch block failed", "error", err)
			time.Sleep(defaultSleepTime)
			continue
		}
		signedTx, err := opts.Signer(opts.From, tx)
		if err != nil {
			return
		}
		err = o.l2Client.SendTransaction(o.ctx, signedTx)
		if err != nil {
			log.Error("send transaction failed,retry...", "error", err)
			time.Sleep(defaultSleepTime)
			continue
		}
		receipt, err := o.waitReceiptWithCtx(o.ctx, tx.Hash())
		if err != nil {
			log.Error("TransactionReceipt failed,retry...", "error", err)
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			log.Error("set stark block failed")
			continue
		}
		log.Info("set start block success")
	}
}
