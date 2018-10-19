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
const CalypsoABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addWriteRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"canWrite\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"checkIfRequestIsValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"isValidReadRequest\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"AddReadRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"holder\",\"type\":\"address\"},{\"name\":\"rrholder\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// CalypsoBin is the compiled bytecode used for deploying new contracts.
const CalypsoBin = `0x608060405234801561001057600080fd5b506040516060806105fe83398101604090815281516020830151919092015160008054600160a060020a0319908116600160a060020a0395861617825560018054821694861694909417909355600380549093169390911692909217905561058090819061007e90396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631029a01781146100715780633239b71b146100945780634a4e19d3146100c9578063778ad09c146100ea578063bdae9b891461010b575b600080fd5b34801561007d57600080fd5b50610092600160a060020a036004351661012c565b005b3480156100a057600080fd5b506100b5600160a060020a03600435166101ae565b604080519115158252519081900360200190f35b3480156100d557600080fd5b506100b5600160a060020a0360043516610248565b3480156100f657600080fd5b506100b5600160a060020a03600435166102b3565b34801561011757600080fd5b50610092600160a060020a036004351661031e565b600154604080517f1029a017000000000000000000000000000000000000000000000000000000008152600160a060020a03848116600483015291519190921691631029a01791602480830192600092919082900301818387803b15801561019357600080fd5b505af11580156101a7573d6000803e3d6000fd5b5050505050565b60008054604080517f3239b71b000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015291519190921691633239b71b91602480830192602092919082900301818787803b15801561021657600080fd5b505af115801561022a573d6000803e3d6000fd5b505050506040513d602081101561024057600080fd5b505192915050565b600154604080517f42087d4f000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916342087d4f9160248082019260209290919082900301818787803b15801561021657600080fd5b600354604080517ffd6336dc000000000000000000000000000000000000000000000000000000008152600160a060020a0384811660048301529151600093929092169163fd6336dc9160248082019260209290919082900301818787803b15801561021657600080fd5b60008082915081600160a060020a03166364e94e676040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561037b57600080fd5b505af115801561038f573d6000803e3d6000fd5b505050506040513d60208110156103a557600080fd5b5051600154604080517f42087d4f000000000000000000000000000000000000000000000000000000008152600160a060020a03808516600483015291519394509116916342087d4f916024808201926020929091908290030181600087803b15801561041157600080fd5b505af1158015610425573d6000803e3d6000fd5b505050506040513d602081101561043b57600080fd5b505115156104d057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f5468657265206973206e6f20636f72726573706f6e64696e672077726974652060448201527f7265717565737400000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b600354604080517f31419b50000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152915191909216916331419b5091602480830192600092919082900301818387803b15801561053757600080fd5b505af115801561054b573d6000803e3d6000fd5b505050505050505600a165627a7a72305820038b4b14c339549642db480fecb24c4101b70b0648cca50f873800dc83eeb6360029`

// DeployCalypso deploys a new Ethereum contract, binding an instance of Calypso to it.
func DeployCalypso(auth *bind.TransactOpts, backend bind.ContractBackend, a common.Address, holder common.Address, rrholder common.Address) (common.Address, *types.Transaction, *Calypso, error) {
	parsed, err := abi.JSON(strings.NewReader(CalypsoABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CalypsoBin), backend, a, holder, rrholder)
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
const OwnersABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"canWrite\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNrOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"a\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnersBin is the compiled bytecode used for deploying new contracts.
const OwnersBin = `0x608060405234801561001057600080fd5b5060405161025338038061025383398101604052805101805160009061003c90829060208501906100ab565b50600090505b81518163ffffffff1610156100a4576001806000848463ffffffff1681518110151561006a57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff1916911515919091179055600101610042565b5050610137565b828054828255906000526020600020908101928215610100579160200282015b828111156101005782518254600160a060020a031916600160a060020a039091161782556020909201916001909101906100cb565b5061010c929150610110565b5090565b61013491905b8082111561010c578054600160a060020a0319168155600101610116565b90565b61010d806101466000396000f30060806040526004361060485763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633239b71b8114604d578063502bb73414608c575b600080fd5b348015605857600080fd5b50607873ffffffffffffffffffffffffffffffffffffffff6004351660b0565b604080519115158252519081900360200190f35b348015609757600080fd5b50609e60db565b60408051918252519081900360200190f35b73ffffffffffffffffffffffffffffffffffffffff1660009081526001602052604090205460ff1690565b600054905600a165627a7a72305820c3ea7a8893b7712755ea162e0cdce2890fea3d7272e025e2dda8403162171c830029`

// DeployOwners deploys a new Ethereum contract, binding an instance of Owners to it.
func DeployOwners(auth *bind.TransactOpts, backend bind.ContractBackend, a []common.Address) (common.Address, *types.Transaction, *Owners, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnersABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnersBin), backend, a)
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
const WriteRequestABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"LTSID\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"U\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"extraData\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getPolicySize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"d\",\"type\":\"bytes\"},{\"name\":\"ed\",\"type\":\"bytes\"},{\"name\":\"l\",\"type\":\"bytes\"},{\"name\":\"p\",\"type\":\"address[]\"},{\"name\":\"u\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// WriteRequestBin is the compiled bytecode used for deploying new contracts.
const WriteRequestBin = `0x608060405234801561001057600080fd5b5060405161051c38038061051c8339810160409081528151602080840151928401516060850151608086015193860180519096958601959283019491830193919092019161006491600091908801906100bf565b5083516100789060019060208701906100bf565b50825161008c9060029060208601906100bf565b5081516100a090600490602085019061013d565b5080516100b49060039060208401906100bf565b5050505050506101df565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061010057805160ff191683800117855561012d565b8280016001018555821561012d579182015b8281111561012d578251825591602001919060010190610112565b5061013992915061019e565b5090565b828054828255906000526020600020908101928215610192579160200282015b828111156101925782518254600160a060020a031916600160a060020a0390911617825560209092019160019091019061015d565b506101399291506101bb565b6101b891905b8082111561013957600081556001016101a4565b90565b6101b891905b80821115610139578054600160a060020a03191681556001016101c1565b61032e806101ee6000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630fdcdbfd81146100715780632688454a146100fb578063609d33341461011057806373d4a13a14610125578063ecc0f8711461013a575b600080fd5b34801561007d57600080fd5b50610086610161565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100c05781810151838201526020016100a8565b50505050905090810190601f1680156100ed5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561010757600080fd5b506100866101ec565b34801561011c57600080fd5b50610086610247565b34801561013157600080fd5b506100866102a1565b34801561014657600080fd5b5061014f6102fc565b60408051918252519081900360200190f35b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156101e45780601f106101b9576101008083540402835291602001916101e4565b820191906000526020600020905b8154815290600101906020018083116101c757829003601f168201915b505050505081565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101e45780601f106101b9576101008083540402835291602001916101e4565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101e45780601f106101b9576101008083540402835291602001916101e4565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101e45780601f106101b9576101008083540402835291602001916101e4565b600454905600a165627a7a7230582092ec765cb2494f80c74f404956fd425d9e4fb9730e89bbc4c8904a9de00b0e860029`

// DeployWriteRequest deploys a new Ethereum contract, binding an instance of WriteRequest to it.
func DeployWriteRequest(auth *bind.TransactOpts, backend bind.ContractBackend, d []byte, ed []byte, l []byte, p []common.Address, u []byte) (common.Address, *types.Transaction, *WriteRequest, error) {
	parsed, err := abi.JSON(strings.NewReader(WriteRequestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WriteRequestBin), backend, d, ed, l, p, u)
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
const WriteRequestHolderBin = `0x608060405234801561001057600080fd5b5061022d806100206000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631029a017811461007157806342087d4f14610094578063b7312591146100c9578063e216ea2b14610106578063ed85db2314610127575b600080fd5b34801561007d57600080fd5b50610092600160a060020a0360043516610148565b005b3480156100a057600080fd5b506100b5600160a060020a0360043516610195565b604080519115158252519081900360200190f35b3480156100d557600080fd5b506100ea600160a060020a03600435166101b3565b60408051600160a060020a039092168252519081900360200190f35b34801561011257600080fd5b506100b5600160a060020a03600435166101d1565b34801561013357600080fd5b506100ea600160a060020a03600435166101e6565b600160a060020a0316600081815260208181526040808320805473ffffffffffffffffffffffffffffffffffffffff1916909417909355600190819052919020805460ff19169091179055565b600160a060020a031660009081526001602052604090205460ff1690565b600160a060020a039081166000908152602081905260409020541690565b60016020526000908152604090205460ff1681565b600060208190529081526040902054600160a060020a0316815600a165627a7a72305820dd6627bf8943d11fb96afc899f8cecb0a597f83120db82ac58d1ff8b36146fdd0029`

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

