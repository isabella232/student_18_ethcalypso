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
const CalypsoBin = `0x608060405234801561001057600080fd5b506040516060806106b983398101604090815281516020830151919092015160008054600160a060020a0319908116600160a060020a039384161782556001805482169584169590951790945560038054909416919092161790915561063d90819061007c90396000f3006080604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631029a017811461008757806319a444b7146100aa5780633239b71b146100df5780634a4e19d314610100578063778ad09c14610121578063ac1e9ef414610142578063bdae9b8914610163575b600080fd5b34801561009357600080fd5b506100a8600160a060020a0360043516610184565b005b3480156100b657600080fd5b506100cb600160a060020a036004351661027c565b604080519115158252519081900360200190f35b3480156100eb57600080fd5b506100cb600160a060020a0360043516610291565b34801561010c57600080fd5b506100cb600160a060020a03600435166102af565b34801561012d57600080fd5b506100cb600160a060020a036004351661034c565b34801561014e57600080fd5b506100a8600160a060020a03600435166103b7565b34801561016f57600080fd5b506100a8600160a060020a03600435166103db565b61018d33610291565b15156101fa57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f596f7520617265206e6f7420612072656769737465726564206f776e65720000604482015290519081900360640190fd5b600154604080517f1029a017000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291519190921691631029a01791602480830192600092919082900301818387803b15801561026157600080fd5b505af1158015610275573d6000803e3d6000fd5b5050505050565b60046020526000908152604090205460ff1681565b600160a060020a031660009081526004602052604090205460ff1690565b600154604080517f42087d4f000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916342087d4f9160248082019260209290919082900301818787803b15801561031a57600080fd5b505af115801561032e573d6000803e3d6000fd5b505050506040513d602081101561034457600080fd5b505192915050565b600354604080517ffd6336dc000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151600093929092169163fd6336dc9160248082019260209290919082900301818787803b15801561031a57600080fd5b600160a060020a03166000908152600460205260409020805460ff19166001179055565b60008082915081600160a060020a03166364e94e676040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561043857600080fd5b505af115801561044c573d6000803e3d6000fd5b505050506040513d602081101561046257600080fd5b5051600154604080517f42087d4f000000000000000000000000000000000000000000000000000000008152600160a060020a03808516600483015291519394509116916342087d4f916024808201926020929091908290030181600087803b1580156104ce57600080fd5b505af11580156104e2573d6000803e3d6000fd5b505050506040513d60208110156104f857600080fd5b5051151561058d57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f5468657265206973206e6f20636f72726573706f6e64696e672077726974652060448201527f7265717565737400000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600354604080517f31419b50000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152915191909216916331419b5091602480830192600092919082900301818387803b1580156105f457600080fd5b505af1158015610608573d6000803e3d6000fd5b505050505050505600a165627a7a7230582074012473280cad8e448e19bb839b3c85a1123a2f13da9a0b09a6852d21e2d5070029`

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
const ReadRequestABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"xc\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"writeRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"wr\",\"type\":\"address\"}],\"name\":\"CompareReadWrite\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"x\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReadRequestBin is the compiled bytecode used for deploying new contracts.
const ReadRequestBin = `0x608060405234801561001057600080fd5b5060405161036d38038061036d83398101604052805160208083015160008054600160a060020a031916600160a060020a0385161790559092018051919290916100609160019190840190610068565b505050610103565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a957805160ff19168380011785556100d6565b828001600101855582156100d6579182015b828111156100d65782518255916020019190600101906100bb565b506100e29291506100e6565b5090565b61010091905b808211156100e257600081556001016100ec565b90565b61025b806101126000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416635baf4e9f811461005b57806364e94e67146100e55780639d0089b814610123575b600080fd5b34801561006757600080fd5b50610070610165565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100aa578181015183820152602001610092565b50505050905090810190601f1680156100d75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156100f157600080fd5b506100fa6101f2565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561012f57600080fd5b5061015173ffffffffffffffffffffffffffffffffffffffff6004351661020e565b604080519115158252519081900360200190f35b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101ea5780601f106101bf576101008083540402835291602001916101ea565b820191906000526020600020905b8154815290600101906020018083116101cd57829003601f168201915b505050505081565b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff908116911614905600a165627a7a72305820d145554985deda3b942cf92099d6ab9fde0bdda7fd43c02b7528e7fe743ea6590029`

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
const ReadRequestHolderBin = `0x608060405234801561001057600080fd5b506101e1806100206000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166331419b50811461006657806374adad1d146100895780638eb81102146100c6578063fd6336dc146100fb575b600080fd5b34801561007257600080fd5b50610087600160a060020a036004351661011c565b005b34801561009557600080fd5b506100aa600160a060020a0360043516610167565b60408051600160a060020a039092168252519081900360200190f35b3480156100d257600080fd5b506100e7600160a060020a0360043516610182565b604080519115158252519081900360200190f35b34801561010757600080fd5b506100e7600160a060020a0360043516610197565b600160a060020a0316600081815260208181526040808320805460ff191660019081179091559091529020805473ffffffffffffffffffffffffffffffffffffffff19169091179055565b600160205260009081526040902054600160a060020a031681565b60006020819052908152604090205460ff1681565b600160a060020a031660009081526020819052604090205460ff16905600a165627a7a723058201939f5aec546a585edfd0a935899e6fe2027a5a0db8219cd207936bfed3ce00e0029`

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
const WriteRequestABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"LTSID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"U\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"extraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Cs\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getPolicySize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"split\",\"outputs\":[{\"name\":\"\",\"type\":\"int64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"d\",\"type\":\"bytes\"},{\"name\":\"ed\",\"type\":\"bytes\"},{\"name\":\"l\",\"type\":\"bytes\"},{\"name\":\"p\",\"type\":\"address[]\"},{\"name\":\"u\",\"type\":\"bytes\"},{\"name\":\"cs\",\"type\":\"bytes\"},{\"name\":\"s\",\"type\":\"int64\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// WriteRequestBin is the compiled bytecode used for deploying new contracts.
const WriteRequestBin = `0x60806040523480156200001157600080fd5b506040516200097838038062000978833981018060405262000037919081019062000322565b86516200004c9060009060208a0190620000ee565b50855162000062906001906020890190620000ee565b50845162000078906002906020880190620000ee565b5083516200008e90600590602087019062000173565b508251620000a4906003906020860190620000ee565b508151620000ba906004906020850190620000ee565b506006805460079290920b6001604060020a031667ffffffffffffffff199092169190911790555062000511945050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200013157805160ff191683800117855562000161565b8280016001018555821562000161579182015b828111156200016157825182559160200191906001019062000144565b506200016f929150620001d9565b5090565b828054828255906000526020600020908101928215620001cb579160200282015b82811115620001cb5782518254600160a060020a031916600160a060020a0390911617825560209092019160019091019062000194565b506200016f929150620001f9565b620001f691905b808211156200016f5760008155600101620001e0565b90565b620001f691905b808211156200016f578054600160a060020a031916815560010162000200565b60006200022e8251620004cc565b9392505050565b6000601f820183136200024757600080fd5b81516200025e620002588262000483565b6200045c565b915081818352602084019350602081019050838560208402820111156200028457600080fd5b60005b83811015620002b457816200029d888262000220565b845250602092830192919091019060010162000287565b5050505092915050565b6000601f82018313620002d057600080fd5b8151620002e16200025882620004a4565b91508082526020830160208301858383011115620002fe57600080fd5b6200030b838284620004de565b50505092915050565b60006200022e8251620004d8565b600080600080600080600060e0888a0312156200033e57600080fd5b87516001604060020a038111156200035557600080fd5b620003638a828b01620002be565b97505060208801516001604060020a038111156200038057600080fd5b6200038e8a828b01620002be565b96505060408801516001604060020a03811115620003ab57600080fd5b620003b98a828b01620002be565b95505060608801516001604060020a03811115620003d657600080fd5b620003e48a828b0162000235565b94505060808801516001604060020a038111156200040157600080fd5b6200040f8a828b01620002be565b93505060a08801516001604060020a038111156200042c57600080fd5b6200043a8a828b01620002be565b92505060c06200044d8a828b0162000314565b91505092959891949750929550565b6040518181016001604060020a03811182821017156200047b57600080fd5b604052919050565b60006001604060020a038211156200049a57600080fd5b5060209081020190565b60006001604060020a03821115620004bb57600080fd5b506020601f91909101601f19160190565b600160a060020a031690565b60070b90565b60005b83811015620004fb578181015183820152602001620004e1565b838111156200050b576000848401525b50505050565b61045780620005216000396000f3006080604052600436106100825763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630fdcdbfd81146100875780632688454a146100b2578063609d3334146100c757806373d4a13a146100dc578063b6f9cf6a146100f1578063ecc0f87114610106578063f765417614610128575b600080fd5b34801561009357600080fd5b5061009c61014a565b6040516100a9919061039c565b60405180910390f35b3480156100be57600080fd5b5061009c6101d5565b3480156100d357600080fd5b5061009c610230565b3480156100e857600080fd5b5061009c61028a565b3480156100fd57600080fd5b5061009c6102e5565b34801561011257600080fd5b5061011b610340565b6040516100a991906103c8565b34801561013457600080fd5b5061013d610346565b6040516100a991906103b4565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156101cd5780601f106101a2576101008083540402835291602001916101cd565b820191906000526020600020905b8154815290600101906020018083116101b057829003601f168201915b505050505081565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101cd5780601f106101a2576101008083540402835291602001916101cd565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101cd5780601f106101a2576101008083540402835291602001916101cd565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101cd5780601f106101a2576101008083540402835291602001916101cd565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101cd5780601f106101a2576101008083540402835291602001916101cd565b60055490565b60065460070b81565b600061035a826103d6565b80845261036e8160208601602086016103e3565b61037781610413565b9093016020019392505050565b61038d816103da565b82525050565b61038d816103e0565b602080825281016103ad818461034f565b9392505050565b602081016103c28284610384565b92915050565b602081016103c28284610393565b5190565b60070b90565b90565b60005b838110156103fe5781810151838201526020016103e6565b8381111561040d576000848401525b50505050565b601f01601f1916905600a265627a7a72305820599056db1ace1549c7053bcfd2c7fbcbeefbcace5febefa1093e1461eb9210d46c6578706572696d656e74616cf50037`

// DeployWriteRequest deploys a new Ethereum contract, binding an instance of WriteRequest to it.
func DeployWriteRequest(auth *bind.TransactOpts, backend bind.ContractBackend, d []byte, ed []byte, l []byte, p []common.Address, u []byte, cs []byte, s int64) (common.Address, *types.Transaction, *WriteRequest, error) {
	parsed, err := abi.JSON(strings.NewReader(WriteRequestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WriteRequestBin), backend, d, ed, l, p, u, cs, s)
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
const WriteRequestHolderBin = `0x608060405234801561001057600080fd5b5061022d806100206000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631029a017811461007157806342087d4f14610094578063b7312591146100c9578063e216ea2b14610106578063ed85db2314610127575b600080fd5b34801561007d57600080fd5b50610092600160a060020a0360043516610148565b005b3480156100a057600080fd5b506100b5600160a060020a0360043516610195565b604080519115158252519081900360200190f35b3480156100d557600080fd5b506100ea600160a060020a03600435166101b3565b60408051600160a060020a039092168252519081900360200190f35b34801561011257600080fd5b506100b5600160a060020a03600435166101d1565b34801561013357600080fd5b506100ea600160a060020a03600435166101e6565b600160a060020a0316600081815260208181526040808320805473ffffffffffffffffffffffffffffffffffffffff1916909417909355600190819052919020805460ff19169091179055565b600160a060020a031660009081526001602052604090205460ff1690565b600160a060020a039081166000908152602081905260409020541690565b60016020526000908152604090205460ff1681565b600060208190529081526040902054600160a060020a0316815600a165627a7a7230582078ed601c0ffe48ff28c6c456e53e2b4ceb479c61a86ea803dea5c433fb5d80170029`

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

