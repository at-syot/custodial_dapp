// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todo

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

// TodoTask is an auto generated low-level Go binding around an user-defined struct.
type TodoTask struct {
	Content string
	Status  bool
}

// TodoMetaData contains all meta data concerning the Todo contract.
var TodoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"get\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"list\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"internalType\":\"structTodo.Task[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50610ca98061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630f560cd71461005c5780634cc822151461007a5780634e324847146100965780639507d39a146100b2578063b0c8f9dc146100e2575b600080fd5b6100646100fe565b6040516100719190610655565b60405180910390f35b610094600480360381019061008f91906106c1565b61020a565b005b6100b060048036038101906100ab919061084f565b610256565b005b6100cc60048036038101906100c791906106c1565b6102cb565b6040516100d991906108fb565b60405180910390f35b6100fc60048036038101906100f7919061091d565b6103b4565b005b60606000805480602002602001604051908101604052809291908181526020016000905b82821015610201578382906000526020600020906002020160405180604001604052908160008201805461015590610995565b80601f016020809104026020016040519081016040528092919081815260200182805461018190610995565b80156101ce5780601f106101a3576101008083540402835291602001916101ce565b820191906000526020600020905b8154815290600101906020018083116101b157829003601f168201915b505050505081526020016001820160009054906101000a900460ff16151515158152505081526020019060010190610122565b50505050905090565b6000818154811061021e5761021d6109c6565b5b90600052602060002090600202016000808201600061023d9190610432565b6001820160006101000a81549060ff0219169055505050565b816000848154811061026b5761026a6109c6565b5b906000526020600020906002020160000190816102889190610ba1565b50806000848154811061029e5761029d6109c6565b5b906000526020600020906002020160010160006101000a81548160ff021916908315150217905550505050565b6102d3610472565b600082815481106102e7576102e66109c6565b5b906000526020600020906002020160405180604001604052908160008201805461031090610995565b80601f016020809104026020016040519081016040528092919081815260200182805461033c90610995565b80156103895780601f1061035e57610100808354040283529160200191610389565b820191906000526020600020905b81548152906001019060200180831161036c57829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815250509050919050565b60006040518060400160405280838152602001600015158152509080600181540180825580915050600190039060005260206000209060020201600090919091909150600082015181600001908161040c9190610ba1565b5060208201518160010160006101000a81548160ff021916908315150217905550505050565b50805461043e90610995565b6000825580601f10610450575061046f565b601f01602090049060005260206000209081019061046e919061048e565b5b50565b6040518060400160405280606081526020016000151581525090565b5b808211156104a757600081600090555060010161048f565b5090565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156105115780820151818401526020810190506104f6565b60008484015250505050565b6000601f19601f8301169050919050565b6000610539826104d7565b61054381856104e2565b93506105538185602086016104f3565b61055c8161051d565b840191505092915050565b60008115159050919050565b61057c81610567565b82525050565b6000604083016000830151848203600086015261059f828261052e565b91505060208301516105b46020860182610573565b508091505092915050565b60006105cb8383610582565b905092915050565b6000602082019050919050565b60006105eb826104ab565b6105f581856104b6565b935083602082028501610607856104c7565b8060005b85811015610643578484038952815161062485826105bf565b945061062f836105d3565b925060208a0199505060018101905061060b565b50829750879550505050505092915050565b6000602082019050818103600083015261066f81846105e0565b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b61069e8161068b565b81146106a957600080fd5b50565b6000813590506106bb81610695565b92915050565b6000602082840312156106d7576106d6610681565b5b60006106e5848285016106ac565b91505092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6107308261051d565b810181811067ffffffffffffffff8211171561074f5761074e6106f8565b5b80604052505050565b6000610762610677565b905061076e8282610727565b919050565b600067ffffffffffffffff82111561078e5761078d6106f8565b5b6107978261051d565b9050602081019050919050565b82818337600083830152505050565b60006107c66107c184610773565b610758565b9050828152602081018484840111156107e2576107e16106f3565b5b6107ed8482856107a4565b509392505050565b600082601f83011261080a576108096106ee565b5b813561081a8482602086016107b3565b91505092915050565b61082c81610567565b811461083757600080fd5b50565b60008135905061084981610823565b92915050565b60008060006060848603121561086857610867610681565b5b6000610876868287016106ac565b935050602084013567ffffffffffffffff81111561089757610896610686565b5b6108a3868287016107f5565b92505060406108b48682870161083a565b9150509250925092565b600060408301600083015184820360008601526108db828261052e565b91505060208301516108f06020860182610573565b508091505092915050565b6000602082019050818103600083015261091581846108be565b905092915050565b60006020828403121561093357610932610681565b5b600082013567ffffffffffffffff81111561095157610950610686565b5b61095d848285016107f5565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806109ad57607f821691505b6020821081036109c0576109bf610966565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610a577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a1a565b610a618683610a1a565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610a9e610a99610a948461068b565b610a79565b61068b565b9050919050565b6000819050919050565b610ab883610a83565b610acc610ac482610aa5565b848454610a27565b825550505050565b600090565b610ae1610ad4565b610aec818484610aaf565b505050565b5b81811015610b1057610b05600082610ad9565b600181019050610af2565b5050565b601f821115610b5557610b26816109f5565b610b2f84610a0a565b81016020851015610b3e578190505b610b52610b4a85610a0a565b830182610af1565b50505b505050565b600082821c905092915050565b6000610b7860001984600802610b5a565b1980831691505092915050565b6000610b918383610b67565b9150826002028217905092915050565b610baa826104d7565b67ffffffffffffffff811115610bc357610bc26106f8565b5b610bcd8254610995565b610bd8828285610b14565b600060209050601f831160018114610c0b5760008415610bf9578287015190505b610c038582610b85565b865550610c6b565b601f198416610c19866109f5565b60005b82811015610c4157848901518255600182019150602085019450602081019050610c1c565b86831015610c5e5784890151610c5a601f891682610b67565b8355505b6001600288020188555050505b50505050505056fea264697066735822122051e104d8480594eb0c235537aca0c233bc5d2d95e35d883032cb3d148cc2b1e064736f6c634300081c0033",
}

// TodoABI is the input ABI used to generate the binding from.
// Deprecated: Use TodoMetaData.ABI instead.
var TodoABI = TodoMetaData.ABI

// TodoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TodoMetaData.Bin instead.
var TodoBin = TodoMetaData.Bin

// DeployTodo deploys a new Ethereum contract, binding an instance of Todo to it.
func DeployTodo(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Todo, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TodoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// Todo is an auto generated Go binding around an Ethereum contract.
type Todo struct {
	TodoCaller     // Read-only binding to the contract
	TodoTransactor // Write-only binding to the contract
	TodoFilterer   // Log filterer for contract events
}

// TodoCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodoSession struct {
	Contract     *Todo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodoCallerSession struct {
	Contract *TodoCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TodoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodoTransactorSession struct {
	Contract     *TodoTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodoRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodoRaw struct {
	Contract *Todo // Generic contract binding to access the raw methods on
}

// TodoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodoCallerRaw struct {
	Contract *TodoCaller // Generic read-only contract binding to access the raw methods on
}

// TodoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodoTransactorRaw struct {
	Contract *TodoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodo creates a new instance of Todo, bound to a specific deployed contract.
func NewTodo(address common.Address, backend bind.ContractBackend) (*Todo, error) {
	contract, err := bindTodo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todo{TodoCaller: TodoCaller{contract: contract}, TodoTransactor: TodoTransactor{contract: contract}, TodoFilterer: TodoFilterer{contract: contract}}, nil
}

// NewTodoCaller creates a new read-only instance of Todo, bound to a specific deployed contract.
func NewTodoCaller(address common.Address, caller bind.ContractCaller) (*TodoCaller, error) {
	contract, err := bindTodo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodoCaller{contract: contract}, nil
}

// NewTodoTransactor creates a new write-only instance of Todo, bound to a specific deployed contract.
func NewTodoTransactor(address common.Address, transactor bind.ContractTransactor) (*TodoTransactor, error) {
	contract, err := bindTodo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodoTransactor{contract: contract}, nil
}

// NewTodoFilterer creates a new log filterer instance of Todo, bound to a specific deployed contract.
func NewTodoFilterer(address common.Address, filterer bind.ContractFilterer) (*TodoFilterer, error) {
	contract, err := bindTodo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodoFilterer{contract: contract}, nil
}

// bindTodo binds a generic wrapper to an already deployed contract.
func bindTodo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TodoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.TodoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.TodoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todo *TodoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todo *TodoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todo *TodoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todo.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 id) view returns((string,bool))
func (_Todo *TodoCaller) Get(opts *bind.CallOpts, id *big.Int) (TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "get", id)

	if err != nil {
		return *new(TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new(TodoTask)).(*TodoTask)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 id) view returns((string,bool))
func (_Todo *TodoSession) Get(id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, id)
}

// Get is a free data retrieval call binding the contract method 0x9507d39a.
//
// Solidity: function get(uint256 id) view returns((string,bool))
func (_Todo *TodoCallerSession) Get(id *big.Int) (TodoTask, error) {
	return _Todo.Contract.Get(&_Todo.CallOpts, id)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCaller) List(opts *bind.CallOpts) ([]TodoTask, error) {
	var out []interface{}
	err := _Todo.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]TodoTask), err
	}

	out0 := *abi.ConvertType(out[0], new([]TodoTask)).(*[]TodoTask)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((string,bool)[])
func (_Todo *TodoCallerSession) List() ([]TodoTask, error) {
	return _Todo.Contract.List(&_Todo.CallOpts)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string content) returns()
func (_Todo *TodoTransactor) Add(opts *bind.TransactOpts, content string) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "add", content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string content) returns()
func (_Todo *TodoSession) Add(content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, content)
}

// Add is a paid mutator transaction binding the contract method 0xb0c8f9dc.
//
// Solidity: function add(string content) returns()
func (_Todo *TodoTransactorSession) Add(content string) (*types.Transaction, error) {
	return _Todo.Contract.Add(&_Todo.TransactOpts, content)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 id) returns()
func (_Todo *TodoTransactor) Remove(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "remove", id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 id) returns()
func (_Todo *TodoSession) Remove(id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, id)
}

// Remove is a paid mutator transaction binding the contract method 0x4cc82215.
//
// Solidity: function remove(uint256 id) returns()
func (_Todo *TodoTransactorSession) Remove(id *big.Int) (*types.Transaction, error) {
	return _Todo.Contract.Remove(&_Todo.TransactOpts, id)
}

// Update is a paid mutator transaction binding the contract method 0x4e324847.
//
// Solidity: function update(uint256 id, string content, bool status) returns()
func (_Todo *TodoTransactor) Update(opts *bind.TransactOpts, id *big.Int, content string, status bool) (*types.Transaction, error) {
	return _Todo.contract.Transact(opts, "update", id, content, status)
}

// Update is a paid mutator transaction binding the contract method 0x4e324847.
//
// Solidity: function update(uint256 id, string content, bool status) returns()
func (_Todo *TodoSession) Update(id *big.Int, content string, status bool) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, id, content, status)
}

// Update is a paid mutator transaction binding the contract method 0x4e324847.
//
// Solidity: function update(uint256 id, string content, bool status) returns()
func (_Todo *TodoTransactorSession) Update(id *big.Int, content string, status bool) (*types.Transaction, error) {
	return _Todo.Contract.Update(&_Todo.TransactOpts, id, content, status)
}
