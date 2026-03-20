// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package federationsync

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

// FederationsyncMetaData contains all meta data concerning the Federationsync contract.
var FederationsyncMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_federation_node_1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_federation_node_2\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_federation_node_3\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_bridgeContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"nodeIndex\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"changedBy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"NodeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"confirmedBy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"RequestConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"TransferExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"changeNode1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"changeNode2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"changeNode3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"confirmRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"federation_node_1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"federation_node_2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"federation_node_3\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"transfer_made\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"node_1_confirmation\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"node_2_confirmation\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"node_3_confirmation\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"node_1_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"node_2_recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"node_3_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"node_1_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"node_2_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"node_3_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50604051610a95380380610a95833981016040819052602b916096565b5f80546001600160a01b039586166001600160a01b031991821617909155600180549486169482169490941790935560028054928516928416929092179091556003805491909316911617905560de565b80516001600160a01b03811681146091575f5ffd5b919050565b5f5f5f5f6080858703121560a8575f5ffd5b60af85607c565b935060bb60208601607c565b925060c760408601607c565b915060d360608601607c565b905092959194509250565b6109aa806100eb5f395ff3fe608060405234801561000f575f5ffd5b5060043610610090575f3560e01c806381d12c581161006357806381d12c58146100fe57806386e74f1e146101d55780639959e541146101e8578063cd596583146101fb578063dc6e07371461020e575f5ffd5b80630562099a146100945780630c020914146100a95780635e080a07146100d95780636453aded146100eb575b5f5ffd5b6100a76100a236600461090d565b610221565b005b6001546100bc906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b5f546100bc906001600160a01b031681565b6002546100bc906001600160a01b031681565b61017761010c36600461092d565b600460208190525f918252604090912080546001820154600283015460038401549484015460059094015460ff80851696610100860482169662010000870483169663010000008104909316956001600160a01b03640100000000909404841695908416949316928a565b604080519a15158b5298151560208b01529615159789019790975293151560608801526001600160a01b03928316608088015290821660a08701521660c085015260e0840152610100830191909152610120820152610140016100d0565b6100a76101e3366004610944565b6102ce565b6100a76101f636600461090d565b610672565b6003546100bc906001600160a01b031681565b6100a761021c36600461090d565b61071a565b6002546001600160a01b031633146102805760405162461bcd60e51b815260206004820152601d60248201527f4f6e6c79206e6f646520332063616e206368616e676520697473656c6600000060448201526064015b60405180910390fd5b600280546001600160a01b0319166001600160a01b03831690811790915560405133906003907ff5becf5ef6bf281ea3c4f617c266a872c7662757c90aa5f5cdffe02bd2023975905f90a450565b5f818152600460205260409020805460ff161561032d5760405162461bcd60e51b815260206004820152601960248201527f5472616e7366657220616c7265616479206578656375746564000000000000006044820152606401610277565b5f546001600160a01b031633036103ca578054610100900460ff16156103955760405162461bcd60e51b815260206004820152601860248201527f4e6f6465203120616c726561647920636f6e6669726d656400000000000000006044820152606401610277565b80546001600160a01b0385166401000000000263ffff0100600160c01b03199091161761010017815560038101839055610567565b6001546001600160a01b0316330361046e57805462010000900460ff16156104345760405162461bcd60e51b815260206004820152601860248201527f4e6f6465203220616c726561647920636f6e6669726d656400000000000000006044820152606401610277565b805462ff00001916620100001781556001810180546001600160a01b0386166001600160a01b031990911617905560048101839055610567565b6002546001600160a01b031633036105155780546301000000900460ff16156104d95760405162461bcd60e51b815260206004820152601860248201527f4e6f6465203320616c726561647920636f6e6669726d656400000000000000006044820152606401610277565b805463ff000000191663010000001781556002810180546001600160a01b0386166001600160a01b031990911617905560058101839055610567565b60405162461bcd60e51b815260206004820152602160248201527f4f6e6c792066656465726174696f6e206e6f6465732063616e20636f6e6669726044820152606d60f81b6064820152608401610277565b60408051848152602081018490526001600160a01b0386169133917f6718bbbc3a1333e921501591df5fe974bd1de1e3dc5498203221a126967067b8910160405180910390a36105b6826107bf565b1561066c57805460ff191660011781556003546040516369ca02dd60e01b81526001600160a01b03868116600483015260248201869052909116906369ca02dd906044015f604051808303815f87803b158015610611575f5ffd5b505af1158015610623573d5f5f3e3d5ffd5b505060408051868152602081018690526001600160a01b03881693507ff8d62a11d316e69c1460662f86b7046003826db86fdc0d5ffa8c3adaa24f0a5892500160405180910390a25b50505050565b6001546001600160a01b031633146106cc5760405162461bcd60e51b815260206004820152601d60248201527f4f6e6c79206e6f646520322063616e206368616e676520697473656c660000006044820152606401610277565b600180546001600160a01b0319166001600160a01b03831690811790915560405133906002907ff5becf5ef6bf281ea3c4f617c266a872c7662757c90aa5f5cdffe02bd2023975905f90a450565b5f546001600160a01b031633146107735760405162461bcd60e51b815260206004820152601d60248201527f4f6e6c79206e6f646520312063616e206368616e676520697473656c660000006044820152606401610277565b5f80546001600160a01b0319166001600160a01b0383169081178255604051909133916001917ff5becf5ef6bf281ea3c4f617c266a872c7662757c90aa5f5cdffe02bd202397591a450565b5f8181526004602052604081208054610100900460ff1680156107e95750805462010000900460ff165b1561082c576001810154815464010000000090046001600160a01b03908116911614801561081e575080600401548160030154145b1561082c5750600192915050565b8054610100900460ff16801561084a575080546301000000900460ff165b1561088d576002810154815464010000000090046001600160a01b03908116911614801561087f575080600501548160030154145b1561088d5750600192915050565b805462010000900460ff1680156108ac575080546301000000900460ff165b156108ea57600281015460018201546001600160a01b0390811691161480156108dc575080600501548160040154145b156108ea5750600192915050565b505f92915050565b80356001600160a01b0381168114610908575f5ffd5b919050565b5f6020828403121561091d575f5ffd5b610926826108f2565b9392505050565b5f6020828403121561093d575f5ffd5b5035919050565b5f5f5f60608486031215610956575f5ffd5b61095f846108f2565b9560208501359550604090940135939250505056fea2646970667358221220b1768224dd6bb8e5df3bf7fc2277582d6779fbafb067b006c821a4678b9bc71a64736f6c63430008220033",
}

// FederationsyncABI is the input ABI used to generate the binding from.
// Deprecated: Use FederationsyncMetaData.ABI instead.
var FederationsyncABI = FederationsyncMetaData.ABI

// FederationsyncBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FederationsyncMetaData.Bin instead.
var FederationsyncBin = FederationsyncMetaData.Bin

// DeployFederationsync deploys a new Ethereum contract, binding an instance of Federationsync to it.
func DeployFederationsync(auth *bind.TransactOpts, backend bind.ContractBackend, _federation_node_1 common.Address, _federation_node_2 common.Address, _federation_node_3 common.Address, _bridgeContract common.Address) (common.Address, *types.Transaction, *Federationsync, error) {
	parsed, err := FederationsyncMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FederationsyncBin), backend, _federation_node_1, _federation_node_2, _federation_node_3, _bridgeContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Federationsync{FederationsyncCaller: FederationsyncCaller{contract: contract}, FederationsyncTransactor: FederationsyncTransactor{contract: contract}, FederationsyncFilterer: FederationsyncFilterer{contract: contract}}, nil
}

// Federationsync is an auto generated Go binding around an Ethereum contract.
type Federationsync struct {
	FederationsyncCaller     // Read-only binding to the contract
	FederationsyncTransactor // Write-only binding to the contract
	FederationsyncFilterer   // Log filterer for contract events
}

// FederationsyncCaller is an auto generated read-only Go binding around an Ethereum contract.
type FederationsyncCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FederationsyncTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FederationsyncTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FederationsyncFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FederationsyncFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FederationsyncSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FederationsyncSession struct {
	Contract     *Federationsync   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FederationsyncCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FederationsyncCallerSession struct {
	Contract *FederationsyncCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FederationsyncTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FederationsyncTransactorSession struct {
	Contract     *FederationsyncTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FederationsyncRaw is an auto generated low-level Go binding around an Ethereum contract.
type FederationsyncRaw struct {
	Contract *Federationsync // Generic contract binding to access the raw methods on
}

// FederationsyncCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FederationsyncCallerRaw struct {
	Contract *FederationsyncCaller // Generic read-only contract binding to access the raw methods on
}

// FederationsyncTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FederationsyncTransactorRaw struct {
	Contract *FederationsyncTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFederationsync creates a new instance of Federationsync, bound to a specific deployed contract.
func NewFederationsync(address common.Address, backend bind.ContractBackend) (*Federationsync, error) {
	contract, err := bindFederationsync(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Federationsync{FederationsyncCaller: FederationsyncCaller{contract: contract}, FederationsyncTransactor: FederationsyncTransactor{contract: contract}, FederationsyncFilterer: FederationsyncFilterer{contract: contract}}, nil
}

// NewFederationsyncCaller creates a new read-only instance of Federationsync, bound to a specific deployed contract.
func NewFederationsyncCaller(address common.Address, caller bind.ContractCaller) (*FederationsyncCaller, error) {
	contract, err := bindFederationsync(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FederationsyncCaller{contract: contract}, nil
}

// NewFederationsyncTransactor creates a new write-only instance of Federationsync, bound to a specific deployed contract.
func NewFederationsyncTransactor(address common.Address, transactor bind.ContractTransactor) (*FederationsyncTransactor, error) {
	contract, err := bindFederationsync(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FederationsyncTransactor{contract: contract}, nil
}

// NewFederationsyncFilterer creates a new log filterer instance of Federationsync, bound to a specific deployed contract.
func NewFederationsyncFilterer(address common.Address, filterer bind.ContractFilterer) (*FederationsyncFilterer, error) {
	contract, err := bindFederationsync(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FederationsyncFilterer{contract: contract}, nil
}

// bindFederationsync binds a generic wrapper to an already deployed contract.
func bindFederationsync(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FederationsyncMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Federationsync *FederationsyncRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Federationsync.Contract.FederationsyncCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Federationsync *FederationsyncRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Federationsync.Contract.FederationsyncTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Federationsync *FederationsyncRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Federationsync.Contract.FederationsyncTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Federationsync *FederationsyncCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Federationsync.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Federationsync *FederationsyncTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Federationsync.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Federationsync *FederationsyncTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Federationsync.Contract.contract.Transact(opts, method, params...)
}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_Federationsync *FederationsyncCaller) BridgeContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Federationsync.contract.Call(opts, &out, "bridgeContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_Federationsync *FederationsyncSession) BridgeContract() (common.Address, error) {
	return _Federationsync.Contract.BridgeContract(&_Federationsync.CallOpts)
}

// BridgeContract is a free data retrieval call binding the contract method 0xcd596583.
//
// Solidity: function bridgeContract() view returns(address)
func (_Federationsync *FederationsyncCallerSession) BridgeContract() (common.Address, error) {
	return _Federationsync.Contract.BridgeContract(&_Federationsync.CallOpts)
}

// FederationNode1 is a free data retrieval call binding the contract method 0x5e080a07.
//
// Solidity: function federation_node_1() view returns(address)
func (_Federationsync *FederationsyncCaller) FederationNode1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Federationsync.contract.Call(opts, &out, "federation_node_1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FederationNode1 is a free data retrieval call binding the contract method 0x5e080a07.
//
// Solidity: function federation_node_1() view returns(address)
func (_Federationsync *FederationsyncSession) FederationNode1() (common.Address, error) {
	return _Federationsync.Contract.FederationNode1(&_Federationsync.CallOpts)
}

// FederationNode1 is a free data retrieval call binding the contract method 0x5e080a07.
//
// Solidity: function federation_node_1() view returns(address)
func (_Federationsync *FederationsyncCallerSession) FederationNode1() (common.Address, error) {
	return _Federationsync.Contract.FederationNode1(&_Federationsync.CallOpts)
}

// FederationNode2 is a free data retrieval call binding the contract method 0x0c020914.
//
// Solidity: function federation_node_2() view returns(address)
func (_Federationsync *FederationsyncCaller) FederationNode2(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Federationsync.contract.Call(opts, &out, "federation_node_2")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FederationNode2 is a free data retrieval call binding the contract method 0x0c020914.
//
// Solidity: function federation_node_2() view returns(address)
func (_Federationsync *FederationsyncSession) FederationNode2() (common.Address, error) {
	return _Federationsync.Contract.FederationNode2(&_Federationsync.CallOpts)
}

// FederationNode2 is a free data retrieval call binding the contract method 0x0c020914.
//
// Solidity: function federation_node_2() view returns(address)
func (_Federationsync *FederationsyncCallerSession) FederationNode2() (common.Address, error) {
	return _Federationsync.Contract.FederationNode2(&_Federationsync.CallOpts)
}

// FederationNode3 is a free data retrieval call binding the contract method 0x6453aded.
//
// Solidity: function federation_node_3() view returns(address)
func (_Federationsync *FederationsyncCaller) FederationNode3(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Federationsync.contract.Call(opts, &out, "federation_node_3")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FederationNode3 is a free data retrieval call binding the contract method 0x6453aded.
//
// Solidity: function federation_node_3() view returns(address)
func (_Federationsync *FederationsyncSession) FederationNode3() (common.Address, error) {
	return _Federationsync.Contract.FederationNode3(&_Federationsync.CallOpts)
}

// FederationNode3 is a free data retrieval call binding the contract method 0x6453aded.
//
// Solidity: function federation_node_3() view returns(address)
func (_Federationsync *FederationsyncCallerSession) FederationNode3() (common.Address, error) {
	return _Federationsync.Contract.FederationNode3(&_Federationsync.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(bool transfer_made, bool node_1_confirmation, bool node_2_confirmation, bool node_3_confirmation, address node_1_recipient, address node_2_recipient, address node_3_recipient, uint256 node_1_amount, uint256 node_2_amount, uint256 node_3_amount)
func (_Federationsync *FederationsyncCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TransferMade      bool
	Node1Confirmation bool
	Node2Confirmation bool
	Node3Confirmation bool
	Node1Recipient    common.Address
	Node2Recipient    common.Address
	Node3Recipient    common.Address
	Node1Amount       *big.Int
	Node2Amount       *big.Int
	Node3Amount       *big.Int
}, error) {
	var out []interface{}
	err := _Federationsync.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		TransferMade      bool
		Node1Confirmation bool
		Node2Confirmation bool
		Node3Confirmation bool
		Node1Recipient    common.Address
		Node2Recipient    common.Address
		Node3Recipient    common.Address
		Node1Amount       *big.Int
		Node2Amount       *big.Int
		Node3Amount       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TransferMade = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Node1Confirmation = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Node2Confirmation = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Node3Confirmation = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.Node1Recipient = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Node2Recipient = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Node3Recipient = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Node1Amount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Node2Amount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Node3Amount = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(bool transfer_made, bool node_1_confirmation, bool node_2_confirmation, bool node_3_confirmation, address node_1_recipient, address node_2_recipient, address node_3_recipient, uint256 node_1_amount, uint256 node_2_amount, uint256 node_3_amount)
func (_Federationsync *FederationsyncSession) Requests(arg0 *big.Int) (struct {
	TransferMade      bool
	Node1Confirmation bool
	Node2Confirmation bool
	Node3Confirmation bool
	Node1Recipient    common.Address
	Node2Recipient    common.Address
	Node3Recipient    common.Address
	Node1Amount       *big.Int
	Node2Amount       *big.Int
	Node3Amount       *big.Int
}, error) {
	return _Federationsync.Contract.Requests(&_Federationsync.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(bool transfer_made, bool node_1_confirmation, bool node_2_confirmation, bool node_3_confirmation, address node_1_recipient, address node_2_recipient, address node_3_recipient, uint256 node_1_amount, uint256 node_2_amount, uint256 node_3_amount)
func (_Federationsync *FederationsyncCallerSession) Requests(arg0 *big.Int) (struct {
	TransferMade      bool
	Node1Confirmation bool
	Node2Confirmation bool
	Node3Confirmation bool
	Node1Recipient    common.Address
	Node2Recipient    common.Address
	Node3Recipient    common.Address
	Node1Amount       *big.Int
	Node2Amount       *big.Int
	Node3Amount       *big.Int
}, error) {
	return _Federationsync.Contract.Requests(&_Federationsync.CallOpts, arg0)
}

// ChangeNode1 is a paid mutator transaction binding the contract method 0xdc6e0737.
//
// Solidity: function changeNode1(address newAddress) returns()
func (_Federationsync *FederationsyncTransactor) ChangeNode1(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Federationsync.contract.Transact(opts, "changeNode1", newAddress)
}

// ChangeNode1 is a paid mutator transaction binding the contract method 0xdc6e0737.
//
// Solidity: function changeNode1(address newAddress) returns()
func (_Federationsync *FederationsyncSession) ChangeNode1(newAddress common.Address) (*types.Transaction, error) {
	return _Federationsync.Contract.ChangeNode1(&_Federationsync.TransactOpts, newAddress)
}

// ChangeNode1 is a paid mutator transaction binding the contract method 0xdc6e0737.
//
// Solidity: function changeNode1(address newAddress) returns()
func (_Federationsync *FederationsyncTransactorSession) ChangeNode1(newAddress common.Address) (*types.Transaction, error) {
	return _Federationsync.Contract.ChangeNode1(&_Federationsync.TransactOpts, newAddress)
}

// ChangeNode2 is a paid mutator transaction binding the contract method 0x9959e541.
//
// Solidity: function changeNode2(address newAddress) returns()
func (_Federationsync *FederationsyncTransactor) ChangeNode2(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Federationsync.contract.Transact(opts, "changeNode2", newAddress)
}

// ChangeNode2 is a paid mutator transaction binding the contract method 0x9959e541.
//
// Solidity: function changeNode2(address newAddress) returns()
func (_Federationsync *FederationsyncSession) ChangeNode2(newAddress common.Address) (*types.Transaction, error) {
	return _Federationsync.Contract.ChangeNode2(&_Federationsync.TransactOpts, newAddress)
}

// ChangeNode2 is a paid mutator transaction binding the contract method 0x9959e541.
//
// Solidity: function changeNode2(address newAddress) returns()
func (_Federationsync *FederationsyncTransactorSession) ChangeNode2(newAddress common.Address) (*types.Transaction, error) {
	return _Federationsync.Contract.ChangeNode2(&_Federationsync.TransactOpts, newAddress)
}

// ChangeNode3 is a paid mutator transaction binding the contract method 0x0562099a.
//
// Solidity: function changeNode3(address newAddress) returns()
func (_Federationsync *FederationsyncTransactor) ChangeNode3(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Federationsync.contract.Transact(opts, "changeNode3", newAddress)
}

// ChangeNode3 is a paid mutator transaction binding the contract method 0x0562099a.
//
// Solidity: function changeNode3(address newAddress) returns()
func (_Federationsync *FederationsyncSession) ChangeNode3(newAddress common.Address) (*types.Transaction, error) {
	return _Federationsync.Contract.ChangeNode3(&_Federationsync.TransactOpts, newAddress)
}

// ChangeNode3 is a paid mutator transaction binding the contract method 0x0562099a.
//
// Solidity: function changeNode3(address newAddress) returns()
func (_Federationsync *FederationsyncTransactorSession) ChangeNode3(newAddress common.Address) (*types.Transaction, error) {
	return _Federationsync.Contract.ChangeNode3(&_Federationsync.TransactOpts, newAddress)
}

// ConfirmRequest is a paid mutator transaction binding the contract method 0x86e74f1e.
//
// Solidity: function confirmRequest(address recipient, uint256 amount, uint256 nonce) returns()
func (_Federationsync *FederationsyncTransactor) ConfirmRequest(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Federationsync.contract.Transact(opts, "confirmRequest", recipient, amount, nonce)
}

// ConfirmRequest is a paid mutator transaction binding the contract method 0x86e74f1e.
//
// Solidity: function confirmRequest(address recipient, uint256 amount, uint256 nonce) returns()
func (_Federationsync *FederationsyncSession) ConfirmRequest(recipient common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Federationsync.Contract.ConfirmRequest(&_Federationsync.TransactOpts, recipient, amount, nonce)
}

// ConfirmRequest is a paid mutator transaction binding the contract method 0x86e74f1e.
//
// Solidity: function confirmRequest(address recipient, uint256 amount, uint256 nonce) returns()
func (_Federationsync *FederationsyncTransactorSession) ConfirmRequest(recipient common.Address, amount *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Federationsync.Contract.ConfirmRequest(&_Federationsync.TransactOpts, recipient, amount, nonce)
}

// FederationsyncNodeChangedIterator is returned from FilterNodeChanged and is used to iterate over the raw logs and unpacked data for NodeChanged events raised by the Federationsync contract.
type FederationsyncNodeChangedIterator struct {
	Event *FederationsyncNodeChanged // Event containing the contract specifics and raw log

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
func (it *FederationsyncNodeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FederationsyncNodeChanged)
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
		it.Event = new(FederationsyncNodeChanged)
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
func (it *FederationsyncNodeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FederationsyncNodeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FederationsyncNodeChanged represents a NodeChanged event raised by the Federationsync contract.
type FederationsyncNodeChanged struct {
	NodeIndex  uint8
	ChangedBy  common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNodeChanged is a free log retrieval operation binding the contract event 0xf5becf5ef6bf281ea3c4f617c266a872c7662757c90aa5f5cdffe02bd2023975.
//
// Solidity: event NodeChanged(uint8 indexed nodeIndex, address indexed changedBy, address indexed newAddress)
func (_Federationsync *FederationsyncFilterer) FilterNodeChanged(opts *bind.FilterOpts, nodeIndex []uint8, changedBy []common.Address, newAddress []common.Address) (*FederationsyncNodeChangedIterator, error) {

	var nodeIndexRule []interface{}
	for _, nodeIndexItem := range nodeIndex {
		nodeIndexRule = append(nodeIndexRule, nodeIndexItem)
	}
	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Federationsync.contract.FilterLogs(opts, "NodeChanged", nodeIndexRule, changedByRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &FederationsyncNodeChangedIterator{contract: _Federationsync.contract, event: "NodeChanged", logs: logs, sub: sub}, nil
}

// WatchNodeChanged is a free log subscription operation binding the contract event 0xf5becf5ef6bf281ea3c4f617c266a872c7662757c90aa5f5cdffe02bd2023975.
//
// Solidity: event NodeChanged(uint8 indexed nodeIndex, address indexed changedBy, address indexed newAddress)
func (_Federationsync *FederationsyncFilterer) WatchNodeChanged(opts *bind.WatchOpts, sink chan<- *FederationsyncNodeChanged, nodeIndex []uint8, changedBy []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var nodeIndexRule []interface{}
	for _, nodeIndexItem := range nodeIndex {
		nodeIndexRule = append(nodeIndexRule, nodeIndexItem)
	}
	var changedByRule []interface{}
	for _, changedByItem := range changedBy {
		changedByRule = append(changedByRule, changedByItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Federationsync.contract.WatchLogs(opts, "NodeChanged", nodeIndexRule, changedByRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FederationsyncNodeChanged)
				if err := _Federationsync.contract.UnpackLog(event, "NodeChanged", log); err != nil {
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

// ParseNodeChanged is a log parse operation binding the contract event 0xf5becf5ef6bf281ea3c4f617c266a872c7662757c90aa5f5cdffe02bd2023975.
//
// Solidity: event NodeChanged(uint8 indexed nodeIndex, address indexed changedBy, address indexed newAddress)
func (_Federationsync *FederationsyncFilterer) ParseNodeChanged(log types.Log) (*FederationsyncNodeChanged, error) {
	event := new(FederationsyncNodeChanged)
	if err := _Federationsync.contract.UnpackLog(event, "NodeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FederationsyncRequestConfirmedIterator is returned from FilterRequestConfirmed and is used to iterate over the raw logs and unpacked data for RequestConfirmed events raised by the Federationsync contract.
type FederationsyncRequestConfirmedIterator struct {
	Event *FederationsyncRequestConfirmed // Event containing the contract specifics and raw log

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
func (it *FederationsyncRequestConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FederationsyncRequestConfirmed)
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
		it.Event = new(FederationsyncRequestConfirmed)
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
func (it *FederationsyncRequestConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FederationsyncRequestConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FederationsyncRequestConfirmed represents a RequestConfirmed event raised by the Federationsync contract.
type FederationsyncRequestConfirmed struct {
	ConfirmedBy common.Address
	Recipient   common.Address
	Amount      *big.Int
	Nonce       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequestConfirmed is a free log retrieval operation binding the contract event 0x6718bbbc3a1333e921501591df5fe974bd1de1e3dc5498203221a126967067b8.
//
// Solidity: event RequestConfirmed(address indexed confirmedBy, address indexed recipient, uint256 amount, uint256 nonce)
func (_Federationsync *FederationsyncFilterer) FilterRequestConfirmed(opts *bind.FilterOpts, confirmedBy []common.Address, recipient []common.Address) (*FederationsyncRequestConfirmedIterator, error) {

	var confirmedByRule []interface{}
	for _, confirmedByItem := range confirmedBy {
		confirmedByRule = append(confirmedByRule, confirmedByItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Federationsync.contract.FilterLogs(opts, "RequestConfirmed", confirmedByRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &FederationsyncRequestConfirmedIterator{contract: _Federationsync.contract, event: "RequestConfirmed", logs: logs, sub: sub}, nil
}

// WatchRequestConfirmed is a free log subscription operation binding the contract event 0x6718bbbc3a1333e921501591df5fe974bd1de1e3dc5498203221a126967067b8.
//
// Solidity: event RequestConfirmed(address indexed confirmedBy, address indexed recipient, uint256 amount, uint256 nonce)
func (_Federationsync *FederationsyncFilterer) WatchRequestConfirmed(opts *bind.WatchOpts, sink chan<- *FederationsyncRequestConfirmed, confirmedBy []common.Address, recipient []common.Address) (event.Subscription, error) {

	var confirmedByRule []interface{}
	for _, confirmedByItem := range confirmedBy {
		confirmedByRule = append(confirmedByRule, confirmedByItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Federationsync.contract.WatchLogs(opts, "RequestConfirmed", confirmedByRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FederationsyncRequestConfirmed)
				if err := _Federationsync.contract.UnpackLog(event, "RequestConfirmed", log); err != nil {
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

// ParseRequestConfirmed is a log parse operation binding the contract event 0x6718bbbc3a1333e921501591df5fe974bd1de1e3dc5498203221a126967067b8.
//
// Solidity: event RequestConfirmed(address indexed confirmedBy, address indexed recipient, uint256 amount, uint256 nonce)
func (_Federationsync *FederationsyncFilterer) ParseRequestConfirmed(log types.Log) (*FederationsyncRequestConfirmed, error) {
	event := new(FederationsyncRequestConfirmed)
	if err := _Federationsync.contract.UnpackLog(event, "RequestConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FederationsyncTransferExecutedIterator is returned from FilterTransferExecuted and is used to iterate over the raw logs and unpacked data for TransferExecuted events raised by the Federationsync contract.
type FederationsyncTransferExecutedIterator struct {
	Event *FederationsyncTransferExecuted // Event containing the contract specifics and raw log

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
func (it *FederationsyncTransferExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FederationsyncTransferExecuted)
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
		it.Event = new(FederationsyncTransferExecuted)
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
func (it *FederationsyncTransferExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FederationsyncTransferExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FederationsyncTransferExecuted represents a TransferExecuted event raised by the Federationsync contract.
type FederationsyncTransferExecuted struct {
	Recipient common.Address
	Amount    *big.Int
	Nonce     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransferExecuted is a free log retrieval operation binding the contract event 0xf8d62a11d316e69c1460662f86b7046003826db86fdc0d5ffa8c3adaa24f0a58.
//
// Solidity: event TransferExecuted(address indexed recipient, uint256 amount, uint256 nonce)
func (_Federationsync *FederationsyncFilterer) FilterTransferExecuted(opts *bind.FilterOpts, recipient []common.Address) (*FederationsyncTransferExecutedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Federationsync.contract.FilterLogs(opts, "TransferExecuted", recipientRule)
	if err != nil {
		return nil, err
	}
	return &FederationsyncTransferExecutedIterator{contract: _Federationsync.contract, event: "TransferExecuted", logs: logs, sub: sub}, nil
}

// WatchTransferExecuted is a free log subscription operation binding the contract event 0xf8d62a11d316e69c1460662f86b7046003826db86fdc0d5ffa8c3adaa24f0a58.
//
// Solidity: event TransferExecuted(address indexed recipient, uint256 amount, uint256 nonce)
func (_Federationsync *FederationsyncFilterer) WatchTransferExecuted(opts *bind.WatchOpts, sink chan<- *FederationsyncTransferExecuted, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Federationsync.contract.WatchLogs(opts, "TransferExecuted", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FederationsyncTransferExecuted)
				if err := _Federationsync.contract.UnpackLog(event, "TransferExecuted", log); err != nil {
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

// ParseTransferExecuted is a log parse operation binding the contract event 0xf8d62a11d316e69c1460662f86b7046003826db86fdc0d5ffa8c3adaa24f0a58.
//
// Solidity: event TransferExecuted(address indexed recipient, uint256 amount, uint256 nonce)
func (_Federationsync *FederationsyncFilterer) ParseTransferExecuted(log types.Log) (*FederationsyncTransferExecuted, error) {
	event := new(FederationsyncTransferExecuted)
	if err := _Federationsync.contract.UnpackLog(event, "TransferExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
