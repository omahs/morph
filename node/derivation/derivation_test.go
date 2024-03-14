package derivation

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/node/types"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind/backends"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/core"
	"github.com/scroll-tech/go-ethereum/core/rawdb"
	"github.com/scroll-tech/go-ethereum/eth"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"github.com/scroll-tech/go-ethereum/ethclient/authclient"
	"github.com/scroll-tech/go-ethereum/ethdb"
	"github.com/scroll-tech/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

func TestCompareBlock(t *testing.T) {
	eClient, err := ethclient.Dial("http://localhost:7545")
	require.NoError(t, err)
	l2Client, err := ethclient.Dial("http://localhost:8545")
	blockNumber, err := eClient.BlockNumber(context.Background())
	require.NoError(t, err)
	for i := 0; i < int(blockNumber); i++ {
		block, err := l2Client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		require.NoError(t, err)
		dBlock, err := eClient.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		require.True(t, reflect.DeepEqual(block.Header(), dBlock.Header()))
	}
}

func testNewDerivationClient(t *testing.T) *Derivation {
	ctx := context.Background()
	l1Client, err := ethclient.Dial("http://localhost:9545")
	addr := common.HexToAddress("0x6900000000000000000000000000000000000010")
	require.NoError(t, err)
	var secret [32]byte
	jwtSecret := common.FromHex(strings.TrimSpace("688f5d737bad920bdfb2fc2f488d6b6209eebda1dae949a8de91398d932c517a"))
	require.True(t, len(jwtSecret) == 32)
	copy(secret[:], jwtSecret)
	aClient, err := authclient.DialContext(context.Background(), "http://localhost:7551", secret)
	require.NoError(t, err)
	eClient, err := ethclient.Dial("http://localhost:7545")
	require.NoError(t, err)
	logger := tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout))
	baseHttp := NewBasicHTTPClient("http://localhost:3500", logger)
	l1BeaconClient := NewL1BeaconClient(baseHttp)
	d := Derivation{
		ctx:                   ctx,
		l1Client:              l1Client,
		RollupContractAddress: addr,
		confirmations:         rpc.BlockNumber(5),
		l2Client:              types.NewRetryableClient(aClient, eClient, tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout))),
		validator:             nil,
		latestDerivation:      9,
		fetchBlockRange:       100,
		pollInterval:          1,
		l1BeaconClient:        l1BeaconClient,
	}
	return &d
}

func TestFetchRollupData(t *testing.T) {
	d := testNewDerivationClient(t)
	_, err := d.fetchRollupDataByTxHash(common.HexToHash("0xf2eaf0c8c121751a544046bc5358db521dc0b8dce660c674ad64c8f1a2a12c1b"), 390)
	require.NoError(t, err)
}

func newSimulatedBackend(key *ecdsa.PrivateKey) (*backends.SimulatedBackend, ethdb.Database) {
	var gasLimit uint64 = 9_000_000
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	genAlloc := make(core.GenesisAlloc)
	genAlloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(9223372036854775807)}
	db := rawdb.NewMemoryDatabase()
	sim := backends.NewSimulatedBackendWithDatabase(db, genAlloc, gasLimit)
	return sim, db
}

func TestDecodeBatch(t *testing.T) {
	abi, err := bindings.RollupMetaData.GetAbi()
	require.NoError(t, err)
	hexData := "0x92f65af30000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008e0000000000000000000000000000000000000000000000000000000000000098000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000018000000000000000000000000000000000000000000000000000000000000006a02b8abfcc1aa740eb6330093a5ef9c1d09c39672c390926de6db9089554b620322b8abfcc1aa740eb6330093a5ef9c1d09c39672c390926de6db9089554b6203227ae5ba08d7291c96c8cbddcc148bf48a6d68c7974b94356f53754ef6171d75700000000000000000000000000000000000000000000000000000000000006c00000000000000000000000000000000000000000000000000000000000000059000000000000000127000000000000000000000000000000006a5279896e0503ca805ecdb9208afa1ae78e214ceae183cf59fba3a8788e4598b0cd5e7bf4073160681199a5f107e15ec96c350840ebedb522a79c73615037b9000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004b11400000000000015520000000065e43bfb000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015530000000065e43c02000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015540000000065e43c08000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015550000000065e43c0f000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015560000000065e43c15000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015570000000065e43c1c000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015580000000065e43c22000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015590000000065e43c280000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000155a0000000065e43c2f0000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000155b0000000065e43c350000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000155c0000000065e43c3c0000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000155d0000000065e43c420000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000155e0000000065e43c490000000000000000000000000000000000000000000000000000000000000000000000000098968000000000000000000000155f0000000065e43c4f000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015600000000065e43c56000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015610000000065e43c5c000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015620000000065e43c62000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015630000000065e43c69000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015640000000065e43c6f000000000000000000000000000000000000000000000000000000000000000000000000009896800000000000000000000015650000000065e43c7600000000000000000000000000000000000000000000000000000000000000000000000000989680000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000600000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000001656c61193a9eca0201811fc18e071ee949910f696d7464fb3ed32ce9e9f022967f03962fb34c8533f0d01c13e808f4c000000000000000000000000000000000a066efeb662c0804b6c35a0be989a5fbd6a229261b12a176195772febca89e806e3fa2bebe19f95c7491af5afc7334e000000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000030000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000001656c61193a9eca0201811fc18e071ee949910f696d7464fb3ed32ce9e9f022967f03962fb34c8533f0d01c13e808f4c000000000000000000000000000000000a066efeb662c0804b6c35a0be989a5fbd6a229261b12a176195772febca89e806e3fa2bebe19f95c7491af5afc7334e"
	txData, err := hexutil.Decode(hexData)
	require.NoError(t, err)
	require.NoError(t, err)
	args, err := abi.Methods["commitBatch"].Inputs.Unpack(txData[4:])
	rollupBatchData := args[0].(struct {
		Version                uint8     "json:\"version\""
		ParentBatchHeader      []uint8   "json:\"parentBatchHeader\""
		Chunks                 [][]uint8 "json:\"chunks\""
		SkippedL1MessageBitmap []uint8   "json:\"skippedL1MessageBitmap\""
		PrevStateRoot          [32]uint8 "json:\"prevStateRoot\""
		PostStateRoot          [32]uint8 "json:\"postStateRoot\""
		WithdrawalRoot         [32]uint8 "json:\"withdrawalRoot\""
		Signature              struct {
			Version   *big.Int   "json:\"version\""
			Signers   []*big.Int "json:\"signers\""
			Signature []uint8    "json:\"signature\""
		} "json:\"signature\""
	})

	var chunks []hexutil.Bytes
	for _, chunk := range rollupBatchData.Chunks {
		chunks = append(chunks, chunk)
	}
	batch := eth.RPCRollupBatch{
		Version:                uint(rollupBatchData.Version),
		ParentBatchHeader:      rollupBatchData.ParentBatchHeader,
		Chunks:                 chunks,
		SkippedL1MessageBitmap: rollupBatchData.SkippedL1MessageBitmap,
		PrevStateRoot:          common.BytesToHash(rollupBatchData.PrevStateRoot[:]),
		PostStateRoot:          common.BytesToHash(rollupBatchData.PostStateRoot[:]),
		WithdrawRoot:           common.BytesToHash(rollupBatchData.WithdrawalRoot[:]),
	}
	batchInfo := new(BatchInfo)
	err = batchInfo.ParseBatch(batch)
	require.NoError(t, err)
}
