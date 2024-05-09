// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dapp

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
)

// DappMetaData contains all meta data concerning the Dapp contract.
var DappMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"say\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506101688061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610029575f3560e01c8063954ab4b21461002d575b5f80fd5b61003561004b565b6040516100429190610112565b60405180910390f35b60606040518060400160405280601081526020017f68656c6c6f206574686572776f726c6400000000000000000000000000000000815250905090565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156100bf5780820151818401526020810190506100a4565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6100e482610088565b6100ee8185610092565b93506100fe8185602086016100a2565b610107816100ca565b840191505092915050565b5f6020820190508181035f83015261012a81846100da565b90509291505056fea26469706673582212201cf24125e893df2ce791b7f5ee770fee82a370cce47bf855bd91c79c4a570e0c64736f6c63430008150033",
}

// DappABI is the input ABI used to generate the binding from.
// Deprecated: Use DappMetaData.ABI instead.
var DappABI = DappMetaData.ABI

// DappBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DappMetaData.Bin instead.
var DappBin = DappMetaData.Bin

// DeployDapp deploys a new Ethereum contract, binding an instance of Dapp to it.
func DeployDapp(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dapp, error) {
	parsed, err := DappMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DappBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dapp{DappCaller: DappCaller{contract: contract}, DappTransactor: DappTransactor{contract: contract}, DappFilterer: DappFilterer{contract: contract}}, nil
}

// Dapp is an auto generated Go binding around an Ethereum contract.
type Dapp struct {
	DappCaller     // Read-only binding to the contract
	DappTransactor // Write-only binding to the contract
	DappFilterer   // Log filterer for contract events
}

// DappCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappSession struct {
	Contract     *Dapp             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappCallerSession struct {
	Contract *DappCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DappTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappTransactorSession struct {
	Contract     *DappTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappRaw struct {
	Contract *Dapp // Generic contract binding to access the raw methods on
}

// DappCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappCallerRaw struct {
	Contract *DappCaller // Generic read-only contract binding to access the raw methods on
}

// DappTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappTransactorRaw struct {
	Contract *DappTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDapp creates a new instance of Dapp, bound to a specific deployed contract.
func NewDapp(address common.Address, backend bind.ContractBackend) (*Dapp, error) {
	contract, err := bindDapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dapp{DappCaller: DappCaller{contract: contract}, DappTransactor: DappTransactor{contract: contract}, DappFilterer: DappFilterer{contract: contract}}, nil
}

// NewDappCaller creates a new read-only instance of Dapp, bound to a specific deployed contract.
func NewDappCaller(address common.Address, caller bind.ContractCaller) (*DappCaller, error) {
	contract, err := bindDapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappCaller{contract: contract}, nil
}

// NewDappTransactor creates a new write-only instance of Dapp, bound to a specific deployed contract.
func NewDappTransactor(address common.Address, transactor bind.ContractTransactor) (*DappTransactor, error) {
	contract, err := bindDapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappTransactor{contract: contract}, nil
}

// NewDappFilterer creates a new log filterer instance of Dapp, bound to a specific deployed contract.
func NewDappFilterer(address common.Address, filterer bind.ContractFilterer) (*DappFilterer, error) {
	contract, err := bindDapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappFilterer{contract: contract}, nil
}

// bindDapp binds a generic wrapper to an already deployed contract.
func bindDapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dapp *DappRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dapp.Contract.DappCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dapp *DappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dapp.Contract.DappTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dapp *DappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dapp.Contract.DappTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dapp *DappCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dapp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dapp *DappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dapp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dapp *DappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dapp.Contract.contract.Transact(opts, method, params...)
}

// Say is a free data retrieval call binding the contract method 0x954ab4b2.
//
// Solidity: function say() pure returns(string)
func (_Dapp *DappCaller) Say(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dapp.contract.Call(opts, &out, "say")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Say is a free data retrieval call binding the contract method 0x954ab4b2.
//
// Solidity: function say() pure returns(string)
func (_Dapp *DappSession) Say() (string, error) {
	return _Dapp.Contract.Say(&_Dapp.CallOpts)
}

// Say is a free data retrieval call binding the contract method 0x954ab4b2.
//
// Solidity: function say() pure returns(string)
func (_Dapp *DappCallerSession) Say() (string, error) {
	return _Dapp.Contract.Say(&_Dapp.CallOpts)
}
