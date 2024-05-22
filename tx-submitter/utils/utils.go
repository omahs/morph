package utils

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/big"
	"reflect"
	"time"

	"github.com/morph-l2/bindings/bindings"

	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/log"
	"github.com/scroll-tech/go-ethereum/rpc"
)

// Loop Run the f func periodically.
func Loop(ctx context.Context, period time.Duration, f func()) {
	tick := time.NewTicker(period)
	defer tick.Stop()
	for ; ; <-tick.C {
		select {
		case <-ctx.Done():
			return
		default:
			f()
		}
	}
}

func ParseParentBatchIndex(calldata []byte) uint64 {
	///   * Field                   Bytes       Type        Index   Comments
	///   * version                 1           uint8       0       The batch version
	///   * batchIndex              8           uint64      1       The index of the batch
	///   * l1MessagePopped         8           uint64      9       Number of L1 messages popped in the batch

	// get calldata
	// calldata := "4db4d96400000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000116000000000000000000000000000000000000000000000000000000000000011800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001a00000000000000000000000000000000000000000000000000000000000000ec015b28c3cfef1e0371da3b8504a7affae87c2683881ae1b1e0276c0ddca3f27071234cad004e504650b46e4ddf52a3f335eb20f84977d9277856680e65efecc803248d1e8440ea207ae6f89011e4c9b6e10196cf35160b66b17f75f59c5619cb70000000000000000000000000000000000000000000000000000000000000ee000000000000000000000000000000000000000000000000000000000000000790000000000000012b90000000000000000000000000001ed8d919a726f9a2575a7ca11f61376240a1201e11dc586105b6f66008cee08bf6856012e35b6c0f787534ebfdbc42f3230421781c5bb72cd7e40c7270d9e8f639ea1f922bdb7c0056a3b155a5f80b50ac9f12de6341c63a3a491640b69e15ad76af200000000000000000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000002e000000000000000000000000000000000000000000000000000000000000003c000000000000000000000000000000000000000000000000000000000000005200000000000000000000000000000000000000000000000000000000000000680000000000000000000000000000000000000000000000000000000000000076000000000000000000000000000000000000000000000000000000000000008c00000000000000000000000000000000000000000000000000000000000000a600000000000000000000000000000000000000000000000000000000000000b400000000000000000000000000000000000000000000000000000000000000ca0000000000000000000000000000000000000000000000000000000000000016906000000000007d04900000000664585db0000000000000000000000000000000000000000000000000000000000000000000000000098968000030000000000000007d04a00000000664585dd0000000000000000000000000000000000000000000000000000000000000000000000000098968000020000000000000007d04b00000000664585de0000000000000000000000000000000000000000000000000000000000000000000000000098968000010000000000000007d04c00000000664585e00000000000000000000000000000000000000000000000000000000000000000000000000098968000030000000000000007d04d00000000664585e10000000000000000000000000000000000000000000000000000000000000000000000000098968000040000000000000007d04e00000000664585e30000000000000000000000000000000000000000000000000000000000000000000000000098968000040000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b503000000000007d04f00000000664585e50000000000000000000000000000000000000000000000000000000000000000000000000098968000080000000000000007d05000000000664585e700000000000000000000000000000000000000000000000000000000000000000000000000989680000a0000000000000007d05100000000664585e900000000000000000000000000000000000000000000000000000000000000000000000000989680000500000000000000000000000000000000000000000000000000000000000000000000000000000000000000012d05000000000007d05200000000664585ea0000000000000000000000000000000000000000000000000000000000000000000000000098968000050000000000000007d05300000000664585ec0000000000000000000000000000000000000000000000000000000000000000000000000098968000040000000000000007d05400000000664585ee0000000000000000000000000000000000000000000000000000000000000000000000000098968000030000000000000007d05500000000664585ef0000000000000000000000000000000000000000000000000000000000000000000000000098968000020000000000000007d05600000000664585f1000000000000000000000000000000000000000000000000000000000000000000000000009896800005000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000012d05000000000007d05700000000664585f30000000000000000000000000000000000000000000000000000000000000000000000000098968000020000000000000007d05800000000664585f40000000000000000000000000000000000000000000000000000000000000000000000000098968000050000000000000007d05900000000664585f60000000000000000000000000000000000000000000000000000000000000000000000000098968000030000000000000007d05a00000000664585f80000000000000000000000000000000000000000000000000000000000000000000000000098968000040000000000000007d05b00000000664585f900000000000000000000000000000000000000000000000000000000000000000000000000989680000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b503000000000007d05c00000000664585fb0000000000000000000000000000000000000000000000000000000000000000000000000098968000030000000000000007d05d00000000664585fd0000000000000000000000000000000000000000000000000000000000000000000000000098968000030000000000000007d05e00000000664585fe00000000000000000000000000000000000000000000000000000000000000000000000000989680000c00000000000000000000000000000000000000000000000000000000000000000000000000000000000000012d05000000000007d05f00000000664586000000000000000000000000000000000000000000000000000000000000000000000000000098968000080000000000000007d06000000000664586020000000000000000000000000000000000000000000000000000000000000000000000000098968000050000000000000007d06100000000664586040000000000000000000000000000000000000000000000000000000000000000000000000098968000020000000000000007d06200000000664586060000000000000000000000000000000000000000000000000000000000000000000000000098968000030000000000000007d0630000000066458608000000000000000000000000000000000000000000000000000000000000000000000000009896800008000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000016906000000000007d064000000006645860a0000000000000000000000000000000000000000000000000000000000000000000000000098968000030000000000000007d065000000006645860b0000000000000000000000000000000000000000000000000000000000000000000000000098968000020000000000000007d066000000006645860d0000000000000000000000000000000000000000000000000000000000000000000000000098968000050000000000000007d067000000006645860e0000000000000000000000000000000000000000000000000000000000000000000000000098968000020000000000000007d06800000000664586100000000000000000000000000000000000000000000000000000000000000000000000000098968000050000000000000007d06900000000664586110000000000000000000000000000000000000000000000000000000000000000000000000098968000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000b503000000000007d06a00000000664586130000000000000000000000000000000000000000000000000000000000000000000000000098968000050000000000000007d06b00000000664586150000000000000000000000000000000000000000000000000000000000000000000000000098968000070000000000000007d06c000000006645861700000000000000000000000000000000000000000000000000000000000000000000000000989680000600000000000000000000000000000000000000000000000000000000000000000000000000000000000000012d05000000000007d06d00000000664586190000000000000000000000000000000000000000000000000000000000000000000000000098968000040000000000000007d06e000000006645861b0000000000000000000000000000000000000000000000000000000000000000000000000098968000070000000000000007d06f000000006645861c0000000000000000000000000000000000000000000000000000000000000000000000000098968000020000000000000007d070000000006645861e0000000000000000000000000000000000000000000000000000000000000000000000000098968000040000000000000007d0710000000066458620000000000000000000000000000000000000000000000000000000000000000000000000009896800003000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003d01000000000007d0720000000066458621000000000000000000000000000000000000000000000000000000000000000000000000009896800003000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000003000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000001600000000000000000000000000000000000000000000000000000000000000007000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000300000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000005000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000019b46417d979494d5f16a044f1b9fa11ccd54726125a86cc4daa1e0651e735d52dc6ae1d31287cb54d14dc1cfb45813d00000000000000000000000000000000132da877af6f7854f6c6c3502a8a3ba454bea34f26a16d42269667375d6404eb7423c3b4629bff01968e993b12461ade000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000019b46417d979494d5f16a044f1b9fa11ccd54726125a86cc4daa1e0651e735d52dc6ae1d31287cb54d14dc1cfb45813d00000000000000000000000000000000132da877af6f7854f6c6c3502a8a3ba454bea34f26a16d42269667375d6404eb7423c3b4629bff01968e993b12461ade"
	// txpool:=NewTxPool()
	// bs, _ := hex.DecodeString(calldata)

	abi, _ := bindings.RollupMetaData.GetAbi()
	// require.NoError(t, err)
	// ans := make(map[string]interface{})
	parms, _ := abi.Methods["commitBatch"].Inputs.UnpackValues(calldata[4:])
	// ans, err := abi.Unpack("commitBatch", bs[4:])

	v := reflect.ValueOf(parms[0])
	pbh := v.FieldByName("ParentBatchHeader")
	batchIndex := binary.BigEndian.Uint64(pbh.Bytes()[1:9])
	return batchIndex
}

// ParseL1Mempool parses the L1 mempool and returns the transactions.
func ParseL1Mempool(rpc *rpc.Client, addr common.Address) ([]types.Transaction, uint64, error) {

	var result map[string]map[string]types.Transaction
	err := rpc.Call(&result, "txpool_contentFrom", addr)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get txpool content: %v", err)
	}

	latestNonce := uint64(0)

	var txs []types.Transaction

	// get pending txs
	if pendingTxs, ok := result["pending"]; ok {
		for _, tx := range pendingTxs {
			txs = append(txs, tx)
			if tx.Nonce() > latestNonce {
				latestNonce = tx.Nonce()
			}
		}
	}

	return txs, latestNonce, nil

}

func ParseMempoolLatestBatchIndex(id []byte, txs []types.Transaction) uint64 {

	var res uint64
	for _, tx := range txs {
		if bytes.Equal(tx.Data()[:4], id) {
			pindex := ParseParentBatchIndex(tx.Data())
			if pindex > res {
				res = pindex
			}
		}
	}

	return res + 1

}

func ParseBusinessInfo(tx types.Transaction, a *abi.ABI) []interface{} {
	// var method string
	// var batchIndex uint64
	// var finalizedIndex uint64
	var res []interface{}
	if len(tx.Data()) > 0 {
		id := tx.Data()[:4]
		if bytes.Equal(id, a.Methods["commitBatch"].ID) {
			method := "commitBatch"
			batchIndex := ParseParentBatchIndex(tx.Data()) + 1
			res = append(res,
				"method", method,
				"batchIndex", batchIndex,
			)
		} else if bytes.Equal(id, a.Methods["finalizeBatch"].ID) {
			method := "finalizeBatch"
			fIndex, err := a.Methods["finalizeBatch"].Inputs.Unpack(tx.Data()[4:])
			if err != nil {
				log.Error("unpack finalizeBatch error", "err", err)
			}
			finalizedIndex := fIndex[0].(*big.Int).Uint64()
			res = append(res,
				"method", method,
				"finalizedIndex", finalizedIndex,
			)

		}

	} else {
		return []interface{}{}
	}
	return res
}

func ParseMethod(tx types.Transaction, a *abi.ABI) string {
	if tx.Data() == nil || len(tx.Data()) < 4 {
		return ""
	}
	id := tx.Data()[:4]
	if bytes.Equal(id, a.Methods["commitBatch"].ID) {
		return "commitBatch"
	} else if bytes.Equal(id, a.Methods["finalizeBatch"].ID) {
		return "finalizeBatch"
	} else {
		return ""
	}
}
