// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gocontracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// CalypsoABI is the input ABI used to generate the binding from.
const CalypsoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addWriteRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"ownersMap\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"canWrite\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"checkIfRequestIsValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"isValidReadRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddReadRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"holder\",\"type\":\"address\"},{\"name\":\"rrholder\",\"type\":\"address\"},{\"name\":\"o\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// CalypsoBin is the compiled bytecode used for deploying new contracts.
const CalypsoBin = `0x608060405234801561001057600080fd5b5060405160608061090383398101604090815281516020830151919092015160008054600160a060020a0319908116600160a060020a039384161782556001805482169584169590951790945560038054909416919092161790915561088790819061007c90396000f3006080604052600436106100695763ffffffff60e060020a6000350416631029a017811461006e57806319a444b7146100915780633239b71b146100c65780634a4e19d3146100e7578063778ad09c14610108578063ac1e9ef414610129578063bdae9b891461014a575b600080fd5b34801561007a57600080fd5b5061008f600160a060020a036004351661016b565b005b34801561009d57600080fd5b506100b2600160a060020a036004351661024c565b604080519115158252519081900360200190f35b3480156100d257600080fd5b506100b2600160a060020a0360043516610261565b3480156100f357600080fd5b506100b2600160a060020a036004351661027f565b34801561011457600080fd5b506100b2600160a060020a036004351661031c565b34801561013557600080fd5b5061008f600160a060020a0360043516610631565b34801561015657600080fd5b5061008f600160a060020a0360043516610655565b61017433610261565b15156101ca576040805160e560020a62461bcd02815260206004820152601e60248201527f596f7520617265206e6f7420612072656769737465726564206f776e65720000604482015290519081900360640190fd5b600154604080517f1029a017000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291519190921691631029a01791602480830192600092919082900301818387803b15801561023157600080fd5b505af1158015610245573d6000803e3d6000fd5b5050505050565b60046020526000908152604090205460ff1681565b600160a060020a031660009081526004602052604090205460ff1690565b600154604080517f42087d4f000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916342087d4f9160248082019260209290919082900301818787803b1580156102ea57600080fd5b505af11580156102fe573d6000803e3d6000fd5b505050506040513d602081101561031457600080fd5b505192915050565b600354604080517ffd6336dc000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151600093849384938493919092169163fd6336dc91602480830192602092919082900301818787803b15801561038c57600080fd5b505af11580156103a0573d6000803e3d6000fd5b505050506040513d60208110156103b657600080fd5b50511515610434576040805160e560020a62461bcd02815260206004820152602960248201527f4974206d75737420617420666972737420626520612076616c6964207265616460448201527f20726571756573742e0000000000000000000000000000000000000000000000606482015290519081900360840190fd5b84925082600160a060020a03166364e94e676040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561047557600080fd5b505af1158015610489573d6000803e3d6000fd5b505050506040513d602081101561049f57600080fd5b5051604080517f8da5cb5b0000000000000000000000000000000000000000000000000000000081529051919350839250600160a060020a038084169263565542569291871691638da5cb5b9160048083019260209291908290030181600087803b15801561050d57600080fd5b505af1158015610521573d6000803e3d6000fd5b505050506040513d602081101561053757600080fd5b50516040805160e060020a63ffffffff8516028152600160a060020a0390921660048301525160248083019260209291908290030181600087803b15801561057e57600080fd5b505af1158015610592573d6000803e3d6000fd5b505050506040513d60208110156105a857600080fd5b50511515610626576040805160e560020a62461bcd02815260206004820152603a60248201527f497320746869732061206d656d626572206f662074686520706f6c696379206f60448201527f662074686520706f696e74656420577269746552657175657374000000000000606482015290519081900360840190fd5b506001949350505050565b600160a060020a03166000908152600460205260409020805460ff19166001179055565b60008082915081600160a060020a03166364e94e676040518163ffffffff1660e060020a028152600401602060405180830381600087803b15801561069957600080fd5b505af11580156106ad573d6000803e3d6000fd5b505050506040513d60208110156106c357600080fd5b5051600154604080517f42087d4f000000000000000000000000000000000000000000000000000000008152600160a060020a03808516600483015291519394509116916342087d4f916024808201926020929091908290030181600087803b15801561072f57600080fd5b505af1158015610743573d6000803e3d6000fd5b505050506040513d602081101561075957600080fd5b505115156107d7576040805160e560020a62461bcd02815260206004820152602760248201527f5468657265206973206e6f20636f72726573706f6e64696e672077726974652060448201527f7265717565737400000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600354604080517f31419b50000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152915191909216916331419b5091602480830192600092919082900301818387803b15801561083e57600080fd5b505af1158015610852573d6000803e3d6000fd5b505050505050505600a165627a7a72305820254adc2d1ddd7e4a75469a322041a1b08230ab47b0ef65b51da1fd481119e7f00029`

// DeployCalypso deploys a new Ethereum contract, binding an instance of Calypso to it.
func DeployCalypso(auth *bind.TransactOpts, backend bind.ContractBackend, holder common.Address, rrholder common.Address, o common.Address) (common.Address, *types.Transaction, *Calypso, error) {
	parsed, err := abi.JSON(strings.NewReader(CalypsoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CalypsoBin), backend, holder, rrholder, o)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Calypso{CalypsoCaller: CalypsoCaller{contract: contract}, CalypsoTransactor: CalypsoTransactor{contract: contract}, CalypsoFilterer: CalypsoFilterer{contract: contract}}, nil
}

// Calypso is an auto generated Go binding around an Ethereum contract.
type Calypso struct {
	CalypsoCaller     // Read-only binding to the contract
	CalypsoTransactor // Write-only binding to the contract
	CalypsoFilterer   // Log filterer for contract events
}

// CalypsoCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalypsoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalypsoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalypsoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalypsoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalypsoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalypsoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalypsoSession struct {
	Contract     *Calypso          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalypsoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalypsoCallerSession struct {
	Contract *CalypsoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CalypsoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalypsoTransactorSession struct {
	Contract     *CalypsoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CalypsoRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalypsoRaw struct {
	Contract *Calypso // Generic contract binding to access the raw methods on
}

// CalypsoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalypsoCallerRaw struct {
	Contract *CalypsoCaller // Generic read-only contract binding to access the raw methods on
}

// CalypsoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalypsoTransactorRaw struct {
	Contract *CalypsoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCalypso creates a new instance of Calypso, bound to a specific deployed contract.
func NewCalypso(address common.Address, backend bind.ContractBackend) (*Calypso, error) {
	contract, err := bindCalypso(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Calypso{CalypsoCaller: CalypsoCaller{contract: contract}, CalypsoTransactor: CalypsoTransactor{contract: contract}, CalypsoFilterer: CalypsoFilterer{contract: contract}}, nil
}

// NewCalypsoCaller creates a new read-only instance of Calypso, bound to a specific deployed contract.
func NewCalypsoCaller(address common.Address, caller bind.ContractCaller) (*CalypsoCaller, error) {
	contract, err := bindCalypso(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalypsoCaller{contract: contract}, nil
}

// NewCalypsoTransactor creates a new write-only instance of Calypso, bound to a specific deployed contract.
func NewCalypsoTransactor(address common.Address, transactor bind.ContractTransactor) (*CalypsoTransactor, error) {
	contract, err := bindCalypso(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalypsoTransactor{contract: contract}, nil
}

// NewCalypsoFilterer creates a new log filterer instance of Calypso, bound to a specific deployed contract.
func NewCalypsoFilterer(address common.Address, filterer bind.ContractFilterer) (*CalypsoFilterer, error) {
	contract, err := bindCalypso(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalypsoFilterer{contract: contract}, nil
}

// bindCalypso binds a generic wrapper to an already deployed contract.
func bindCalypso(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CalypsoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calypso *CalypsoRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Calypso.Contract.CalypsoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calypso *CalypsoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calypso.Contract.CalypsoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calypso *CalypsoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calypso.Contract.CalypsoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calypso *CalypsoCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Calypso.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calypso *CalypsoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calypso.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calypso *CalypsoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calypso.Contract.contract.Transact(opts, method, params...)
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Calypso *CalypsoCaller) CanWrite(opts *bind.CallOpts, a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Calypso.contract.Call(opts, out, "canWrite", a)
	return *ret0, err
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Calypso *CalypsoSession) CanWrite(a common.Address) (bool, error) {
	return _Calypso.Contract.CanWrite(&_Calypso.CallOpts, a)
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Calypso *CalypsoCallerSession) CanWrite(a common.Address) (bool, error) {
	return _Calypso.Contract.CanWrite(&_Calypso.CallOpts, a)
}

// IsValidReadRequest is a free data retrieval call binding the contract method 0x778ad09c.
//
// Solidity: function isValidReadRequest(a address) constant returns(bool)
func (_Calypso *CalypsoCaller) IsValidReadRequest(opts *bind.CallOpts, a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Calypso.contract.Call(opts, out, "isValidReadRequest", a)
	return *ret0, err
}

// IsValidReadRequest is a free data retrieval call binding the contract method 0x778ad09c.
//
// Solidity: function isValidReadRequest(a address) constant returns(bool)
func (_Calypso *CalypsoSession) IsValidReadRequest(a common.Address) (bool, error) {
	return _Calypso.Contract.IsValidReadRequest(&_Calypso.CallOpts, a)
}

// IsValidReadRequest is a free data retrieval call binding the contract method 0x778ad09c.
//
// Solidity: function isValidReadRequest(a address) constant returns(bool)
func (_Calypso *CalypsoCallerSession) IsValidReadRequest(a common.Address) (bool, error) {
	return _Calypso.Contract.IsValidReadRequest(&_Calypso.CallOpts, a)
}

// OwnersMap is a free data retrieval call binding the contract method 0x19a444b7.
//
// Solidity: function ownersMap( address) constant returns(bool)
func (_Calypso *CalypsoCaller) OwnersMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Calypso.contract.Call(opts, out, "ownersMap", arg0)
	return *ret0, err
}

// OwnersMap is a free data retrieval call binding the contract method 0x19a444b7.
//
// Solidity: function ownersMap( address) constant returns(bool)
func (_Calypso *CalypsoSession) OwnersMap(arg0 common.Address) (bool, error) {
	return _Calypso.Contract.OwnersMap(&_Calypso.CallOpts, arg0)
}

// OwnersMap is a free data retrieval call binding the contract method 0x19a444b7.
//
// Solidity: function ownersMap( address) constant returns(bool)
func (_Calypso *CalypsoCallerSession) OwnersMap(arg0 common.Address) (bool, error) {
	return _Calypso.Contract.OwnersMap(&_Calypso.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Calypso *CalypsoTransactor) AddOwner(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Calypso.contract.Transact(opts, "AddOwner", a)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Calypso *CalypsoSession) AddOwner(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddOwner(&_Calypso.TransactOpts, a)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Calypso *CalypsoTransactorSession) AddOwner(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddOwner(&_Calypso.TransactOpts, a)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0xbdae9b89.
//
// Solidity: function AddReadRequest(a address) returns()
func (_Calypso *CalypsoTransactor) AddReadRequest(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Calypso.contract.Transact(opts, "AddReadRequest", a)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0xbdae9b89.
//
// Solidity: function AddReadRequest(a address) returns()
func (_Calypso *CalypsoSession) AddReadRequest(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddReadRequest(&_Calypso.TransactOpts, a)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0xbdae9b89.
//
// Solidity: function AddReadRequest(a address) returns()
func (_Calypso *CalypsoTransactorSession) AddReadRequest(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddReadRequest(&_Calypso.TransactOpts, a)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_Calypso *CalypsoTransactor) AddWriteRequest(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Calypso.contract.Transact(opts, "addWriteRequest", a)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_Calypso *CalypsoSession) AddWriteRequest(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddWriteRequest(&_Calypso.TransactOpts, a)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_Calypso *CalypsoTransactorSession) AddWriteRequest(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.AddWriteRequest(&_Calypso.TransactOpts, a)
}

// CheckIfRequestIsValid is a paid mutator transaction binding the contract method 0x4a4e19d3.
//
// Solidity: function checkIfRequestIsValid(a address) returns(bool)
func (_Calypso *CalypsoTransactor) CheckIfRequestIsValid(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Calypso.contract.Transact(opts, "checkIfRequestIsValid", a)
}

// CheckIfRequestIsValid is a paid mutator transaction binding the contract method 0x4a4e19d3.
//
// Solidity: function checkIfRequestIsValid(a address) returns(bool)
func (_Calypso *CalypsoSession) CheckIfRequestIsValid(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.CheckIfRequestIsValid(&_Calypso.TransactOpts, a)
}

// CheckIfRequestIsValid is a paid mutator transaction binding the contract method 0x4a4e19d3.
//
// Solidity: function checkIfRequestIsValid(a address) returns(bool)
func (_Calypso *CalypsoTransactorSession) CheckIfRequestIsValid(a common.Address) (*types.Transaction, error) {
	return _Calypso.Contract.CheckIfRequestIsValid(&_Calypso.TransactOpts, a)
}

// OwnersABI is the input ABI used to generate the binding from.
const OwnersABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"canWrite\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNrOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnersBin is the compiled bytecode used for deploying new contracts.
const OwnersBin = `0x608060405234801561001057600080fd5b50610185806100206000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633239b71b811461005b578063502bb7341461009d578063ac1e9ef4146100c4575b600080fd5b34801561006757600080fd5b5061008973ffffffffffffffffffffffffffffffffffffffff600435166100f4565b604080519115158252519081900360200190f35b3480156100a957600080fd5b506100b261011f565b60408051918252519081900360200190f35b3480156100d057600080fd5b506100f273ffffffffffffffffffffffffffffffffffffffff60043516610125565b005b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205460ff1690565b60005490565b73ffffffffffffffffffffffffffffffffffffffff166000908152600160208190526040909120805460ff191690911790555600a165627a7a723058202ed371272478c472c1407e55d25087a68b749c661521e332f63cda9556e676580029`

// DeployOwners deploys a new Ethereum contract, binding an instance of Owners to it.
func DeployOwners(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owners, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owners{OwnersCaller: OwnersCaller{contract: contract}, OwnersTransactor: OwnersTransactor{contract: contract}, OwnersFilterer: OwnersFilterer{contract: contract}}, nil
}

// Owners is an auto generated Go binding around an Ethereum contract.
type Owners struct {
	OwnersCaller     // Read-only binding to the contract
	OwnersTransactor // Write-only binding to the contract
	OwnersFilterer   // Log filterer for contract events
}

// OwnersCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnersSession struct {
	Contract     *Owners           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnersCallerSession struct {
	Contract *OwnersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnersTransactorSession struct {
	Contract     *OwnersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnersRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnersRaw struct {
	Contract *Owners // Generic contract binding to access the raw methods on
}

// OwnersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnersCallerRaw struct {
	Contract *OwnersCaller // Generic read-only contract binding to access the raw methods on
}

// OwnersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnersTransactorRaw struct {
	Contract *OwnersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwners creates a new instance of Owners, bound to a specific deployed contract.
func NewOwners(address common.Address, backend bind.ContractBackend) (*Owners, error) {
	contract, err := bindOwners(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owners{OwnersCaller: OwnersCaller{contract: contract}, OwnersTransactor: OwnersTransactor{contract: contract}, OwnersFilterer: OwnersFilterer{contract: contract}}, nil
}

// NewOwnersCaller creates a new read-only instance of Owners, bound to a specific deployed contract.
func NewOwnersCaller(address common.Address, caller bind.ContractCaller) (*OwnersCaller, error) {
	contract, err := bindOwners(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnersCaller{contract: contract}, nil
}

// NewOwnersTransactor creates a new write-only instance of Owners, bound to a specific deployed contract.
func NewOwnersTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnersTransactor, error) {
	contract, err := bindOwners(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnersTransactor{contract: contract}, nil
}

// NewOwnersFilterer creates a new log filterer instance of Owners, bound to a specific deployed contract.
func NewOwnersFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnersFilterer, error) {
	contract, err := bindOwners(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnersFilterer{contract: contract}, nil
}

// bindOwners binds a generic wrapper to an already deployed contract.
func bindOwners(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owners *OwnersRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owners.Contract.OwnersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owners *OwnersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owners.Contract.OwnersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owners *OwnersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owners.Contract.OwnersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owners *OwnersCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owners.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owners *OwnersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owners.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owners *OwnersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owners.Contract.contract.Transact(opts, method, params...)
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Owners *OwnersCaller) CanWrite(opts *bind.CallOpts, a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Owners.contract.Call(opts, out, "canWrite", a)
	return *ret0, err
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Owners *OwnersSession) CanWrite(a common.Address) (bool, error) {
	return _Owners.Contract.CanWrite(&_Owners.CallOpts, a)
}

// CanWrite is a free data retrieval call binding the contract method 0x3239b71b.
//
// Solidity: function canWrite(a address) constant returns(bool)
func (_Owners *OwnersCallerSession) CanWrite(a common.Address) (bool, error) {
	return _Owners.Contract.CanWrite(&_Owners.CallOpts, a)
}

// GetNrOwners is a free data retrieval call binding the contract method 0x502bb734.
//
// Solidity: function getNrOwners() constant returns(uint256)
func (_Owners *OwnersCaller) GetNrOwners(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Owners.contract.Call(opts, out, "getNrOwners")
	return *ret0, err
}

// GetNrOwners is a free data retrieval call binding the contract method 0x502bb734.
//
// Solidity: function getNrOwners() constant returns(uint256)
func (_Owners *OwnersSession) GetNrOwners() (*big.Int, error) {
	return _Owners.Contract.GetNrOwners(&_Owners.CallOpts)
}

// GetNrOwners is a free data retrieval call binding the contract method 0x502bb734.
//
// Solidity: function getNrOwners() constant returns(uint256)
func (_Owners *OwnersCallerSession) GetNrOwners() (*big.Int, error) {
	return _Owners.Contract.GetNrOwners(&_Owners.CallOpts)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Owners *OwnersTransactor) AddOwner(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _Owners.contract.Transact(opts, "AddOwner", a)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Owners *OwnersSession) AddOwner(a common.Address) (*types.Transaction, error) {
	return _Owners.Contract.AddOwner(&_Owners.TransactOpts, a)
}

// AddOwner is a paid mutator transaction binding the contract method 0xac1e9ef4.
//
// Solidity: function AddOwner(a address) returns()
func (_Owners *OwnersTransactorSession) AddOwner(a common.Address) (*types.Transaction, error) {
	return _Owners.Contract.AddOwner(&_Owners.TransactOpts, a)
}

// ReadRequestABI is the input ABI used to generate the binding from.
const ReadRequestABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"xc\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writeRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"wr\",\"type\":\"address\"}],\"name\":\"CompareReadWrite\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"x\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReadRequestBin is the compiled bytecode used for deploying new contracts.
const ReadRequestBin = `0x608060405234801561001057600080fd5b506040516103b73803806103b78339810160405280516020808301516002805433600160a060020a03199182161790915560008054909116600160a060020a03851617905590920180519192909161006e9160019190840190610076565b505050610111565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100b757805160ff19168380011785556100e4565b828001600101855582156100e4579182015b828111156100e45782518255916020019190600101906100c9565b506100f09291506100f4565b5090565b61010e91905b808211156100f057600081556001016100fa565b90565b610297806101206000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416635baf4e9f811461006657806364e94e67146100f05780638da5cb5b1461012e5780639d0089b814610143575b600080fd5b34801561007257600080fd5b5061007b610185565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100b557818101518382015260200161009d565b50505050905090810190601f1680156100e25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156100fc57600080fd5b50610105610212565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561013a57600080fd5b5061010561022e565b34801561014f57600080fd5b5061017173ffffffffffffffffffffffffffffffffffffffff6004351661024a565b604080519115158252519081900360200190f35b60018054604080516020600284861615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561020a5780601f106101df5761010080835404028352916020019161020a565b820191906000526020600020905b8154815290600101906020018083116101ed57829003601f168201915b505050505081565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60025473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff908116911614905600a165627a7a72305820b89c9530f16355c8c15d22e1f7e8d6ede7371d517741f773357a812ba6a55b920029`

// DeployReadRequest deploys a new Ethereum contract, binding an instance of ReadRequest to it.
func DeployReadRequest(auth *bind.TransactOpts, backend bind.ContractBackend, a common.Address, x []byte) (common.Address, *types.Transaction, *ReadRequest, error) {
	parsed, err := abi.JSON(strings.NewReader(ReadRequestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReadRequestBin), backend, a, x)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReadRequest{ReadRequestCaller: ReadRequestCaller{contract: contract}, ReadRequestTransactor: ReadRequestTransactor{contract: contract}, ReadRequestFilterer: ReadRequestFilterer{contract: contract}}, nil
}

// ReadRequest is an auto generated Go binding around an Ethereum contract.
type ReadRequest struct {
	ReadRequestCaller     // Read-only binding to the contract
	ReadRequestTransactor // Write-only binding to the contract
	ReadRequestFilterer   // Log filterer for contract events
}

// ReadRequestCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReadRequestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReadRequestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReadRequestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReadRequestSession struct {
	Contract     *ReadRequest      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReadRequestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReadRequestCallerSession struct {
	Contract *ReadRequestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ReadRequestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReadRequestTransactorSession struct {
	Contract     *ReadRequestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ReadRequestRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReadRequestRaw struct {
	Contract *ReadRequest // Generic contract binding to access the raw methods on
}

// ReadRequestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReadRequestCallerRaw struct {
	Contract *ReadRequestCaller // Generic read-only contract binding to access the raw methods on
}

// ReadRequestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReadRequestTransactorRaw struct {
	Contract *ReadRequestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReadRequest creates a new instance of ReadRequest, bound to a specific deployed contract.
func NewReadRequest(address common.Address, backend bind.ContractBackend) (*ReadRequest, error) {
	contract, err := bindReadRequest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReadRequest{ReadRequestCaller: ReadRequestCaller{contract: contract}, ReadRequestTransactor: ReadRequestTransactor{contract: contract}, ReadRequestFilterer: ReadRequestFilterer{contract: contract}}, nil
}

// NewReadRequestCaller creates a new read-only instance of ReadRequest, bound to a specific deployed contract.
func NewReadRequestCaller(address common.Address, caller bind.ContractCaller) (*ReadRequestCaller, error) {
	contract, err := bindReadRequest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReadRequestCaller{contract: contract}, nil
}

// NewReadRequestTransactor creates a new write-only instance of ReadRequest, bound to a specific deployed contract.
func NewReadRequestTransactor(address common.Address, transactor bind.ContractTransactor) (*ReadRequestTransactor, error) {
	contract, err := bindReadRequest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReadRequestTransactor{contract: contract}, nil
}

// NewReadRequestFilterer creates a new log filterer instance of ReadRequest, bound to a specific deployed contract.
func NewReadRequestFilterer(address common.Address, filterer bind.ContractFilterer) (*ReadRequestFilterer, error) {
	contract, err := bindReadRequest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReadRequestFilterer{contract: contract}, nil
}

// bindReadRequest binds a generic wrapper to an already deployed contract.
func bindReadRequest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReadRequestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReadRequest *ReadRequestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReadRequest.Contract.ReadRequestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReadRequest *ReadRequestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReadRequest.Contract.ReadRequestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReadRequest *ReadRequestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReadRequest.Contract.ReadRequestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReadRequest *ReadRequestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReadRequest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReadRequest *ReadRequestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReadRequest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReadRequest *ReadRequestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReadRequest.Contract.contract.Transact(opts, method, params...)
}

// CompareReadWrite is a free data retrieval call binding the contract method 0x9d0089b8.
//
// Solidity: function CompareReadWrite(wr address) constant returns(bool)
func (_ReadRequest *ReadRequestCaller) CompareReadWrite(opts *bind.CallOpts, wr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReadRequest.contract.Call(opts, out, "CompareReadWrite", wr)
	return *ret0, err
}

// CompareReadWrite is a free data retrieval call binding the contract method 0x9d0089b8.
//
// Solidity: function CompareReadWrite(wr address) constant returns(bool)
func (_ReadRequest *ReadRequestSession) CompareReadWrite(wr common.Address) (bool, error) {
	return _ReadRequest.Contract.CompareReadWrite(&_ReadRequest.CallOpts, wr)
}

// CompareReadWrite is a free data retrieval call binding the contract method 0x9d0089b8.
//
// Solidity: function CompareReadWrite(wr address) constant returns(bool)
func (_ReadRequest *ReadRequestCallerSession) CompareReadWrite(wr common.Address) (bool, error) {
	return _ReadRequest.Contract.CompareReadWrite(&_ReadRequest.CallOpts, wr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReadRequest *ReadRequestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReadRequest.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReadRequest *ReadRequestSession) Owner() (common.Address, error) {
	return _ReadRequest.Contract.Owner(&_ReadRequest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReadRequest *ReadRequestCallerSession) Owner() (common.Address, error) {
	return _ReadRequest.Contract.Owner(&_ReadRequest.CallOpts)
}

// WriteRequest is a free data retrieval call binding the contract method 0x64e94e67.
//
// Solidity: function writeRequest() constant returns(address)
func (_ReadRequest *ReadRequestCaller) WriteRequest(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReadRequest.contract.Call(opts, out, "writeRequest")
	return *ret0, err
}

// WriteRequest is a free data retrieval call binding the contract method 0x64e94e67.
//
// Solidity: function writeRequest() constant returns(address)
func (_ReadRequest *ReadRequestSession) WriteRequest() (common.Address, error) {
	return _ReadRequest.Contract.WriteRequest(&_ReadRequest.CallOpts)
}

// WriteRequest is a free data retrieval call binding the contract method 0x64e94e67.
//
// Solidity: function writeRequest() constant returns(address)
func (_ReadRequest *ReadRequestCallerSession) WriteRequest() (common.Address, error) {
	return _ReadRequest.Contract.WriteRequest(&_ReadRequest.CallOpts)
}

// Xc is a free data retrieval call binding the contract method 0x5baf4e9f.
//
// Solidity: function xc() constant returns(bytes)
func (_ReadRequest *ReadRequestCaller) Xc(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _ReadRequest.contract.Call(opts, out, "xc")
	return *ret0, err
}

// Xc is a free data retrieval call binding the contract method 0x5baf4e9f.
//
// Solidity: function xc() constant returns(bytes)
func (_ReadRequest *ReadRequestSession) Xc() ([]byte, error) {
	return _ReadRequest.Contract.Xc(&_ReadRequest.CallOpts)
}

// Xc is a free data retrieval call binding the contract method 0x5baf4e9f.
//
// Solidity: function xc() constant returns(bytes)
func (_ReadRequest *ReadRequestCallerSession) Xc() ([]byte, error) {
	return _ReadRequest.Contract.Xc(&_ReadRequest.CallOpts)
}

// ReadRequestHolderABI is the input ABI used to generate the binding from.
const ReadRequestHolderABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addReadRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"requests\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasReq\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"address\"}],\"name\":\"hasReadRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReadRequestHolderBin is the compiled bytecode used for deploying new contracts.
const ReadRequestHolderBin = `0x608060405234801561001057600080fd5b506101e1806100206000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166331419b50811461006657806374adad1d146100895780638eb81102146100c6578063fd6336dc146100fb575b600080fd5b34801561007257600080fd5b50610087600160a060020a036004351661011c565b005b34801561009557600080fd5b506100aa600160a060020a0360043516610167565b60408051600160a060020a039092168252519081900360200190f35b3480156100d257600080fd5b506100e7600160a060020a0360043516610182565b604080519115158252519081900360200190f35b34801561010757600080fd5b506100e7600160a060020a0360043516610197565b600160a060020a0316600081815260208181526040808320805460ff191660019081179091559091529020805473ffffffffffffffffffffffffffffffffffffffff19169091179055565b600160205260009081526040902054600160a060020a031681565b60006020819052908152604090205460ff1681565b600160a060020a031660009081526020819052604090205460ff16905600a165627a7a72305820d689a2f6c4b8872b6039e598e8f27ddd7d4a1ab31335bf7e83067efc571c96b70029`

// DeployReadRequestHolder deploys a new Ethereum contract, binding an instance of ReadRequestHolder to it.
func DeployReadRequestHolder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReadRequestHolder, error) {
	parsed, err := abi.JSON(strings.NewReader(ReadRequestHolderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReadRequestHolderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReadRequestHolder{ReadRequestHolderCaller: ReadRequestHolderCaller{contract: contract}, ReadRequestHolderTransactor: ReadRequestHolderTransactor{contract: contract}, ReadRequestHolderFilterer: ReadRequestHolderFilterer{contract: contract}}, nil
}

// ReadRequestHolder is an auto generated Go binding around an Ethereum contract.
type ReadRequestHolder struct {
	ReadRequestHolderCaller     // Read-only binding to the contract
	ReadRequestHolderTransactor // Write-only binding to the contract
	ReadRequestHolderFilterer   // Log filterer for contract events
}

// ReadRequestHolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReadRequestHolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestHolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReadRequestHolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestHolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReadRequestHolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReadRequestHolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReadRequestHolderSession struct {
	Contract     *ReadRequestHolder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReadRequestHolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReadRequestHolderCallerSession struct {
	Contract *ReadRequestHolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ReadRequestHolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReadRequestHolderTransactorSession struct {
	Contract     *ReadRequestHolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ReadRequestHolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReadRequestHolderRaw struct {
	Contract *ReadRequestHolder // Generic contract binding to access the raw methods on
}

// ReadRequestHolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReadRequestHolderCallerRaw struct {
	Contract *ReadRequestHolderCaller // Generic read-only contract binding to access the raw methods on
}

// ReadRequestHolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReadRequestHolderTransactorRaw struct {
	Contract *ReadRequestHolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReadRequestHolder creates a new instance of ReadRequestHolder, bound to a specific deployed contract.
func NewReadRequestHolder(address common.Address, backend bind.ContractBackend) (*ReadRequestHolder, error) {
	contract, err := bindReadRequestHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReadRequestHolder{ReadRequestHolderCaller: ReadRequestHolderCaller{contract: contract}, ReadRequestHolderTransactor: ReadRequestHolderTransactor{contract: contract}, ReadRequestHolderFilterer: ReadRequestHolderFilterer{contract: contract}}, nil
}

// NewReadRequestHolderCaller creates a new read-only instance of ReadRequestHolder, bound to a specific deployed contract.
func NewReadRequestHolderCaller(address common.Address, caller bind.ContractCaller) (*ReadRequestHolderCaller, error) {
	contract, err := bindReadRequestHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReadRequestHolderCaller{contract: contract}, nil
}

// NewReadRequestHolderTransactor creates a new write-only instance of ReadRequestHolder, bound to a specific deployed contract.
func NewReadRequestHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*ReadRequestHolderTransactor, error) {
	contract, err := bindReadRequestHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReadRequestHolderTransactor{contract: contract}, nil
}

// NewReadRequestHolderFilterer creates a new log filterer instance of ReadRequestHolder, bound to a specific deployed contract.
func NewReadRequestHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*ReadRequestHolderFilterer, error) {
	contract, err := bindReadRequestHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReadRequestHolderFilterer{contract: contract}, nil
}

// bindReadRequestHolder binds a generic wrapper to an already deployed contract.
func bindReadRequestHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReadRequestHolderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReadRequestHolder *ReadRequestHolderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReadRequestHolder.Contract.ReadRequestHolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReadRequestHolder *ReadRequestHolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.ReadRequestHolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReadRequestHolder *ReadRequestHolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.ReadRequestHolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReadRequestHolder *ReadRequestHolderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReadRequestHolder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReadRequestHolder *ReadRequestHolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReadRequestHolder *ReadRequestHolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.contract.Transact(opts, method, params...)
}

// HasReadRequest is a free data retrieval call binding the contract method 0xfd6336dc.
//
// Solidity: function hasReadRequest(id address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderCaller) HasReadRequest(opts *bind.CallOpts, id common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReadRequestHolder.contract.Call(opts, out, "hasReadRequest", id)
	return *ret0, err
}

// HasReadRequest is a free data retrieval call binding the contract method 0xfd6336dc.
//
// Solidity: function hasReadRequest(id address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderSession) HasReadRequest(id common.Address) (bool, error) {
	return _ReadRequestHolder.Contract.HasReadRequest(&_ReadRequestHolder.CallOpts, id)
}

// HasReadRequest is a free data retrieval call binding the contract method 0xfd6336dc.
//
// Solidity: function hasReadRequest(id address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderCallerSession) HasReadRequest(id common.Address) (bool, error) {
	return _ReadRequestHolder.Contract.HasReadRequest(&_ReadRequestHolder.CallOpts, id)
}

// HasReq is a free data retrieval call binding the contract method 0x8eb81102.
//
// Solidity: function hasReq( address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderCaller) HasReq(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReadRequestHolder.contract.Call(opts, out, "hasReq", arg0)
	return *ret0, err
}

// HasReq is a free data retrieval call binding the contract method 0x8eb81102.
//
// Solidity: function hasReq( address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderSession) HasReq(arg0 common.Address) (bool, error) {
	return _ReadRequestHolder.Contract.HasReq(&_ReadRequestHolder.CallOpts, arg0)
}

// HasReq is a free data retrieval call binding the contract method 0x8eb81102.
//
// Solidity: function hasReq( address) constant returns(bool)
func (_ReadRequestHolder *ReadRequestHolderCallerSession) HasReq(arg0 common.Address) (bool, error) {
	return _ReadRequestHolder.Contract.HasReq(&_ReadRequestHolder.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x74adad1d.
//
// Solidity: function requests( address) constant returns(address)
func (_ReadRequestHolder *ReadRequestHolderCaller) Requests(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReadRequestHolder.contract.Call(opts, out, "requests", arg0)
	return *ret0, err
}

// Requests is a free data retrieval call binding the contract method 0x74adad1d.
//
// Solidity: function requests( address) constant returns(address)
func (_ReadRequestHolder *ReadRequestHolderSession) Requests(arg0 common.Address) (common.Address, error) {
	return _ReadRequestHolder.Contract.Requests(&_ReadRequestHolder.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x74adad1d.
//
// Solidity: function requests( address) constant returns(address)
func (_ReadRequestHolder *ReadRequestHolderCallerSession) Requests(arg0 common.Address) (common.Address, error) {
	return _ReadRequestHolder.Contract.Requests(&_ReadRequestHolder.CallOpts, arg0)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0x31419b50.
//
// Solidity: function addReadRequest(a address) returns()
func (_ReadRequestHolder *ReadRequestHolderTransactor) AddReadRequest(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _ReadRequestHolder.contract.Transact(opts, "addReadRequest", a)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0x31419b50.
//
// Solidity: function addReadRequest(a address) returns()
func (_ReadRequestHolder *ReadRequestHolderSession) AddReadRequest(a common.Address) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.AddReadRequest(&_ReadRequestHolder.TransactOpts, a)
}

// AddReadRequest is a paid mutator transaction binding the contract method 0x31419b50.
//
// Solidity: function addReadRequest(a address) returns()
func (_ReadRequestHolder *ReadRequestHolderTransactorSession) AddReadRequest(a common.Address) (*types.Transaction, error) {
	return _ReadRequestHolder.Contract.AddReadRequest(&_ReadRequestHolder.TransactOpts, a)
}

// WriteRequestABI is the input ABI used to generate the binding from.
const WriteRequestABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"policyMap\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"LTSID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"U\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"F\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"CanRead\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetSlice\",\"outputs\":[{\"name\":\"\",\"type\":\"int64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"extraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"E\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Cs\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"X\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slice\",\"outputs\":[{\"name\":\"\",\"type\":\"int64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getPolicySize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Ubar\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"split\",\"outputs\":[{\"name\":\"\",\"type\":\"int64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"d\",\"type\":\"bytes\"},{\"name\":\"ed\",\"type\":\"bytes\"},{\"name\":\"l\",\"type\":\"bytes\"},{\"name\":\"p\",\"type\":\"address[]\"},{\"name\":\"u\",\"type\":\"bytes\"},{\"name\":\"cs\",\"type\":\"bytes\"},{\"name\":\"sl\",\"type\":\"int64[]\"},{\"name\":\"f\",\"type\":\"bytes\"},{\"name\":\"e\",\"type\":\"bytes\"},{\"name\":\"ub\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// WriteRequestBin is the compiled bytecode used for deploying new contracts.
const WriteRequestBin = `0x60806040523480156200001157600080fd5b5060405162000bf838038062000bf88339810160409081528151602080840151928401516060850151608086015160a087015160c088015160e08901516101008a01516101208b0151988b018051909b9a8b019a978801999688019895880197948501969385019592850194918201939091019160009162000099918391908e0190620001ee565b508951620000af9060019060208d0190620001ee565b508851620000c59060029060208c0190620001ee565b508751620000db9060099060208b019062000273565b508651620000f19060039060208a0190620001ee565b50855162000107906008906020890190620001ee565b5050600b80546001604060020a031916600117905560005b87518163ffffffff16101562000185576001600a60008a8463ffffffff168151811015156200014a57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff19169115159190911790556001016200011f565b84516200019a90600c906020880190620002d9565b508351620001b0906006906020870190620001ee565b508251620001c6906005906020860190620001ee565b508151620001dc906004906020850190620001ee565b50505050505050505050505062000405565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200023157805160ff191683800117855562000261565b8280016001018555821562000261579182015b828111156200026157825182559160200191906001019062000244565b506200026f92915062000397565b5090565b828054828255906000526020600020908101928215620002cb579160200282015b82811115620002cb5782518254600160a060020a031916600160a060020a0390911617825560209092019160019091019062000294565b506200026f929150620003b7565b82805482825590600052602060002090600301600490048101928215620003895791602002820160005b838211156200035257835183826101000a8154816001604060020a03021916908360070b6001604060020a03160217905550926020019260080160208160070104928301926001030262000303565b8015620003875782816101000a8154906001604060020a03021916905560080160208160070104928301926001030262000352565b505b506200026f929150620003de565b620003b491905b808211156200026f57600081556001016200039e565b90565b620003b491905b808211156200026f578054600160a060020a0319168155600101620003be565b620003b491905b808211156200026f5780546001604060020a0319168155600101620003e5565b6107e380620004156000396000f3006080604052600436106100da5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663069cc4e481146100df5780630fdcdbfd146101215780632688454a146101ab57806328811f59146101c057806356554256146101d55780635bf2abcd14610203578063609d33341461026857806373d4a13a1461027d57806392bbf6e814610292578063b6f9cf6a146102a7578063c1599bd9146102bc578063e6a01b5a146102d1578063ecc0f87114610302578063eeeee14d14610329578063f76541761461033e575b600080fd5b3480156100eb57600080fd5b5061010d73ffffffffffffffffffffffffffffffffffffffff60043516610353565b604080519115158252519081900360200190f35b34801561012d57600080fd5b50610136610368565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610170578181015183820152602001610158565b50505050905090810190601f16801561019d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101b757600080fd5b506101366103f3565b3480156101cc57600080fd5b5061013661044e565b3480156101e157600080fd5b5061010d73ffffffffffffffffffffffffffffffffffffffff600435166104a9565b34801561020f57600080fd5b506102186104d4565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561025457818101518382015260200161023c565b505050509050019250505060405180910390f35b34801561027457600080fd5b50610136610552565b34801561028957600080fd5b506101366105ac565b34801561029e57600080fd5b50610136610607565b3480156102b357600080fd5b50610136610662565b3480156102c857600080fd5b506101366106bd565b3480156102dd57600080fd5b506102e9600435610718565b60408051600792830b90920b8252519081900360200190f35b34801561030e57600080fd5b5061031761074d565b60408051918252519081900360200190f35b34801561033557600080fd5b50610136610753565b34801561034a57600080fd5b506102e96107ae565b600a6020526000908152604090205460ff1681565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156103eb5780601f106103c0576101008083540402835291602001916103eb565b820191906000526020600020905b8154815290600101906020018083116103ce57829003601f168201915b505050505081565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103eb5780601f106103c0576101008083540402835291602001916103eb565b6006805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103eb5780601f106103c0576101008083540402835291602001916103eb565b73ffffffffffffffffffffffffffffffffffffffff166000908152600a602052604090205460ff1690565b6060600c80548060200260200160405190810160405280929190818152602001828054801561054857602002820191906000526020600020906000905b82829054906101000a900460070b60070b815260200190600801906020826007010492830192600103820291508084116105115790505b5050505050905090565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103eb5780601f106103c0576101008083540402835291602001916103eb565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103eb5780601f106103c0576101008083540402835291602001916103eb565b6005805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103eb5780601f106103c0576101008083540402835291602001916103eb565b6008805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103eb5780601f106103c0576101008083540402835291602001916103eb565b6007805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103eb5780601f106103c0576101008083540402835291602001916103eb565b600c80548290811061072657fe5b9060005260206000209060049182820401919006600802915054906101000a900460070b81565b60095490565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103eb5780601f106103c0576101008083540402835291602001916103eb565b600b5460070b815600a165627a7a723058202fc3664cb3a25071071e9dc7f2e01147cd65a6638d7855625e89c103113dbfc80029`

// DeployWriteRequest deploys a new Ethereum contract, binding an instance of WriteRequest to it.
func DeployWriteRequest(auth *bind.TransactOpts, backend bind.ContractBackend, d []byte, ed []byte, l []byte, p []common.Address, u []byte, cs []byte, sl []int64, f []byte, e []byte, ub []byte) (common.Address, *types.Transaction, *WriteRequest, error) {
	parsed, err := abi.JSON(strings.NewReader(WriteRequestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WriteRequestBin), backend, d, ed, l, p, u, cs, sl, f, e, ub)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WriteRequest{WriteRequestCaller: WriteRequestCaller{contract: contract}, WriteRequestTransactor: WriteRequestTransactor{contract: contract}, WriteRequestFilterer: WriteRequestFilterer{contract: contract}}, nil
}

// WriteRequest is an auto generated Go binding around an Ethereum contract.
type WriteRequest struct {
	WriteRequestCaller     // Read-only binding to the contract
	WriteRequestTransactor // Write-only binding to the contract
	WriteRequestFilterer   // Log filterer for contract events
}

// WriteRequestCaller is an auto generated read-only Go binding around an Ethereum contract.
type WriteRequestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WriteRequestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WriteRequestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WriteRequestSession struct {
	Contract     *WriteRequest     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WriteRequestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WriteRequestCallerSession struct {
	Contract *WriteRequestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WriteRequestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WriteRequestTransactorSession struct {
	Contract     *WriteRequestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WriteRequestRaw is an auto generated low-level Go binding around an Ethereum contract.
type WriteRequestRaw struct {
	Contract *WriteRequest // Generic contract binding to access the raw methods on
}

// WriteRequestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WriteRequestCallerRaw struct {
	Contract *WriteRequestCaller // Generic read-only contract binding to access the raw methods on
}

// WriteRequestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WriteRequestTransactorRaw struct {
	Contract *WriteRequestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWriteRequest creates a new instance of WriteRequest, bound to a specific deployed contract.
func NewWriteRequest(address common.Address, backend bind.ContractBackend) (*WriteRequest, error) {
	contract, err := bindWriteRequest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WriteRequest{WriteRequestCaller: WriteRequestCaller{contract: contract}, WriteRequestTransactor: WriteRequestTransactor{contract: contract}, WriteRequestFilterer: WriteRequestFilterer{contract: contract}}, nil
}

// NewWriteRequestCaller creates a new read-only instance of WriteRequest, bound to a specific deployed contract.
func NewWriteRequestCaller(address common.Address, caller bind.ContractCaller) (*WriteRequestCaller, error) {
	contract, err := bindWriteRequest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WriteRequestCaller{contract: contract}, nil
}

// NewWriteRequestTransactor creates a new write-only instance of WriteRequest, bound to a specific deployed contract.
func NewWriteRequestTransactor(address common.Address, transactor bind.ContractTransactor) (*WriteRequestTransactor, error) {
	contract, err := bindWriteRequest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WriteRequestTransactor{contract: contract}, nil
}

// NewWriteRequestFilterer creates a new log filterer instance of WriteRequest, bound to a specific deployed contract.
func NewWriteRequestFilterer(address common.Address, filterer bind.ContractFilterer) (*WriteRequestFilterer, error) {
	contract, err := bindWriteRequest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WriteRequestFilterer{contract: contract}, nil
}

// bindWriteRequest binds a generic wrapper to an already deployed contract.
func bindWriteRequest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WriteRequestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WriteRequest *WriteRequestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WriteRequest.Contract.WriteRequestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WriteRequest *WriteRequestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WriteRequest.Contract.WriteRequestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WriteRequest *WriteRequestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WriteRequest.Contract.WriteRequestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WriteRequest *WriteRequestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WriteRequest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WriteRequest *WriteRequestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WriteRequest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WriteRequest *WriteRequestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WriteRequest.Contract.contract.Transact(opts, method, params...)
}

// CanRead is a free data retrieval call binding the contract method 0x56554256.
//
// Solidity: function CanRead(a address) constant returns(bool)
func (_WriteRequest *WriteRequestCaller) CanRead(opts *bind.CallOpts, a common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "CanRead", a)
	return *ret0, err
}

// CanRead is a free data retrieval call binding the contract method 0x56554256.
//
// Solidity: function CanRead(a address) constant returns(bool)
func (_WriteRequest *WriteRequestSession) CanRead(a common.Address) (bool, error) {
	return _WriteRequest.Contract.CanRead(&_WriteRequest.CallOpts, a)
}

// CanRead is a free data retrieval call binding the contract method 0x56554256.
//
// Solidity: function CanRead(a address) constant returns(bool)
func (_WriteRequest *WriteRequestCallerSession) CanRead(a common.Address) (bool, error) {
	return _WriteRequest.Contract.CanRead(&_WriteRequest.CallOpts, a)
}

// Cs is a free data retrieval call binding the contract method 0xb6f9cf6a.
//
// Solidity: function Cs() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) Cs(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "Cs")
	return *ret0, err
}

// Cs is a free data retrieval call binding the contract method 0xb6f9cf6a.
//
// Solidity: function Cs() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) Cs() ([]byte, error) {
	return _WriteRequest.Contract.Cs(&_WriteRequest.CallOpts)
}

// Cs is a free data retrieval call binding the contract method 0xb6f9cf6a.
//
// Solidity: function Cs() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) Cs() ([]byte, error) {
	return _WriteRequest.Contract.Cs(&_WriteRequest.CallOpts)
}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) E(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "E")
	return *ret0, err
}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) E() ([]byte, error) {
	return _WriteRequest.Contract.E(&_WriteRequest.CallOpts)
}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) E() ([]byte, error) {
	return _WriteRequest.Contract.E(&_WriteRequest.CallOpts)
}

// F is a free data retrieval call binding the contract method 0x28811f59.
//
// Solidity: function F() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) F(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "F")
	return *ret0, err
}

// F is a free data retrieval call binding the contract method 0x28811f59.
//
// Solidity: function F() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) F() ([]byte, error) {
	return _WriteRequest.Contract.F(&_WriteRequest.CallOpts)
}

// F is a free data retrieval call binding the contract method 0x28811f59.
//
// Solidity: function F() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) F() ([]byte, error) {
	return _WriteRequest.Contract.F(&_WriteRequest.CallOpts)
}

// GetSlice is a free data retrieval call binding the contract method 0x5bf2abcd.
//
// Solidity: function GetSlice() constant returns(int64[])
func (_WriteRequest *WriteRequestCaller) GetSlice(opts *bind.CallOpts) ([]int64, error) {
	var (
		ret0 = new([]int64)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "GetSlice")
	return *ret0, err
}

// GetSlice is a free data retrieval call binding the contract method 0x5bf2abcd.
//
// Solidity: function GetSlice() constant returns(int64[])
func (_WriteRequest *WriteRequestSession) GetSlice() ([]int64, error) {
	return _WriteRequest.Contract.GetSlice(&_WriteRequest.CallOpts)
}

// GetSlice is a free data retrieval call binding the contract method 0x5bf2abcd.
//
// Solidity: function GetSlice() constant returns(int64[])
func (_WriteRequest *WriteRequestCallerSession) GetSlice() ([]int64, error) {
	return _WriteRequest.Contract.GetSlice(&_WriteRequest.CallOpts)
}

// LTSID is a free data retrieval call binding the contract method 0x0fdcdbfd.
//
// Solidity: function LTSID() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) LTSID(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "LTSID")
	return *ret0, err
}

// LTSID is a free data retrieval call binding the contract method 0x0fdcdbfd.
//
// Solidity: function LTSID() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) LTSID() ([]byte, error) {
	return _WriteRequest.Contract.LTSID(&_WriteRequest.CallOpts)
}

// LTSID is a free data retrieval call binding the contract method 0x0fdcdbfd.
//
// Solidity: function LTSID() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) LTSID() ([]byte, error) {
	return _WriteRequest.Contract.LTSID(&_WriteRequest.CallOpts)
}

// U is a free data retrieval call binding the contract method 0x2688454a.
//
// Solidity: function U() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) U(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "U")
	return *ret0, err
}

// U is a free data retrieval call binding the contract method 0x2688454a.
//
// Solidity: function U() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) U() ([]byte, error) {
	return _WriteRequest.Contract.U(&_WriteRequest.CallOpts)
}

// U is a free data retrieval call binding the contract method 0x2688454a.
//
// Solidity: function U() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) U() ([]byte, error) {
	return _WriteRequest.Contract.U(&_WriteRequest.CallOpts)
}

// Ubar is a free data retrieval call binding the contract method 0xeeeee14d.
//
// Solidity: function Ubar() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) Ubar(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "Ubar")
	return *ret0, err
}

// Ubar is a free data retrieval call binding the contract method 0xeeeee14d.
//
// Solidity: function Ubar() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) Ubar() ([]byte, error) {
	return _WriteRequest.Contract.Ubar(&_WriteRequest.CallOpts)
}

// Ubar is a free data retrieval call binding the contract method 0xeeeee14d.
//
// Solidity: function Ubar() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) Ubar() ([]byte, error) {
	return _WriteRequest.Contract.Ubar(&_WriteRequest.CallOpts)
}

// X is a free data retrieval call binding the contract method 0xc1599bd9.
//
// Solidity: function X() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) X(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "X")
	return *ret0, err
}

// X is a free data retrieval call binding the contract method 0xc1599bd9.
//
// Solidity: function X() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) X() ([]byte, error) {
	return _WriteRequest.Contract.X(&_WriteRequest.CallOpts)
}

// X is a free data retrieval call binding the contract method 0xc1599bd9.
//
// Solidity: function X() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) X() ([]byte, error) {
	return _WriteRequest.Contract.X(&_WriteRequest.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) Data(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) Data() ([]byte, error) {
	return _WriteRequest.Contract.Data(&_WriteRequest.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) Data() ([]byte, error) {
	return _WriteRequest.Contract.Data(&_WriteRequest.CallOpts)
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() constant returns(bytes)
func (_WriteRequest *WriteRequestCaller) ExtraData(opts *bind.CallOpts) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "extraData")
	return *ret0, err
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() constant returns(bytes)
func (_WriteRequest *WriteRequestSession) ExtraData() ([]byte, error) {
	return _WriteRequest.Contract.ExtraData(&_WriteRequest.CallOpts)
}

// ExtraData is a free data retrieval call binding the contract method 0x609d3334.
//
// Solidity: function extraData() constant returns(bytes)
func (_WriteRequest *WriteRequestCallerSession) ExtraData() ([]byte, error) {
	return _WriteRequest.Contract.ExtraData(&_WriteRequest.CallOpts)
}

// PolicyMap is a free data retrieval call binding the contract method 0x069cc4e4.
//
// Solidity: function policyMap( address) constant returns(bool)
func (_WriteRequest *WriteRequestCaller) PolicyMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "policyMap", arg0)
	return *ret0, err
}

// PolicyMap is a free data retrieval call binding the contract method 0x069cc4e4.
//
// Solidity: function policyMap( address) constant returns(bool)
func (_WriteRequest *WriteRequestSession) PolicyMap(arg0 common.Address) (bool, error) {
	return _WriteRequest.Contract.PolicyMap(&_WriteRequest.CallOpts, arg0)
}

// PolicyMap is a free data retrieval call binding the contract method 0x069cc4e4.
//
// Solidity: function policyMap( address) constant returns(bool)
func (_WriteRequest *WriteRequestCallerSession) PolicyMap(arg0 common.Address) (bool, error) {
	return _WriteRequest.Contract.PolicyMap(&_WriteRequest.CallOpts, arg0)
}

// Slice is a free data retrieval call binding the contract method 0xe6a01b5a.
//
// Solidity: function slice( uint256) constant returns(int64)
func (_WriteRequest *WriteRequestCaller) Slice(opts *bind.CallOpts, arg0 *big.Int) (int64, error) {
	var (
		ret0 = new(int64)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "slice", arg0)
	return *ret0, err
}

// Slice is a free data retrieval call binding the contract method 0xe6a01b5a.
//
// Solidity: function slice( uint256) constant returns(int64)
func (_WriteRequest *WriteRequestSession) Slice(arg0 *big.Int) (int64, error) {
	return _WriteRequest.Contract.Slice(&_WriteRequest.CallOpts, arg0)
}

// Slice is a free data retrieval call binding the contract method 0xe6a01b5a.
//
// Solidity: function slice( uint256) constant returns(int64)
func (_WriteRequest *WriteRequestCallerSession) Slice(arg0 *big.Int) (int64, error) {
	return _WriteRequest.Contract.Slice(&_WriteRequest.CallOpts, arg0)
}

// Split is a free data retrieval call binding the contract method 0xf7654176.
//
// Solidity: function split() constant returns(int64)
func (_WriteRequest *WriteRequestCaller) Split(opts *bind.CallOpts) (int64, error) {
	var (
		ret0 = new(int64)
	)
	out := ret0
	err := _WriteRequest.contract.Call(opts, out, "split")
	return *ret0, err
}

// Split is a free data retrieval call binding the contract method 0xf7654176.
//
// Solidity: function split() constant returns(int64)
func (_WriteRequest *WriteRequestSession) Split() (int64, error) {
	return _WriteRequest.Contract.Split(&_WriteRequest.CallOpts)
}

// Split is a free data retrieval call binding the contract method 0xf7654176.
//
// Solidity: function split() constant returns(int64)
func (_WriteRequest *WriteRequestCallerSession) Split() (int64, error) {
	return _WriteRequest.Contract.Split(&_WriteRequest.CallOpts)
}

// GetPolicySize is a paid mutator transaction binding the contract method 0xecc0f871.
//
// Solidity: function getPolicySize() returns(uint256)
func (_WriteRequest *WriteRequestTransactor) GetPolicySize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WriteRequest.contract.Transact(opts, "getPolicySize")
}

// GetPolicySize is a paid mutator transaction binding the contract method 0xecc0f871.
//
// Solidity: function getPolicySize() returns(uint256)
func (_WriteRequest *WriteRequestSession) GetPolicySize() (*types.Transaction, error) {
	return _WriteRequest.Contract.GetPolicySize(&_WriteRequest.TransactOpts)
}

// GetPolicySize is a paid mutator transaction binding the contract method 0xecc0f871.
//
// Solidity: function getPolicySize() returns(uint256)
func (_WriteRequest *WriteRequestTransactorSession) GetPolicySize() (*types.Transaction, error) {
	return _WriteRequest.Contract.GetPolicySize(&_WriteRequest.TransactOpts)
}

// WriteRequestHolderABI is the input ABI used to generate the binding from.
const WriteRequestHolderABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addWriteRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"canRead\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"getRequestByID\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"wrmap\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// WriteRequestHolderBin is the compiled bytecode used for deploying new contracts.
const WriteRequestHolderBin = `0x608060405234801561001057600080fd5b5061022d806100206000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631029a017811461007157806342087d4f14610094578063b7312591146100c9578063e216ea2b14610106578063ed85db2314610127575b600080fd5b34801561007d57600080fd5b50610092600160a060020a0360043516610148565b005b3480156100a057600080fd5b506100b5600160a060020a0360043516610195565b604080519115158252519081900360200190f35b3480156100d557600080fd5b506100ea600160a060020a03600435166101b3565b60408051600160a060020a039092168252519081900360200190f35b34801561011257600080fd5b506100b5600160a060020a03600435166101d1565b34801561013357600080fd5b506100ea600160a060020a03600435166101e6565b600160a060020a0316600081815260208181526040808320805473ffffffffffffffffffffffffffffffffffffffff1916909417909355600190819052919020805460ff19169091179055565b600160a060020a031660009081526001602052604090205460ff1690565b600160a060020a039081166000908152602081905260409020541690565b60016020526000908152604090205460ff1681565b600060208190529081526040902054600160a060020a0316815600a165627a7a723058208a8c45930c2bb970c5519698029898f5aa2cec025772013786a8f24cc819e81e0029`

// DeployWriteRequestHolder deploys a new Ethereum contract, binding an instance of WriteRequestHolder to it.
func DeployWriteRequestHolder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WriteRequestHolder, error) {
	parsed, err := abi.JSON(strings.NewReader(WriteRequestHolderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WriteRequestHolderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WriteRequestHolder{WriteRequestHolderCaller: WriteRequestHolderCaller{contract: contract}, WriteRequestHolderTransactor: WriteRequestHolderTransactor{contract: contract}, WriteRequestHolderFilterer: WriteRequestHolderFilterer{contract: contract}}, nil
}

// WriteRequestHolder is an auto generated Go binding around an Ethereum contract.
type WriteRequestHolder struct {
	WriteRequestHolderCaller     // Read-only binding to the contract
	WriteRequestHolderTransactor // Write-only binding to the contract
	WriteRequestHolderFilterer   // Log filterer for contract events
}

// WriteRequestHolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type WriteRequestHolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestHolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WriteRequestHolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestHolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WriteRequestHolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WriteRequestHolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WriteRequestHolderSession struct {
	Contract     *WriteRequestHolder // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WriteRequestHolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WriteRequestHolderCallerSession struct {
	Contract *WriteRequestHolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// WriteRequestHolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WriteRequestHolderTransactorSession struct {
	Contract     *WriteRequestHolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// WriteRequestHolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type WriteRequestHolderRaw struct {
	Contract *WriteRequestHolder // Generic contract binding to access the raw methods on
}

// WriteRequestHolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WriteRequestHolderCallerRaw struct {
	Contract *WriteRequestHolderCaller // Generic read-only contract binding to access the raw methods on
}

// WriteRequestHolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WriteRequestHolderTransactorRaw struct {
	Contract *WriteRequestHolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWriteRequestHolder creates a new instance of WriteRequestHolder, bound to a specific deployed contract.
func NewWriteRequestHolder(address common.Address, backend bind.ContractBackend) (*WriteRequestHolder, error) {
	contract, err := bindWriteRequestHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WriteRequestHolder{WriteRequestHolderCaller: WriteRequestHolderCaller{contract: contract}, WriteRequestHolderTransactor: WriteRequestHolderTransactor{contract: contract}, WriteRequestHolderFilterer: WriteRequestHolderFilterer{contract: contract}}, nil
}

// NewWriteRequestHolderCaller creates a new read-only instance of WriteRequestHolder, bound to a specific deployed contract.
func NewWriteRequestHolderCaller(address common.Address, caller bind.ContractCaller) (*WriteRequestHolderCaller, error) {
	contract, err := bindWriteRequestHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WriteRequestHolderCaller{contract: contract}, nil
}

// NewWriteRequestHolderTransactor creates a new write-only instance of WriteRequestHolder, bound to a specific deployed contract.
func NewWriteRequestHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*WriteRequestHolderTransactor, error) {
	contract, err := bindWriteRequestHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WriteRequestHolderTransactor{contract: contract}, nil
}

// NewWriteRequestHolderFilterer creates a new log filterer instance of WriteRequestHolder, bound to a specific deployed contract.
func NewWriteRequestHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*WriteRequestHolderFilterer, error) {
	contract, err := bindWriteRequestHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WriteRequestHolderFilterer{contract: contract}, nil
}

// bindWriteRequestHolder binds a generic wrapper to an already deployed contract.
func bindWriteRequestHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WriteRequestHolderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WriteRequestHolder *WriteRequestHolderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WriteRequestHolder.Contract.WriteRequestHolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WriteRequestHolder *WriteRequestHolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.WriteRequestHolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WriteRequestHolder *WriteRequestHolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.WriteRequestHolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WriteRequestHolder *WriteRequestHolderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WriteRequestHolder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WriteRequestHolder *WriteRequestHolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WriteRequestHolder *WriteRequestHolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.contract.Transact(opts, method, params...)
}

// GetRequestByID is a free data retrieval call binding the contract method 0xb7312591.
//
// Solidity: function getRequestByID(a address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderCaller) GetRequestByID(opts *bind.CallOpts, a common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WriteRequestHolder.contract.Call(opts, out, "getRequestByID", a)
	return *ret0, err
}

// GetRequestByID is a free data retrieval call binding the contract method 0xb7312591.
//
// Solidity: function getRequestByID(a address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderSession) GetRequestByID(a common.Address) (common.Address, error) {
	return _WriteRequestHolder.Contract.GetRequestByID(&_WriteRequestHolder.CallOpts, a)
}

// GetRequestByID is a free data retrieval call binding the contract method 0xb7312591.
//
// Solidity: function getRequestByID(a address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderCallerSession) GetRequestByID(a common.Address) (common.Address, error) {
	return _WriteRequestHolder.Contract.GetRequestByID(&_WriteRequestHolder.CallOpts, a)
}

// HasRequest is a free data retrieval call binding the contract method 0xe216ea2b.
//
// Solidity: function hasRequest( address) constant returns(bool)
func (_WriteRequestHolder *WriteRequestHolderCaller) HasRequest(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _WriteRequestHolder.contract.Call(opts, out, "hasRequest", arg0)
	return *ret0, err
}

// HasRequest is a free data retrieval call binding the contract method 0xe216ea2b.
//
// Solidity: function hasRequest( address) constant returns(bool)
func (_WriteRequestHolder *WriteRequestHolderSession) HasRequest(arg0 common.Address) (bool, error) {
	return _WriteRequestHolder.Contract.HasRequest(&_WriteRequestHolder.CallOpts, arg0)
}

// HasRequest is a free data retrieval call binding the contract method 0xe216ea2b.
//
// Solidity: function hasRequest( address) constant returns(bool)
func (_WriteRequestHolder *WriteRequestHolderCallerSession) HasRequest(arg0 common.Address) (bool, error) {
	return _WriteRequestHolder.Contract.HasRequest(&_WriteRequestHolder.CallOpts, arg0)
}

// Wrmap is a free data retrieval call binding the contract method 0xed85db23.
//
// Solidity: function wrmap( address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderCaller) Wrmap(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _WriteRequestHolder.contract.Call(opts, out, "wrmap", arg0)
	return *ret0, err
}

// Wrmap is a free data retrieval call binding the contract method 0xed85db23.
//
// Solidity: function wrmap( address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderSession) Wrmap(arg0 common.Address) (common.Address, error) {
	return _WriteRequestHolder.Contract.Wrmap(&_WriteRequestHolder.CallOpts, arg0)
}

// Wrmap is a free data retrieval call binding the contract method 0xed85db23.
//
// Solidity: function wrmap( address) constant returns(address)
func (_WriteRequestHolder *WriteRequestHolderCallerSession) Wrmap(arg0 common.Address) (common.Address, error) {
	return _WriteRequestHolder.Contract.Wrmap(&_WriteRequestHolder.CallOpts, arg0)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_WriteRequestHolder *WriteRequestHolderTransactor) AddWriteRequest(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.contract.Transact(opts, "addWriteRequest", a)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_WriteRequestHolder *WriteRequestHolderSession) AddWriteRequest(a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.AddWriteRequest(&_WriteRequestHolder.TransactOpts, a)
}

// AddWriteRequest is a paid mutator transaction binding the contract method 0x1029a017.
//
// Solidity: function addWriteRequest(a address) returns()
func (_WriteRequestHolder *WriteRequestHolderTransactorSession) AddWriteRequest(a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.AddWriteRequest(&_WriteRequestHolder.TransactOpts, a)
}

// CanRead is a paid mutator transaction binding the contract method 0x42087d4f.
//
// Solidity: function canRead(a address) returns(bool)
func (_WriteRequestHolder *WriteRequestHolderTransactor) CanRead(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.contract.Transact(opts, "canRead", a)
}

// CanRead is a paid mutator transaction binding the contract method 0x42087d4f.
//
// Solidity: function canRead(a address) returns(bool)
func (_WriteRequestHolder *WriteRequestHolderSession) CanRead(a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.CanRead(&_WriteRequestHolder.TransactOpts, a)
}

// CanRead is a paid mutator transaction binding the contract method 0x42087d4f.
//
// Solidity: function canRead(a address) returns(bool)
func (_WriteRequestHolder *WriteRequestHolderTransactorSession) CanRead(a common.Address) (*types.Transaction, error) {
	return _WriteRequestHolder.Contract.CanRead(&_WriteRequestHolder.TransactOpts, a)
}

