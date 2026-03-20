// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
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
	_ = abi.ConvertType
)

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_own_balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_external_balance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"Request_Approved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_multisig_address\",\"type\":\"address\"}],\"name\":\"ChangeMultisigAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"external_balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"multisig_contract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"own_balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"start_balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b506040516107d23803806107d2833981016040819052602b916069565b5f8390556001829055603c838360a9565b600255600480546001600160a01b0319166001600160a01b039290921691909117905550505f60035560cd565b5f5f5f60608486031215607a575f5ffd5b83516020850151604086015191945092506001600160a01b0381168114609e575f5ffd5b809150509250925092565b8082018082111560c757634e487b7160e01b5f52601160045260245ffd5b92915050565b6106f8806100da5f395ff3fe608060405234801561000f575f5ffd5b5060043610610090575f3560e01c80636ce5edff116100635780636ce5edff146100eb5780639bfb252914610116578063affed0e01461011e578063dc7ccf4314610127578063fc0c546a14610130575f5ffd5b806302411f5a1461009457806322a272c3146100a957806362ef5dce146100bc57806369ca02dd146100d8575b5f5ffd5b6100a76100a2366004610500565b610143565b005b6100a76100b73660046105d8565b610268565b6100c560025481565b6040519081526020015b60405180910390f35b6100a76100e63660046105f8565b6102dd565b6005546100fe906001600160a01b031681565b6040516001600160a01b0390911681526020016100cf565b6100c55f5481565b6100c560035481565b6100c560015481565b6004546100fe906001600160a01b031681565b5f82116101975760405162461bcd60e51b815260206004820181905260248201527f416d6f756e74206d7573742062652067726561746572207468616e207a65726f60448201526064015b60405180910390fd5b600480546040516323b872dd60e01b81523392810192909252306024830152604482018490525f916001600160a01b03909116906323b872dd906064016020604051808303815f875af11580156101f0573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102149190610620565b9050806102595760405162461bcd60e51b81526020600482015260136024820152721d1c985b9cd9995c919c9bdb4819985a5b1959606a1b604482015260640161018e565b6102638383610403565b505050565b739bb51a70c346d9cff52b382e85851d78671accd033146102bb5760405162461bcd60e51b815260206004820152600d60248201526c6e6f742061206e6f2d6869766560981b604482015260640161018e565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b6005546001600160a01b031633146103435760405162461bcd60e51b815260206004820152602360248201527f43616c6c6572206973206e6f7420746865206d756c746973696720636f6e74726044820152621858dd60ea1b606482015260840161018e565b6004805460405163a9059cbb60e01b81526001600160a01b0385811693820193909352602481018490525f929091169063a9059cbb906044016020604051808303815f875af1158015610398573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103bc9190610620565b9050806102635760405162461bcd60e51b8152602060048201526015602482015274151bdad95b881d1c985b9cd9995c8819985a5b1959605a1b604482015260640161018e565b81600154116104645760405162461bcd60e51b815260206004820152602760248201527f496e73756666696369656e742066756e6473206f6e2064657374696e6174696f604482015266371031b430b4b760c91b606482015260840161018e565b8160015f8282546104759190610653565b92505081905550815f5f82825461048c919061066c565b909155505060035460405133917f4ae16418416b042e51fc7d3cd38a51016f389d862442eaa688117bcdd497fff9916104c991869186919061067f565b60405180910390a2600160035f8282546104e3919061066c565b90915550505050565b634e487b7160e01b5f52604160045260245ffd5b5f5f60408385031215610511575f5ffd5b82359150602083013567ffffffffffffffff81111561052e575f5ffd5b8301601f8101851361053e575f5ffd5b803567ffffffffffffffff811115610558576105586104ec565b604051601f8201601f19908116603f0116810167ffffffffffffffff81118282101715610587576105876104ec565b60405281815282820160200187101561059e575f5ffd5b816020840160208301375f602083830101528093505050509250929050565b80356001600160a01b03811681146105d3575f5ffd5b919050565b5f602082840312156105e8575f5ffd5b6105f1826105bd565b9392505050565b5f5f60408385031215610609575f5ffd5b610612836105bd565b946020939093013593505050565b5f60208284031215610630575f5ffd5b815180151581146105f1575f5ffd5b634e487b7160e01b5f52601160045260245ffd5b818103818111156106665761066661063f565b92915050565b808201808211156106665761066661063f565b838152606060208201525f83518060608401528060208601608085015e5f608082850101526080601f19601f83011684010191505082604083015294935050505056fea26469706673582212202d94703344980e746642853a4950ab56f0e8343360bf52d315cd6692255f3a0664736f6c63430008220033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _own_balance *big.Int, _external_balance *big.Int, _token common.Address) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend, _own_balance, _external_balance, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// ExternalBalance is a free data retrieval call binding the contract method 0xdc7ccf43.
//
// Solidity: function external_balance() view returns(uint256)
func (_Bridge *BridgeCaller) ExternalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "external_balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExternalBalance is a free data retrieval call binding the contract method 0xdc7ccf43.
//
// Solidity: function external_balance() view returns(uint256)
func (_Bridge *BridgeSession) ExternalBalance() (*big.Int, error) {
	return _Bridge.Contract.ExternalBalance(&_Bridge.CallOpts)
}

// ExternalBalance is a free data retrieval call binding the contract method 0xdc7ccf43.
//
// Solidity: function external_balance() view returns(uint256)
func (_Bridge *BridgeCallerSession) ExternalBalance() (*big.Int, error) {
	return _Bridge.Contract.ExternalBalance(&_Bridge.CallOpts)
}

// MultisigContract is a free data retrieval call binding the contract method 0x6ce5edff.
//
// Solidity: function multisig_contract() view returns(address)
func (_Bridge *BridgeCaller) MultisigContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "multisig_contract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MultisigContract is a free data retrieval call binding the contract method 0x6ce5edff.
//
// Solidity: function multisig_contract() view returns(address)
func (_Bridge *BridgeSession) MultisigContract() (common.Address, error) {
	return _Bridge.Contract.MultisigContract(&_Bridge.CallOpts)
}

// MultisigContract is a free data retrieval call binding the contract method 0x6ce5edff.
//
// Solidity: function multisig_contract() view returns(address)
func (_Bridge *BridgeCallerSession) MultisigContract() (common.Address, error) {
	return _Bridge.Contract.MultisigContract(&_Bridge.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeSession) Nonce() (*big.Int, error) {
	return _Bridge.Contract.Nonce(&_Bridge.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_Bridge *BridgeCallerSession) Nonce() (*big.Int, error) {
	return _Bridge.Contract.Nonce(&_Bridge.CallOpts)
}

// OwnBalance is a free data retrieval call binding the contract method 0x9bfb2529.
//
// Solidity: function own_balance() view returns(uint256)
func (_Bridge *BridgeCaller) OwnBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "own_balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnBalance is a free data retrieval call binding the contract method 0x9bfb2529.
//
// Solidity: function own_balance() view returns(uint256)
func (_Bridge *BridgeSession) OwnBalance() (*big.Int, error) {
	return _Bridge.Contract.OwnBalance(&_Bridge.CallOpts)
}

// OwnBalance is a free data retrieval call binding the contract method 0x9bfb2529.
//
// Solidity: function own_balance() view returns(uint256)
func (_Bridge *BridgeCallerSession) OwnBalance() (*big.Int, error) {
	return _Bridge.Contract.OwnBalance(&_Bridge.CallOpts)
}

// StartBalance is a free data retrieval call binding the contract method 0x62ef5dce.
//
// Solidity: function start_balance() view returns(uint256)
func (_Bridge *BridgeCaller) StartBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "start_balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBalance is a free data retrieval call binding the contract method 0x62ef5dce.
//
// Solidity: function start_balance() view returns(uint256)
func (_Bridge *BridgeSession) StartBalance() (*big.Int, error) {
	return _Bridge.Contract.StartBalance(&_Bridge.CallOpts)
}

// StartBalance is a free data retrieval call binding the contract method 0x62ef5dce.
//
// Solidity: function start_balance() view returns(uint256)
func (_Bridge *BridgeCallerSession) StartBalance() (*big.Int, error) {
	return _Bridge.Contract.StartBalance(&_Bridge.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bridge *BridgeCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bridge *BridgeSession) Token() (common.Address, error) {
	return _Bridge.Contract.Token(&_Bridge.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bridge *BridgeCallerSession) Token() (common.Address, error) {
	return _Bridge.Contract.Token(&_Bridge.CallOpts)
}

// ChangeMultisigAddress is a paid mutator transaction binding the contract method 0x22a272c3.
//
// Solidity: function ChangeMultisigAddress(address _multisig_address) returns()
func (_Bridge *BridgeTransactor) ChangeMultisigAddress(opts *bind.TransactOpts, _multisig_address common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "ChangeMultisigAddress", _multisig_address)
}

// ChangeMultisigAddress is a paid mutator transaction binding the contract method 0x22a272c3.
//
// Solidity: function ChangeMultisigAddress(address _multisig_address) returns()
func (_Bridge *BridgeSession) ChangeMultisigAddress(_multisig_address common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeMultisigAddress(&_Bridge.TransactOpts, _multisig_address)
}

// ChangeMultisigAddress is a paid mutator transaction binding the contract method 0x22a272c3.
//
// Solidity: function ChangeMultisigAddress(address _multisig_address) returns()
func (_Bridge *BridgeTransactorSession) ChangeMultisigAddress(_multisig_address common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ChangeMultisigAddress(&_Bridge.TransactOpts, _multisig_address)
}

// Deposit is a paid mutator transaction binding the contract method 0x02411f5a.
//
// Solidity: function Deposit(uint256 amount, string receiver) returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int, receiver string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "Deposit", amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x02411f5a.
//
// Solidity: function Deposit(uint256 amount, string receiver) returns()
func (_Bridge *BridgeSession) Deposit(amount *big.Int, receiver string) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, amount, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x02411f5a.
//
// Solidity: function Deposit(uint256 amount, string receiver) returns()
func (_Bridge *BridgeTransactorSession) Deposit(amount *big.Int, receiver string) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, amount, receiver)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address recipient, uint256 amount) returns()
func (_Bridge *BridgeTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "Transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address recipient, uint256 amount) returns()
func (_Bridge *BridgeSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Transfer(&_Bridge.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x69ca02dd.
//
// Solidity: function Transfer(address recipient, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Transfer(&_Bridge.TransactOpts, recipient, amount)
}

// BridgeRequestApprovedIterator is returned from FilterRequestApproved and is used to iterate over the raw logs and unpacked data for RequestApproved events raised by the Bridge contract.
type BridgeRequestApprovedIterator struct {
	Event *BridgeRequestApproved // Event containing the contract specifics and raw log

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
func (it *BridgeRequestApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRequestApproved)
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
		it.Event = new(BridgeRequestApproved)
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
func (it *BridgeRequestApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRequestApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRequestApproved represents a RequestApproved event raised by the Bridge contract.
type BridgeRequestApproved struct {
	Sender   common.Address
	Amount   *big.Int
	Receiver string
	Nonce    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequestApproved is a free log retrieval operation binding the contract event 0x4ae16418416b042e51fc7d3cd38a51016f389d862442eaa688117bcdd497fff9.
//
// Solidity: event Request_Approved(address indexed sender, uint256 amount, string receiver, uint256 nonce)
func (_Bridge *BridgeFilterer) FilterRequestApproved(opts *bind.FilterOpts, sender []common.Address) (*BridgeRequestApprovedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Request_Approved", senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRequestApprovedIterator{contract: _Bridge.contract, event: "Request_Approved", logs: logs, sub: sub}, nil
}

// WatchRequestApproved is a free log subscription operation binding the contract event 0x4ae16418416b042e51fc7d3cd38a51016f389d862442eaa688117bcdd497fff9.
//
// Solidity: event Request_Approved(address indexed sender, uint256 amount, string receiver, uint256 nonce)
func (_Bridge *BridgeFilterer) WatchRequestApproved(opts *bind.WatchOpts, sink chan<- *BridgeRequestApproved, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Request_Approved", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRequestApproved)
				if err := _Bridge.contract.UnpackLog(event, "Request_Approved", log); err != nil {
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

// ParseRequestApproved is a log parse operation binding the contract event 0x4ae16418416b042e51fc7d3cd38a51016f389d862442eaa688117bcdd497fff9.
//
// Solidity: event Request_Approved(address indexed sender, uint256 amount, string receiver, uint256 nonce)
func (_Bridge *BridgeFilterer) ParseRequestApproved(log types.Log) (*BridgeRequestApproved, error) {
	event := new(BridgeRequestApproved)
	if err := _Bridge.contract.UnpackLog(event, "Request_Approved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
