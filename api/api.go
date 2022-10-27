// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"CloseDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"CreateDeposit\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deposites\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506103f5806100206000396000f3fe6080604052600436106100345760003560e01c8063427544fe14610039578063c35789cc14610076578063ed21248c1461008d575b600080fd5b34801561004557600080fd5b50610060600480360381019061005b919061028e565b610097565b60405161006d91906102d4565b60405180910390f35b34801561008257600080fd5b5061008b6100af565b005b6100956101ad565b005b60006020528060005260406000206000915090505481565b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205403610130576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101279061034c565b60405180910390fd5b7f837d23a798bf91d3f58cf49011204a506b6d8689d70754231081c2c05f081d093360405161015f919061037b565b60405180910390a160008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550565b7f5c60bd3a970d085ac5923a358ca5d94faf619ccc740a0abb1df309108c62cf9933346040516101de929190610396565b60405180910390a1346000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061025b82610230565b9050919050565b61026b81610250565b811461027657600080fd5b50565b60008135905061028881610262565b92915050565b6000602082840312156102a4576102a361022b565b5b60006102b284828501610279565b91505092915050565b6000819050919050565b6102ce816102bb565b82525050565b60006020820190506102e960008301846102c5565b92915050565b600082825260208201905092915050565b7f4e6f742065786973740000000000000000000000000000000000000000000000600082015250565b60006103366009836102ef565b915061034182610300565b602082019050919050565b6000602082019050818103600083015261036581610329565b9050919050565b61037581610250565b82525050565b6000602082019050610390600083018461036c565b92915050565b60006040820190506103ab600083018561036c565b6103b860208301846102c5565b939250505056fea2646970667358221220bb7ead9ececf5d6f6590f6ce634003cff169940c46200a43afbbf20cbd4fe1a364736f6c63430008110033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// Deposites is a free data retrieval call binding the contract method 0x427544fe.
//
// Solidity: function deposites(address ) view returns(uint256)
func (_Api *ApiCaller) Deposites(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "deposites", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deposites is a free data retrieval call binding the contract method 0x427544fe.
//
// Solidity: function deposites(address ) view returns(uint256)
func (_Api *ApiSession) Deposites(arg0 common.Address) (*big.Int, error) {
	return _Api.Contract.Deposites(&_Api.CallOpts, arg0)
}

// Deposites is a free data retrieval call binding the contract method 0x427544fe.
//
// Solidity: function deposites(address ) view returns(uint256)
func (_Api *ApiCallerSession) Deposites(arg0 common.Address) (*big.Int, error) {
	return _Api.Contract.Deposites(&_Api.CallOpts, arg0)
}

// Close is a paid mutator transaction binding the contract method 0xc35789cc.
//
// Solidity: function Close() returns()
func (_Api *ApiTransactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "Close")
}

// Close is a paid mutator transaction binding the contract method 0xc35789cc.
//
// Solidity: function Close() returns()
func (_Api *ApiSession) Close() (*types.Transaction, error) {
	return _Api.Contract.Close(&_Api.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0xc35789cc.
//
// Solidity: function Close() returns()
func (_Api *ApiTransactorSession) Close() (*types.Transaction, error) {
	return _Api.Contract.Close(&_Api.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Api *ApiTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "Deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Api *ApiSession) Deposit() (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xed21248c.
//
// Solidity: function Deposit() payable returns()
func (_Api *ApiTransactorSession) Deposit() (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts)
}

// ApiCloseDepositIterator is returned from FilterCloseDeposit and is used to iterate over the raw logs and unpacked data for CloseDeposit events raised by the Api contract.
type ApiCloseDepositIterator struct {
	Event *ApiCloseDeposit // Event containing the contract specifics and raw log

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
func (it *ApiCloseDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiCloseDeposit)
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
		it.Event = new(ApiCloseDeposit)
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
func (it *ApiCloseDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiCloseDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiCloseDeposit represents a CloseDeposit event raised by the Api contract.
type ApiCloseDeposit struct {
	From common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCloseDeposit is a free log retrieval operation binding the contract event 0x837d23a798bf91d3f58cf49011204a506b6d8689d70754231081c2c05f081d09.
//
// Solidity: event CloseDeposit(address _from)
func (_Api *ApiFilterer) FilterCloseDeposit(opts *bind.FilterOpts) (*ApiCloseDepositIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "CloseDeposit")
	if err != nil {
		return nil, err
	}
	return &ApiCloseDepositIterator{contract: _Api.contract, event: "CloseDeposit", logs: logs, sub: sub}, nil
}

// WatchCloseDeposit is a free log subscription operation binding the contract event 0x837d23a798bf91d3f58cf49011204a506b6d8689d70754231081c2c05f081d09.
//
// Solidity: event CloseDeposit(address _from)
func (_Api *ApiFilterer) WatchCloseDeposit(opts *bind.WatchOpts, sink chan<- *ApiCloseDeposit) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "CloseDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiCloseDeposit)
				if err := _Api.contract.UnpackLog(event, "CloseDeposit", log); err != nil {
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

// ParseCloseDeposit is a log parse operation binding the contract event 0x837d23a798bf91d3f58cf49011204a506b6d8689d70754231081c2c05f081d09.
//
// Solidity: event CloseDeposit(address _from)
func (_Api *ApiFilterer) ParseCloseDeposit(log types.Log) (*ApiCloseDeposit, error) {
	event := new(ApiCloseDeposit)
	if err := _Api.contract.UnpackLog(event, "CloseDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiCreateDepositIterator is returned from FilterCreateDeposit and is used to iterate over the raw logs and unpacked data for CreateDeposit events raised by the Api contract.
type ApiCreateDepositIterator struct {
	Event *ApiCreateDeposit // Event containing the contract specifics and raw log

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
func (it *ApiCreateDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiCreateDeposit)
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
		it.Event = new(ApiCreateDeposit)
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
func (it *ApiCreateDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiCreateDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiCreateDeposit represents a CreateDeposit event raised by the Api contract.
type ApiCreateDeposit struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCreateDeposit is a free log retrieval operation binding the contract event 0x5c60bd3a970d085ac5923a358ca5d94faf619ccc740a0abb1df309108c62cf99.
//
// Solidity: event CreateDeposit(address _from, uint256 _value)
func (_Api *ApiFilterer) FilterCreateDeposit(opts *bind.FilterOpts) (*ApiCreateDepositIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "CreateDeposit")
	if err != nil {
		return nil, err
	}
	return &ApiCreateDepositIterator{contract: _Api.contract, event: "CreateDeposit", logs: logs, sub: sub}, nil
}

// WatchCreateDeposit is a free log subscription operation binding the contract event 0x5c60bd3a970d085ac5923a358ca5d94faf619ccc740a0abb1df309108c62cf99.
//
// Solidity: event CreateDeposit(address _from, uint256 _value)
func (_Api *ApiFilterer) WatchCreateDeposit(opts *bind.WatchOpts, sink chan<- *ApiCreateDeposit) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "CreateDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiCreateDeposit)
				if err := _Api.contract.UnpackLog(event, "CreateDeposit", log); err != nil {
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

// ParseCreateDeposit is a log parse operation binding the contract event 0x5c60bd3a970d085ac5923a358ca5d94faf619ccc740a0abb1df309108c62cf99.
//
// Solidity: event CreateDeposit(address _from, uint256 _value)
func (_Api *ApiFilterer) ParseCreateDeposit(log types.Log) (*ApiCreateDeposit, error) {
	event := new(ApiCreateDeposit)
	if err := _Api.contract.UnpackLog(event, "CreateDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
