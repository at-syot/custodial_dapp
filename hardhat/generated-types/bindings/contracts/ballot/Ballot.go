// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ballot

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

// BallotMetaData contains all meta data concerning the Ballot contract.
var BallotMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proposalNames\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposal\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"voted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winningName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"winningName_\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516115193803806115198339818101604052810190610032919061030e565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018060008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555060005b8151811015610164576002604051806040016040528084848151811061010657610105610357565b5b60200260200101518152602001600081525090806001815401808255809150506001900390600052602060002090600202016000909190919091506000820151816000015560208201518160010155505080806001019150506100dd565b5050610386565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101cd82610184565b810181811067ffffffffffffffff821117156101ec576101eb610195565b5b80604052505050565b60006101ff61016b565b905061020b82826101c4565b919050565b600067ffffffffffffffff82111561022b5761022a610195565b5b602082029050602081019050919050565b600080fd5b6000819050919050565b61025481610241565b811461025f57600080fd5b50565b6000815190506102718161024b565b92915050565b600061028a61028584610210565b6101f5565b905080838252602082019050602084028301858111156102ad576102ac61023c565b5b835b818110156102d657806102c28882610262565b8452602084019350506020810190506102af565b5050509392505050565b600082601f8301126102f5576102f461017f565b5b8151610305848260208601610277565b91505092915050565b60006020828403121561032457610323610175565b5b600082015167ffffffffffffffff8111156103425761034161017a565b5b61034e848285016102e0565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b611184806103956000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80632e4176cf1161005b5780632e4176cf146100ed5780635c19a95c1461010b5780639e7b8d6114610127578063a3ec138d146101435761007d565b80630121b93f14610082578063013cf08b1461009e5780631f9ff6d7146100cf575b600080fd5b61009c60048036038101906100979190610a79565b610176565b005b6100b860048036038101906100b39190610a79565b6102bd565b6040516100c6929190610ace565b60405180910390f35b6100d76102f1565b6040516100e49190610af7565b60405180910390f35b6100f5610325565b6040516101029190610b53565b60405180910390f35b61012560048036038101906101209190610b9a565b610349565b005b610141600480360381019061013c9190610b9a565b610772565b005b61015d60048036038101906101589190610b9a565b61095f565b60405161016d9493929190610be2565b60405180910390f35b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600181600001541015610201576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f890610c84565b60405180910390fd5b8060010160009054906101000a900460ff1615610253576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024a90610cf0565b60405180910390fd5b60018160010160006101000a81548160ff02191690831515021790555081816002018190555080600001546002838154811061029257610291610d10565b5b906000526020600020906002020160010160008282546102b29190610d6e565b925050819055505050565b600281815481106102cd57600080fd5b90600052602060002090600202016000915090508060000154908060010154905082565b600060026102fd6109bc565b8154811061030e5761030d610d10565b5b906000526020600020906002020160000154905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160000154036103d3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103ca90610dee565b60405180910390fd5b8060010160009054906101000a900460ff1615610425576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041c90610e5a565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1603610493576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048a90610ec6565b60405180910390fd5b5b600073ffffffffffffffffffffffffffffffffffffffff16600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461060257600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691503373ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036105fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f490610f32565b60405180910390fd5b610494565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060018160000154101561068d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068490610fc4565b60405180910390fd5b60018260010160006101000a81548160ff021916908315150217905550828260010160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060010160009054906101000a900460ff161561074d578160000154600282600201548154811061072157610720610d10565b5b906000526020600020906002020160010160008282546107419190610d6e565b9250508190555061076d565b81600001548160000160008282546107659190610d6e565b925050819055505b505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610800576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107f790611056565b60405180910390fd5b600160008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900460ff1615610890576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610887906110c2565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000015414610915576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161090c9061112e565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000018190555050565b60016020528060005260406000206000915090508060000154908060010160009054906101000a900460ff16908060010160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020154905084565b6000806000905060005b600280549050811015610a395781600282815481106109e8576109e7610d10565b5b9060005260206000209060020201600101541115610a2c5760028181548110610a1457610a13610d10565b5b90600052602060002090600202016001015491508092505b80806001019150506109c6565b505090565b600080fd5b6000819050919050565b610a5681610a43565b8114610a6157600080fd5b50565b600081359050610a7381610a4d565b92915050565b600060208284031215610a8f57610a8e610a3e565b5b6000610a9d84828501610a64565b91505092915050565b6000819050919050565b610ab981610aa6565b82525050565b610ac881610a43565b82525050565b6000604082019050610ae36000830185610ab0565b610af06020830184610abf565b9392505050565b6000602082019050610b0c6000830184610ab0565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b3d82610b12565b9050919050565b610b4d81610b32565b82525050565b6000602082019050610b686000830184610b44565b92915050565b610b7781610b32565b8114610b8257600080fd5b50565b600081359050610b9481610b6e565b92915050565b600060208284031215610bb057610baf610a3e565b5b6000610bbe84828501610b85565b91505092915050565b60008115159050919050565b610bdc81610bc7565b82525050565b6000608082019050610bf76000830187610abf565b610c046020830186610bd3565b610c116040830185610b44565b610c1e6060830184610abf565b95945050505050565b600082825260208201905092915050565b7f596f7520646f6e2774206861766520766f74696e672072696768742e00000000600082015250565b6000610c6e601c83610c27565b9150610c7982610c38565b602082019050919050565b60006020820190508181036000830152610c9d81610c61565b9050919050565b7f596f7520616c726561647920766f7465642e0000000000000000000000000000600082015250565b6000610cda601283610c27565b9150610ce582610ca4565b602082019050919050565b60006020820190508181036000830152610d0981610ccd565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610d7982610a43565b9150610d8483610a43565b9250828201905080821115610d9c57610d9b610d3f565b5b92915050565b7f73656e6465723a596f75206e6f7420616c6c6f7720746f20766f746500000000600082015250565b6000610dd8601c83610c27565b9150610de382610da2565b602082019050919050565b60006020820190508181036000830152610e0781610dcb565b9050919050565b7f73656e6465723a596f7520616c726561647920766f7465640000000000000000600082015250565b6000610e44601883610c27565b9150610e4f82610e0e565b602082019050919050565b60006020820190508181036000830152610e7381610e37565b9050919050565b7f53656c662d64656c65676174696f6e206973206e6f7420616c6c6f7700000000600082015250565b6000610eb0601c83610c27565b9150610ebb82610e7a565b602082019050919050565b60006020820190508181036000830152610edf81610ea3565b9050919050565b7f576520666f756e6420612064656c65676174696f6e206c6f6f70000000000000600082015250565b6000610f1c601a83610c27565b9150610f2782610ee6565b602082019050919050565b60006020820190508181036000830152610f4b81610f0f565b9050919050565b7f64656c61676174656420766f7465723a5468697320766f74657220686176652060008201527f6e6f20726967687420746f20766f746500000000000000000000000000000000602082015250565b6000610fae603083610c27565b9150610fb982610f52565b604082019050919050565b60006020820190508181036000830152610fdd81610fa1565b9050919050565b7f6f6e6c79206368616972706572736f6e2063616e20676976652074686520726960008201527f67687420746f20766f74652e0000000000000000000000000000000000000000602082015250565b6000611040602c83610c27565b915061104b82610fe4565b604082019050919050565b6000602082019050818103600083015261106f81611033565b9050919050565b7f54686520766f74657220616c726561647920766f7465642e0000000000000000600082015250565b60006110ac601883610c27565b91506110b782611076565b602082019050919050565b600060208201905081810360008301526110db8161109f565b9050919050565b7f4e6f206e65656420746f206769766520726967687420746f20766f7465720000600082015250565b6000611118601e83610c27565b9150611123826110e2565b602082019050919050565b600060208201905081810360008301526111478161110b565b905091905056fea26469706673582212207692354ef45e478da665ea714bea06086ebd3e686916b3e086f7b73ce96bd39f64736f6c634300081c0033",
}

// BallotABI is the input ABI used to generate the binding from.
// Deprecated: Use BallotMetaData.ABI instead.
var BallotABI = BallotMetaData.ABI

// BallotBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BallotMetaData.Bin instead.
var BallotBin = BallotMetaData.Bin

// DeployBallot deploys a new Ethereum contract, binding an instance of Ballot to it.
func DeployBallot(auth *bind.TransactOpts, backend bind.ContractBackend, proposalNames [][32]byte) (common.Address, *types.Transaction, *Ballot, error) {
	parsed, err := BallotMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BallotBin), backend, proposalNames)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ballot{BallotCaller: BallotCaller{contract: contract}, BallotTransactor: BallotTransactor{contract: contract}, BallotFilterer: BallotFilterer{contract: contract}}, nil
}

// Ballot is an auto generated Go binding around an Ethereum contract.
type Ballot struct {
	BallotCaller     // Read-only binding to the contract
	BallotTransactor // Write-only binding to the contract
	BallotFilterer   // Log filterer for contract events
}

// BallotCaller is an auto generated read-only Go binding around an Ethereum contract.
type BallotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BallotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BallotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BallotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BallotSession struct {
	Contract     *Ballot           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BallotCallerSession struct {
	Contract *BallotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BallotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BallotTransactorSession struct {
	Contract     *BallotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BallotRaw is an auto generated low-level Go binding around an Ethereum contract.
type BallotRaw struct {
	Contract *Ballot // Generic contract binding to access the raw methods on
}

// BallotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BallotCallerRaw struct {
	Contract *BallotCaller // Generic read-only contract binding to access the raw methods on
}

// BallotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BallotTransactorRaw struct {
	Contract *BallotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBallot creates a new instance of Ballot, bound to a specific deployed contract.
func NewBallot(address common.Address, backend bind.ContractBackend) (*Ballot, error) {
	contract, err := bindBallot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ballot{BallotCaller: BallotCaller{contract: contract}, BallotTransactor: BallotTransactor{contract: contract}, BallotFilterer: BallotFilterer{contract: contract}}, nil
}

// NewBallotCaller creates a new read-only instance of Ballot, bound to a specific deployed contract.
func NewBallotCaller(address common.Address, caller bind.ContractCaller) (*BallotCaller, error) {
	contract, err := bindBallot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BallotCaller{contract: contract}, nil
}

// NewBallotTransactor creates a new write-only instance of Ballot, bound to a specific deployed contract.
func NewBallotTransactor(address common.Address, transactor bind.ContractTransactor) (*BallotTransactor, error) {
	contract, err := bindBallot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BallotTransactor{contract: contract}, nil
}

// NewBallotFilterer creates a new log filterer instance of Ballot, bound to a specific deployed contract.
func NewBallotFilterer(address common.Address, filterer bind.ContractFilterer) (*BallotFilterer, error) {
	contract, err := bindBallot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BallotFilterer{contract: contract}, nil
}

// bindBallot binds a generic wrapper to an already deployed contract.
func bindBallot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BallotMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ballot *BallotRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ballot.Contract.BallotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ballot *BallotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.Contract.BallotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ballot *BallotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ballot.Contract.BallotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ballot *BallotCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ballot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ballot *BallotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ballot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ballot *BallotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ballot.Contract.contract.Transact(opts, method, params...)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Ballot *BallotCaller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "chairperson")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Ballot *BallotSession) Chairperson() (common.Address, error) {
	return _Ballot.Contract.Chairperson(&_Ballot.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Ballot *BallotCallerSession) Chairperson() (common.Address, error) {
	return _Ballot.Contract.Chairperson(&_Ballot.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Ballot *BallotCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Name      [32]byte
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Ballot *BallotSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Ballot.Contract.Proposals(&_Ballot.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount)
func (_Ballot *BallotCallerSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
}, error) {
	return _Ballot.Contract.Proposals(&_Ballot.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Ballot *BallotCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		Weight   *big.Int
		Voted    bool
		Delegate common.Address
		Vote     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Voted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Delegate = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Vote = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Ballot *BallotSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Ballot.Contract.Voters(&_Ballot.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Ballot *BallotCallerSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Ballot.Contract.Voters(&_Ballot.CallOpts, arg0)
}

// WinningName is a free data retrieval call binding the contract method 0x1f9ff6d7.
//
// Solidity: function winningName() view returns(bytes32 winningName_)
func (_Ballot *BallotCaller) WinningName(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ballot.contract.Call(opts, &out, "winningName")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinningName is a free data retrieval call binding the contract method 0x1f9ff6d7.
//
// Solidity: function winningName() view returns(bytes32 winningName_)
func (_Ballot *BallotSession) WinningName() ([32]byte, error) {
	return _Ballot.Contract.WinningName(&_Ballot.CallOpts)
}

// WinningName is a free data retrieval call binding the contract method 0x1f9ff6d7.
//
// Solidity: function winningName() view returns(bytes32 winningName_)
func (_Ballot *BallotCallerSession) WinningName() ([32]byte, error) {
	return _Ballot.Contract.WinningName(&_Ballot.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Ballot *BallotTransactor) Delegate(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "delegate", to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Ballot *BallotSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.Delegate(&_Ballot.TransactOpts, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Ballot *BallotTransactorSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.Delegate(&_Ballot.TransactOpts, to)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Ballot *BallotTransactor) GiveRightToVote(opts *bind.TransactOpts, voter common.Address) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "giveRightToVote", voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Ballot *BallotSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.GiveRightToVote(&_Ballot.TransactOpts, voter)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0x9e7b8d61.
//
// Solidity: function giveRightToVote(address voter) returns()
func (_Ballot *BallotTransactorSession) GiveRightToVote(voter common.Address) (*types.Transaction, error) {
	return _Ballot.Contract.GiveRightToVote(&_Ballot.TransactOpts, voter)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Ballot *BallotTransactor) Vote(opts *bind.TransactOpts, proposal *big.Int) (*types.Transaction, error) {
	return _Ballot.contract.Transact(opts, "vote", proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Ballot *BallotSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Vote(&_Ballot.TransactOpts, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Ballot *BallotTransactorSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Ballot.Contract.Vote(&_Ballot.TransactOpts, proposal)
}
