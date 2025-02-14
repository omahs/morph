// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// L2TxFeeVaultMetaData contains all meta data concerning the L2TxFeeVault contract.
var L2TxFeeVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minWithdrawalAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldMessenger\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newMessenger\",\"type\":\"address\"}],\"name\":\"UpdateMessenger\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMinWithdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMinWithdrawAmount\",\"type\":\"uint256\"}],\"name\":\"UpdateMinWithdrawAmount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"UpdateRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"messenger\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minWithdrawAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newMessenger\",\"type\":\"address\"}],\"name\":\"updateMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMinWithdrawAmount\",\"type\":\"uint256\"}],\"name\":\"updateMinWithdrawAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRecipient\",\"type\":\"address\"}],\"name\":\"updateRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50604051610c8f380380610c8f83398101604081905261002e916100ca565b61003783610060565b600155600380546001600160a01b0319166001600160a01b039290921691909117905550610103565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b80516001600160a01b03811681146100c5575f80fd5b919050565b5f805f606084860312156100dc575f80fd5b6100e5846100af565b92506100f3602085016100af565b9150604084015190509250925092565b610b7f806101105f395ff3fe6080604052600436106100c6575f3560e01c806384411d6511610071578063f2fde38b1161004c578063f2fde38b1461021e578063feec756c1461023d578063ff4f35461461025c575f80fd5b806384411d65146101bf5780638da5cb5b146101d45780639e7adc79146101ff575f80fd5b8063457e1a49116100a1578063457e1a491461015c57806366d003ac1461017f578063715018a6146101ab575f80fd5b80632e1a7d4d146100d15780633cb747bf146100f25780633ccfd60b14610148575f80fd5b366100cd57005b5f80fd5b3480156100dc575f80fd5b506100f06100eb366004610a8b565b61027b565b005b3480156100fd575f80fd5b5060025461011e9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b348015610153575f80fd5b506100f0610548565b348015610167575f80fd5b5061017160015481565b60405190815260200161013f565b34801561018a575f80fd5b5060035461011e9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156101b6575f80fd5b506100f06105d5565b3480156101ca575f80fd5b5061017160045481565b3480156101df575f80fd5b505f5461011e9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561020a575f80fd5b506100f0610219366004610aa2565b610660565b348015610229575f80fd5b506100f0610238366004610aa2565b610756565b348015610248575f80fd5b506100f0610257366004610aa2565b61085c565b348015610267575f80fd5b506100f0610276366004610a8b565b610952565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610300576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064015b60405180910390fd5b6001548110156103b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a4016102f7565b4780821115610449576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f4665655661756c743a20696e73756666696369656e742062616c616e6365207460448201527f6f2077697468647261770000000000000000000000000000000000000000000060648201526084016102f7565b60048054830190556003546040805184815273ffffffffffffffffffffffffffffffffffffffff90921660208301523382820152517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1600254600354604080516020810182525f80825291517fb2267a7b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9485169463b2267a7b948894610516949190921692859290600401610adc565b5f604051808303818588803b15801561052d575f80fd5b505af115801561053f573d5f803e3d5ffd5b50505050505050565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146105c8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016102f7565b476105d28161027b565b50565b5f5473ffffffffffffffffffffffffffffffffffffffff163314610655576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016102f7565b61065e5f610a17565b565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146106e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016102f7565b6002805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f1c928c417a10a21c3cddad148c5dba5d710e4b1442d6d8a36de345935ad84612905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146107d6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016102f7565b73ffffffffffffffffffffffffffffffffffffffff8116610853576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f6e6577206f776e657220697320746865207a65726f206164647265737300000060448201526064016102f7565b6105d281610a17565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146108dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016102f7565b6003805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f7e1e96961a397c8aa26162fe259cc837afc95e33aad4945ddc61c18dabb7a6ad905f90a35050565b5f5473ffffffffffffffffffffffffffffffffffffffff1633146109d2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f63616c6c6572206973206e6f7420746865206f776e657200000000000000000060448201526064016102f7565b600180549082905560408051828152602081018490527f0d3c80219fe57713b9f9c83d1e51426792d0c14d8e330e65b102571816140965910160405180910390a15050565b5f805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f60208284031215610a9b575f80fd5b5035919050565b5f60208284031215610ab2575f80fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610ad5575f80fd5b9392505050565b73ffffffffffffffffffffffffffffffffffffffff851681525f60208560208401526080604084015284518060808501525f5b81811015610b2b5786810183015185820160a001528201610b0f565b505f60a0828601015260a07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050508260608301529594505050505056fea164736f6c6343000818000a",
}

// L2TxFeeVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use L2TxFeeVaultMetaData.ABI instead.
var L2TxFeeVaultABI = L2TxFeeVaultMetaData.ABI

// L2TxFeeVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2TxFeeVaultMetaData.Bin instead.
var L2TxFeeVaultBin = L2TxFeeVaultMetaData.Bin

// DeployL2TxFeeVault deploys a new Ethereum contract, binding an instance of L2TxFeeVault to it.
func DeployL2TxFeeVault(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _recipient common.Address, _minWithdrawalAmount *big.Int) (common.Address, *types.Transaction, *L2TxFeeVault, error) {
	parsed, err := L2TxFeeVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2TxFeeVaultBin), backend, _owner, _recipient, _minWithdrawalAmount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2TxFeeVault{L2TxFeeVaultCaller: L2TxFeeVaultCaller{contract: contract}, L2TxFeeVaultTransactor: L2TxFeeVaultTransactor{contract: contract}, L2TxFeeVaultFilterer: L2TxFeeVaultFilterer{contract: contract}}, nil
}

// L2TxFeeVault is an auto generated Go binding around an Ethereum contract.
type L2TxFeeVault struct {
	L2TxFeeVaultCaller     // Read-only binding to the contract
	L2TxFeeVaultTransactor // Write-only binding to the contract
	L2TxFeeVaultFilterer   // Log filterer for contract events
}

// L2TxFeeVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type L2TxFeeVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TxFeeVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type L2TxFeeVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TxFeeVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2TxFeeVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2TxFeeVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2TxFeeVaultSession struct {
	Contract     *L2TxFeeVault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2TxFeeVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2TxFeeVaultCallerSession struct {
	Contract *L2TxFeeVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// L2TxFeeVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2TxFeeVaultTransactorSession struct {
	Contract     *L2TxFeeVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// L2TxFeeVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type L2TxFeeVaultRaw struct {
	Contract *L2TxFeeVault // Generic contract binding to access the raw methods on
}

// L2TxFeeVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2TxFeeVaultCallerRaw struct {
	Contract *L2TxFeeVaultCaller // Generic read-only contract binding to access the raw methods on
}

// L2TxFeeVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2TxFeeVaultTransactorRaw struct {
	Contract *L2TxFeeVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewL2TxFeeVault creates a new instance of L2TxFeeVault, bound to a specific deployed contract.
func NewL2TxFeeVault(address common.Address, backend bind.ContractBackend) (*L2TxFeeVault, error) {
	contract, err := bindL2TxFeeVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVault{L2TxFeeVaultCaller: L2TxFeeVaultCaller{contract: contract}, L2TxFeeVaultTransactor: L2TxFeeVaultTransactor{contract: contract}, L2TxFeeVaultFilterer: L2TxFeeVaultFilterer{contract: contract}}, nil
}

// NewL2TxFeeVaultCaller creates a new read-only instance of L2TxFeeVault, bound to a specific deployed contract.
func NewL2TxFeeVaultCaller(address common.Address, caller bind.ContractCaller) (*L2TxFeeVaultCaller, error) {
	contract, err := bindL2TxFeeVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultCaller{contract: contract}, nil
}

// NewL2TxFeeVaultTransactor creates a new write-only instance of L2TxFeeVault, bound to a specific deployed contract.
func NewL2TxFeeVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*L2TxFeeVaultTransactor, error) {
	contract, err := bindL2TxFeeVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultTransactor{contract: contract}, nil
}

// NewL2TxFeeVaultFilterer creates a new log filterer instance of L2TxFeeVault, bound to a specific deployed contract.
func NewL2TxFeeVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*L2TxFeeVaultFilterer, error) {
	contract, err := bindL2TxFeeVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultFilterer{contract: contract}, nil
}

// bindL2TxFeeVault binds a generic wrapper to an already deployed contract.
func bindL2TxFeeVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2TxFeeVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2TxFeeVault *L2TxFeeVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2TxFeeVault.Contract.L2TxFeeVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2TxFeeVault *L2TxFeeVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.L2TxFeeVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2TxFeeVault *L2TxFeeVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.L2TxFeeVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2TxFeeVault *L2TxFeeVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2TxFeeVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2TxFeeVault *L2TxFeeVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2TxFeeVault *L2TxFeeVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.contract.Transact(opts, method, params...)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCaller) Messenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "messenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultSession) Messenger() (common.Address, error) {
	return _L2TxFeeVault.Contract.Messenger(&_L2TxFeeVault.CallOpts)
}

// Messenger is a free data retrieval call binding the contract method 0x3cb747bf.
//
// Solidity: function messenger() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) Messenger() (common.Address, error) {
	return _L2TxFeeVault.Contract.Messenger(&_L2TxFeeVault.CallOpts)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultCaller) MinWithdrawAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "minWithdrawAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultSession) MinWithdrawAmount() (*big.Int, error) {
	return _L2TxFeeVault.Contract.MinWithdrawAmount(&_L2TxFeeVault.CallOpts)
}

// MinWithdrawAmount is a free data retrieval call binding the contract method 0x457e1a49.
//
// Solidity: function minWithdrawAmount() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) MinWithdrawAmount() (*big.Int, error) {
	return _L2TxFeeVault.Contract.MinWithdrawAmount(&_L2TxFeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultSession) Owner() (common.Address, error) {
	return _L2TxFeeVault.Contract.Owner(&_L2TxFeeVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) Owner() (common.Address, error) {
	return _L2TxFeeVault.Contract.Owner(&_L2TxFeeVault.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCaller) Recipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "recipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultSession) Recipient() (common.Address, error) {
	return _L2TxFeeVault.Contract.Recipient(&_L2TxFeeVault.CallOpts)
}

// Recipient is a free data retrieval call binding the contract method 0x66d003ac.
//
// Solidity: function recipient() view returns(address)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) Recipient() (common.Address, error) {
	return _L2TxFeeVault.Contract.Recipient(&_L2TxFeeVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultCaller) TotalProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2TxFeeVault.contract.Call(opts, &out, "totalProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultSession) TotalProcessed() (*big.Int, error) {
	return _L2TxFeeVault.Contract.TotalProcessed(&_L2TxFeeVault.CallOpts)
}

// TotalProcessed is a free data retrieval call binding the contract method 0x84411d65.
//
// Solidity: function totalProcessed() view returns(uint256)
func (_L2TxFeeVault *L2TxFeeVaultCallerSession) TotalProcessed() (*big.Int, error) {
	return _L2TxFeeVault.Contract.TotalProcessed(&_L2TxFeeVault.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.RenounceOwnership(&_L2TxFeeVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.RenounceOwnership(&_L2TxFeeVault.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.TransferOwnership(&_L2TxFeeVault.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.TransferOwnership(&_L2TxFeeVault.TransactOpts, _newOwner)
}

// UpdateMessenger is a paid mutator transaction binding the contract method 0x9e7adc79.
//
// Solidity: function updateMessenger(address _newMessenger) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) UpdateMessenger(opts *bind.TransactOpts, _newMessenger common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "updateMessenger", _newMessenger)
}

// UpdateMessenger is a paid mutator transaction binding the contract method 0x9e7adc79.
//
// Solidity: function updateMessenger(address _newMessenger) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) UpdateMessenger(_newMessenger common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateMessenger(&_L2TxFeeVault.TransactOpts, _newMessenger)
}

// UpdateMessenger is a paid mutator transaction binding the contract method 0x9e7adc79.
//
// Solidity: function updateMessenger(address _newMessenger) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) UpdateMessenger(_newMessenger common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateMessenger(&_L2TxFeeVault.TransactOpts, _newMessenger)
}

// UpdateMinWithdrawAmount is a paid mutator transaction binding the contract method 0xff4f3546.
//
// Solidity: function updateMinWithdrawAmount(uint256 _newMinWithdrawAmount) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) UpdateMinWithdrawAmount(opts *bind.TransactOpts, _newMinWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "updateMinWithdrawAmount", _newMinWithdrawAmount)
}

// UpdateMinWithdrawAmount is a paid mutator transaction binding the contract method 0xff4f3546.
//
// Solidity: function updateMinWithdrawAmount(uint256 _newMinWithdrawAmount) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) UpdateMinWithdrawAmount(_newMinWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateMinWithdrawAmount(&_L2TxFeeVault.TransactOpts, _newMinWithdrawAmount)
}

// UpdateMinWithdrawAmount is a paid mutator transaction binding the contract method 0xff4f3546.
//
// Solidity: function updateMinWithdrawAmount(uint256 _newMinWithdrawAmount) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) UpdateMinWithdrawAmount(_newMinWithdrawAmount *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateMinWithdrawAmount(&_L2TxFeeVault.TransactOpts, _newMinWithdrawAmount)
}

// UpdateRecipient is a paid mutator transaction binding the contract method 0xfeec756c.
//
// Solidity: function updateRecipient(address _newRecipient) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) UpdateRecipient(opts *bind.TransactOpts, _newRecipient common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "updateRecipient", _newRecipient)
}

// UpdateRecipient is a paid mutator transaction binding the contract method 0xfeec756c.
//
// Solidity: function updateRecipient(address _newRecipient) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) UpdateRecipient(_newRecipient common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateRecipient(&_L2TxFeeVault.TransactOpts, _newRecipient)
}

// UpdateRecipient is a paid mutator transaction binding the contract method 0xfeec756c.
//
// Solidity: function updateRecipient(address _newRecipient) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) UpdateRecipient(_newRecipient common.Address) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.UpdateRecipient(&_L2TxFeeVault.TransactOpts, _newRecipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) Withdraw(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "withdraw", _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Withdraw(&_L2TxFeeVault.TransactOpts, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _value) returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) Withdraw(_value *big.Int) (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Withdraw(&_L2TxFeeVault.TransactOpts, _value)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) Withdraw0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.Transact(opts, "withdraw0")
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) Withdraw0() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Withdraw0(&_L2TxFeeVault.TransactOpts)
}

// Withdraw0 is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) Withdraw0() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Withdraw0(&_L2TxFeeVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2TxFeeVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2TxFeeVault *L2TxFeeVaultSession) Receive() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Receive(&_L2TxFeeVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_L2TxFeeVault *L2TxFeeVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _L2TxFeeVault.Contract.Receive(&_L2TxFeeVault.TransactOpts)
}

// L2TxFeeVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L2TxFeeVault contract.
type L2TxFeeVaultOwnershipTransferredIterator struct {
	Event *L2TxFeeVaultOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L2TxFeeVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L2TxFeeVaultOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L2TxFeeVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultOwnershipTransferred represents a OwnershipTransferred event raised by the L2TxFeeVault contract.
type L2TxFeeVaultOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _oldOwner []common.Address, _newOwner []common.Address) (*L2TxFeeVaultOwnershipTransferredIterator, error) {

	var _oldOwnerRule []interface{}
	for _, _oldOwnerItem := range _oldOwner {
		_oldOwnerRule = append(_oldOwnerRule, _oldOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "OwnershipTransferred", _oldOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultOwnershipTransferredIterator{contract: _L2TxFeeVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultOwnershipTransferred, _oldOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _oldOwnerRule []interface{}
	for _, _oldOwnerItem := range _oldOwner {
		_oldOwnerRule = append(_oldOwnerRule, _oldOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "OwnershipTransferred", _oldOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultOwnershipTransferred)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _oldOwner, address indexed _newOwner)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseOwnershipTransferred(log types.Log) (*L2TxFeeVaultOwnershipTransferred, error) {
	event := new(L2TxFeeVaultOwnershipTransferred)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TxFeeVaultUpdateMessengerIterator is returned from FilterUpdateMessenger and is used to iterate over the raw logs and unpacked data for UpdateMessenger events raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateMessengerIterator struct {
	Event *L2TxFeeVaultUpdateMessenger // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L2TxFeeVaultUpdateMessengerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultUpdateMessenger)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L2TxFeeVaultUpdateMessenger)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L2TxFeeVaultUpdateMessengerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultUpdateMessengerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultUpdateMessenger represents a UpdateMessenger event raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateMessenger struct {
	OldMessenger common.Address
	NewMessenger common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateMessenger is a free log retrieval operation binding the contract event 0x1c928c417a10a21c3cddad148c5dba5d710e4b1442d6d8a36de345935ad84612.
//
// Solidity: event UpdateMessenger(address indexed oldMessenger, address indexed newMessenger)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterUpdateMessenger(opts *bind.FilterOpts, oldMessenger []common.Address, newMessenger []common.Address) (*L2TxFeeVaultUpdateMessengerIterator, error) {

	var oldMessengerRule []interface{}
	for _, oldMessengerItem := range oldMessenger {
		oldMessengerRule = append(oldMessengerRule, oldMessengerItem)
	}
	var newMessengerRule []interface{}
	for _, newMessengerItem := range newMessenger {
		newMessengerRule = append(newMessengerRule, newMessengerItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "UpdateMessenger", oldMessengerRule, newMessengerRule)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultUpdateMessengerIterator{contract: _L2TxFeeVault.contract, event: "UpdateMessenger", logs: logs, sub: sub}, nil
}

// WatchUpdateMessenger is a free log subscription operation binding the contract event 0x1c928c417a10a21c3cddad148c5dba5d710e4b1442d6d8a36de345935ad84612.
//
// Solidity: event UpdateMessenger(address indexed oldMessenger, address indexed newMessenger)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchUpdateMessenger(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultUpdateMessenger, oldMessenger []common.Address, newMessenger []common.Address) (event.Subscription, error) {

	var oldMessengerRule []interface{}
	for _, oldMessengerItem := range oldMessenger {
		oldMessengerRule = append(oldMessengerRule, oldMessengerItem)
	}
	var newMessengerRule []interface{}
	for _, newMessengerItem := range newMessenger {
		newMessengerRule = append(newMessengerRule, newMessengerItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "UpdateMessenger", oldMessengerRule, newMessengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultUpdateMessenger)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateMessenger", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateMessenger is a log parse operation binding the contract event 0x1c928c417a10a21c3cddad148c5dba5d710e4b1442d6d8a36de345935ad84612.
//
// Solidity: event UpdateMessenger(address indexed oldMessenger, address indexed newMessenger)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseUpdateMessenger(log types.Log) (*L2TxFeeVaultUpdateMessenger, error) {
	event := new(L2TxFeeVaultUpdateMessenger)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateMessenger", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TxFeeVaultUpdateMinWithdrawAmountIterator is returned from FilterUpdateMinWithdrawAmount and is used to iterate over the raw logs and unpacked data for UpdateMinWithdrawAmount events raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateMinWithdrawAmountIterator struct {
	Event *L2TxFeeVaultUpdateMinWithdrawAmount // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L2TxFeeVaultUpdateMinWithdrawAmountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultUpdateMinWithdrawAmount)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L2TxFeeVaultUpdateMinWithdrawAmount)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L2TxFeeVaultUpdateMinWithdrawAmountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultUpdateMinWithdrawAmountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultUpdateMinWithdrawAmount represents a UpdateMinWithdrawAmount event raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateMinWithdrawAmount struct {
	OldMinWithdrawAmount *big.Int
	NewMinWithdrawAmount *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateMinWithdrawAmount is a free log retrieval operation binding the contract event 0x0d3c80219fe57713b9f9c83d1e51426792d0c14d8e330e65b102571816140965.
//
// Solidity: event UpdateMinWithdrawAmount(uint256 oldMinWithdrawAmount, uint256 newMinWithdrawAmount)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterUpdateMinWithdrawAmount(opts *bind.FilterOpts) (*L2TxFeeVaultUpdateMinWithdrawAmountIterator, error) {

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "UpdateMinWithdrawAmount")
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultUpdateMinWithdrawAmountIterator{contract: _L2TxFeeVault.contract, event: "UpdateMinWithdrawAmount", logs: logs, sub: sub}, nil
}

// WatchUpdateMinWithdrawAmount is a free log subscription operation binding the contract event 0x0d3c80219fe57713b9f9c83d1e51426792d0c14d8e330e65b102571816140965.
//
// Solidity: event UpdateMinWithdrawAmount(uint256 oldMinWithdrawAmount, uint256 newMinWithdrawAmount)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchUpdateMinWithdrawAmount(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultUpdateMinWithdrawAmount) (event.Subscription, error) {

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "UpdateMinWithdrawAmount")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultUpdateMinWithdrawAmount)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateMinWithdrawAmount", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateMinWithdrawAmount is a log parse operation binding the contract event 0x0d3c80219fe57713b9f9c83d1e51426792d0c14d8e330e65b102571816140965.
//
// Solidity: event UpdateMinWithdrawAmount(uint256 oldMinWithdrawAmount, uint256 newMinWithdrawAmount)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseUpdateMinWithdrawAmount(log types.Log) (*L2TxFeeVaultUpdateMinWithdrawAmount, error) {
	event := new(L2TxFeeVaultUpdateMinWithdrawAmount)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateMinWithdrawAmount", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TxFeeVaultUpdateRecipientIterator is returned from FilterUpdateRecipient and is used to iterate over the raw logs and unpacked data for UpdateRecipient events raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateRecipientIterator struct {
	Event *L2TxFeeVaultUpdateRecipient // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L2TxFeeVaultUpdateRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultUpdateRecipient)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L2TxFeeVaultUpdateRecipient)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L2TxFeeVaultUpdateRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultUpdateRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultUpdateRecipient represents a UpdateRecipient event raised by the L2TxFeeVault contract.
type L2TxFeeVaultUpdateRecipient struct {
	OldRecipient common.Address
	NewRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUpdateRecipient is a free log retrieval operation binding the contract event 0x7e1e96961a397c8aa26162fe259cc837afc95e33aad4945ddc61c18dabb7a6ad.
//
// Solidity: event UpdateRecipient(address indexed oldRecipient, address indexed newRecipient)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterUpdateRecipient(opts *bind.FilterOpts, oldRecipient []common.Address, newRecipient []common.Address) (*L2TxFeeVaultUpdateRecipientIterator, error) {

	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "UpdateRecipient", oldRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultUpdateRecipientIterator{contract: _L2TxFeeVault.contract, event: "UpdateRecipient", logs: logs, sub: sub}, nil
}

// WatchUpdateRecipient is a free log subscription operation binding the contract event 0x7e1e96961a397c8aa26162fe259cc837afc95e33aad4945ddc61c18dabb7a6ad.
//
// Solidity: event UpdateRecipient(address indexed oldRecipient, address indexed newRecipient)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchUpdateRecipient(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultUpdateRecipient, oldRecipient []common.Address, newRecipient []common.Address) (event.Subscription, error) {

	var oldRecipientRule []interface{}
	for _, oldRecipientItem := range oldRecipient {
		oldRecipientRule = append(oldRecipientRule, oldRecipientItem)
	}
	var newRecipientRule []interface{}
	for _, newRecipientItem := range newRecipient {
		newRecipientRule = append(newRecipientRule, newRecipientItem)
	}

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "UpdateRecipient", oldRecipientRule, newRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultUpdateRecipient)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateRecipient", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateRecipient is a log parse operation binding the contract event 0x7e1e96961a397c8aa26162fe259cc837afc95e33aad4945ddc61c18dabb7a6ad.
//
// Solidity: event UpdateRecipient(address indexed oldRecipient, address indexed newRecipient)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseUpdateRecipient(log types.Log) (*L2TxFeeVaultUpdateRecipient, error) {
	event := new(L2TxFeeVaultUpdateRecipient)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "UpdateRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TxFeeVaultWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the L2TxFeeVault contract.
type L2TxFeeVaultWithdrawalIterator struct {
	Event *L2TxFeeVaultWithdrawal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L2TxFeeVaultWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2TxFeeVaultWithdrawal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L2TxFeeVaultWithdrawal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L2TxFeeVaultWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TxFeeVaultWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2TxFeeVaultWithdrawal represents a Withdrawal event raised by the L2TxFeeVault contract.
type L2TxFeeVaultWithdrawal struct {
	Value *big.Int
	To    common.Address
	From  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*L2TxFeeVaultWithdrawalIterator, error) {

	logs, sub, err := _L2TxFeeVault.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &L2TxFeeVaultWithdrawalIterator{contract: _L2TxFeeVault.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *L2TxFeeVaultWithdrawal) (event.Subscription, error) {

	logs, sub, err := _L2TxFeeVault.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2TxFeeVaultWithdrawal)
				if err := _L2TxFeeVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawal is a log parse operation binding the contract event 0xc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba.
//
// Solidity: event Withdrawal(uint256 value, address to, address from)
func (_L2TxFeeVault *L2TxFeeVaultFilterer) ParseWithdrawal(log types.Log) (*L2TxFeeVaultWithdrawal, error) {
	event := new(L2TxFeeVaultWithdrawal)
	if err := _L2TxFeeVault.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
