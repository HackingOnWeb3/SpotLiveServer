// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package SpotLive

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

// SpotLiveCheckInInfo is an auto generated low-level Go binding around an user-defined struct.
type SpotLiveCheckInInfo struct {
	Latitude      *big.Int
	Longitude     *big.Int
	Time          *big.Int
	Origin        bool
	BelongToScope common.Address
}

// SpotLiveMetaData contains all meta data concerning the SpotLive contract.
var SpotLiveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"userName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"scope\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"scopename\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"addLiveInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scope\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buyPoint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lat\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"long\",\"type\":\"uint256\"}],\"name\":\"checkDistance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scopeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"latitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOrigin\",\"type\":\"bool\"}],\"name\":\"checkIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"checkInList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"latitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"origin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"belongToScope\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scope\",\"type\":\"address\"}],\"name\":\"checkMembers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTokenID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCheckInList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"latitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"origin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"belongToScope\",\"type\":\"address\"}],\"internalType\":\"structSpotLive.CheckInInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getCheckInList\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"latitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"origin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"belongToScope\",\"type\":\"address\"}],\"internalType\":\"structSpotLive.CheckInInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"lat1\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"lon1\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"lat2\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"lon2\",\"type\":\"int256\"}],\"name\":\"getDistance\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOriginList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"live\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getPriceByLive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserTokenList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liveEvenInfoMap\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"liveSupplyMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"memberCheckList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"latitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"origin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"belongToScope\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"scopeAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"location\",\"type\":\"string\"}],\"name\":\"mintPASS\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"originList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"scopeMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"latitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"longitude\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"origin\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"belongToScope\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"scopeMembersMap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"scopeTokenList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdToScope\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userTokenList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SpotLiveABI is the input ABI used to generate the binding from.
// Deprecated: Use SpotLiveMetaData.ABI instead.
var SpotLiveABI = SpotLiveMetaData.ABI

// SpotLive is an auto generated Go binding around an Ethereum contract.
type SpotLive struct {
	SpotLiveCaller     // Read-only binding to the contract
	SpotLiveTransactor // Write-only binding to the contract
	SpotLiveFilterer   // Log filterer for contract events
}

// SpotLiveCaller is an auto generated read-only Go binding around an Ethereum contract.
type SpotLiveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotLiveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SpotLiveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotLiveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SpotLiveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SpotLiveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SpotLiveSession struct {
	Contract     *SpotLive         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SpotLiveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SpotLiveCallerSession struct {
	Contract *SpotLiveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SpotLiveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SpotLiveTransactorSession struct {
	Contract     *SpotLiveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SpotLiveRaw is an auto generated low-level Go binding around an Ethereum contract.
type SpotLiveRaw struct {
	Contract *SpotLive // Generic contract binding to access the raw methods on
}

// SpotLiveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SpotLiveCallerRaw struct {
	Contract *SpotLiveCaller // Generic read-only contract binding to access the raw methods on
}

// SpotLiveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SpotLiveTransactorRaw struct {
	Contract *SpotLiveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSpotLive creates a new instance of SpotLive, bound to a specific deployed contract.
func NewSpotLive(address common.Address, backend bind.ContractBackend) (*SpotLive, error) {
	contract, err := bindSpotLive(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SpotLive{SpotLiveCaller: SpotLiveCaller{contract: contract}, SpotLiveTransactor: SpotLiveTransactor{contract: contract}, SpotLiveFilterer: SpotLiveFilterer{contract: contract}}, nil
}

// NewSpotLiveCaller creates a new read-only instance of SpotLive, bound to a specific deployed contract.
func NewSpotLiveCaller(address common.Address, caller bind.ContractCaller) (*SpotLiveCaller, error) {
	contract, err := bindSpotLive(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SpotLiveCaller{contract: contract}, nil
}

// NewSpotLiveTransactor creates a new write-only instance of SpotLive, bound to a specific deployed contract.
func NewSpotLiveTransactor(address common.Address, transactor bind.ContractTransactor) (*SpotLiveTransactor, error) {
	contract, err := bindSpotLive(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SpotLiveTransactor{contract: contract}, nil
}

// NewSpotLiveFilterer creates a new log filterer instance of SpotLive, bound to a specific deployed contract.
func NewSpotLiveFilterer(address common.Address, filterer bind.ContractFilterer) (*SpotLiveFilterer, error) {
	contract, err := bindSpotLive(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SpotLiveFilterer{contract: contract}, nil
}

// bindSpotLive binds a generic wrapper to an already deployed contract.
func bindSpotLive(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SpotLiveMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SpotLive *SpotLiveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SpotLive.Contract.SpotLiveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SpotLive *SpotLiveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpotLive.Contract.SpotLiveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SpotLive *SpotLiveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SpotLive.Contract.SpotLiveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SpotLive *SpotLiveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SpotLive.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SpotLive *SpotLiveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SpotLive.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SpotLive *SpotLiveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SpotLive.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SpotLive *SpotLiveCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SpotLive *SpotLiveSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SpotLive.Contract.BalanceOf(&_SpotLive.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_SpotLive *SpotLiveCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _SpotLive.Contract.BalanceOf(&_SpotLive.CallOpts, owner)
}

// CheckDistance is a free data retrieval call binding the contract method 0x3b8eee6d.
//
// Solidity: function checkDistance(uint256 lat, uint256 long) view returns(address)
func (_SpotLive *SpotLiveCaller) CheckDistance(opts *bind.CallOpts, lat *big.Int, long *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "checkDistance", lat, long)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckDistance is a free data retrieval call binding the contract method 0x3b8eee6d.
//
// Solidity: function checkDistance(uint256 lat, uint256 long) view returns(address)
func (_SpotLive *SpotLiveSession) CheckDistance(lat *big.Int, long *big.Int) (common.Address, error) {
	return _SpotLive.Contract.CheckDistance(&_SpotLive.CallOpts, lat, long)
}

// CheckDistance is a free data retrieval call binding the contract method 0x3b8eee6d.
//
// Solidity: function checkDistance(uint256 lat, uint256 long) view returns(address)
func (_SpotLive *SpotLiveCallerSession) CheckDistance(lat *big.Int, long *big.Int) (common.Address, error) {
	return _SpotLive.Contract.CheckDistance(&_SpotLive.CallOpts, lat, long)
}

// CheckInList is a free data retrieval call binding the contract method 0xd22e6818.
//
// Solidity: function checkInList(uint256 ) view returns(uint256 latitude, uint256 longitude, uint256 time, bool origin, address belongToScope)
func (_SpotLive *SpotLiveCaller) CheckInList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Latitude      *big.Int
	Longitude     *big.Int
	Time          *big.Int
	Origin        bool
	BelongToScope common.Address
}, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "checkInList", arg0)

	outstruct := new(struct {
		Latitude      *big.Int
		Longitude     *big.Int
		Time          *big.Int
		Origin        bool
		BelongToScope common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Latitude = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Longitude = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Origin = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.BelongToScope = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// CheckInList is a free data retrieval call binding the contract method 0xd22e6818.
//
// Solidity: function checkInList(uint256 ) view returns(uint256 latitude, uint256 longitude, uint256 time, bool origin, address belongToScope)
func (_SpotLive *SpotLiveSession) CheckInList(arg0 *big.Int) (struct {
	Latitude      *big.Int
	Longitude     *big.Int
	Time          *big.Int
	Origin        bool
	BelongToScope common.Address
}, error) {
	return _SpotLive.Contract.CheckInList(&_SpotLive.CallOpts, arg0)
}

// CheckInList is a free data retrieval call binding the contract method 0xd22e6818.
//
// Solidity: function checkInList(uint256 ) view returns(uint256 latitude, uint256 longitude, uint256 time, bool origin, address belongToScope)
func (_SpotLive *SpotLiveCallerSession) CheckInList(arg0 *big.Int) (struct {
	Latitude      *big.Int
	Longitude     *big.Int
	Time          *big.Int
	Origin        bool
	BelongToScope common.Address
}, error) {
	return _SpotLive.Contract.CheckInList(&_SpotLive.CallOpts, arg0)
}

// CheckMembers is a free data retrieval call binding the contract method 0xebd757a4.
//
// Solidity: function checkMembers(address scope) view returns(uint256)
func (_SpotLive *SpotLiveCaller) CheckMembers(opts *bind.CallOpts, scope common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "checkMembers", scope)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckMembers is a free data retrieval call binding the contract method 0xebd757a4.
//
// Solidity: function checkMembers(address scope) view returns(uint256)
func (_SpotLive *SpotLiveSession) CheckMembers(scope common.Address) (*big.Int, error) {
	return _SpotLive.Contract.CheckMembers(&_SpotLive.CallOpts, scope)
}

// CheckMembers is a free data retrieval call binding the contract method 0xebd757a4.
//
// Solidity: function checkMembers(address scope) view returns(uint256)
func (_SpotLive *SpotLiveCallerSession) CheckMembers(scope common.Address) (*big.Int, error) {
	return _SpotLive.Contract.CheckMembers(&_SpotLive.CallOpts, scope)
}

// CurrentTokenID is a free data retrieval call binding the contract method 0xbb62115e.
//
// Solidity: function currentTokenID() view returns(uint256)
func (_SpotLive *SpotLiveCaller) CurrentTokenID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "currentTokenID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTokenID is a free data retrieval call binding the contract method 0xbb62115e.
//
// Solidity: function currentTokenID() view returns(uint256)
func (_SpotLive *SpotLiveSession) CurrentTokenID() (*big.Int, error) {
	return _SpotLive.Contract.CurrentTokenID(&_SpotLive.CallOpts)
}

// CurrentTokenID is a free data retrieval call binding the contract method 0xbb62115e.
//
// Solidity: function currentTokenID() view returns(uint256)
func (_SpotLive *SpotLiveCallerSession) CurrentTokenID() (*big.Int, error) {
	return _SpotLive.Contract.CurrentTokenID(&_SpotLive.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SpotLive *SpotLiveCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SpotLive *SpotLiveSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SpotLive.Contract.GetApproved(&_SpotLive.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_SpotLive *SpotLiveCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _SpotLive.Contract.GetApproved(&_SpotLive.CallOpts, tokenId)
}

// GetCheckInList is a free data retrieval call binding the contract method 0x1b2f3258.
//
// Solidity: function getCheckInList() view returns((uint256,uint256,uint256,bool,address)[], uint256)
func (_SpotLive *SpotLiveCaller) GetCheckInList(opts *bind.CallOpts) ([]SpotLiveCheckInInfo, *big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "getCheckInList")

	if err != nil {
		return *new([]SpotLiveCheckInInfo), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]SpotLiveCheckInInfo)).(*[]SpotLiveCheckInInfo)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCheckInList is a free data retrieval call binding the contract method 0x1b2f3258.
//
// Solidity: function getCheckInList() view returns((uint256,uint256,uint256,bool,address)[], uint256)
func (_SpotLive *SpotLiveSession) GetCheckInList() ([]SpotLiveCheckInInfo, *big.Int, error) {
	return _SpotLive.Contract.GetCheckInList(&_SpotLive.CallOpts)
}

// GetCheckInList is a free data retrieval call binding the contract method 0x1b2f3258.
//
// Solidity: function getCheckInList() view returns((uint256,uint256,uint256,bool,address)[], uint256)
func (_SpotLive *SpotLiveCallerSession) GetCheckInList() ([]SpotLiveCheckInInfo, *big.Int, error) {
	return _SpotLive.Contract.GetCheckInList(&_SpotLive.CallOpts)
}

// GetCheckInList0 is a free data retrieval call binding the contract method 0x56d6d02e.
//
// Solidity: function getCheckInList(address user) view returns((uint256,uint256,uint256,bool,address)[])
func (_SpotLive *SpotLiveCaller) GetCheckInList0(opts *bind.CallOpts, user common.Address) ([]SpotLiveCheckInInfo, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "getCheckInList0", user)

	if err != nil {
		return *new([]SpotLiveCheckInInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]SpotLiveCheckInInfo)).(*[]SpotLiveCheckInInfo)

	return out0, err

}

// GetCheckInList0 is a free data retrieval call binding the contract method 0x56d6d02e.
//
// Solidity: function getCheckInList(address user) view returns((uint256,uint256,uint256,bool,address)[])
func (_SpotLive *SpotLiveSession) GetCheckInList0(user common.Address) ([]SpotLiveCheckInInfo, error) {
	return _SpotLive.Contract.GetCheckInList0(&_SpotLive.CallOpts, user)
}

// GetCheckInList0 is a free data retrieval call binding the contract method 0x56d6d02e.
//
// Solidity: function getCheckInList(address user) view returns((uint256,uint256,uint256,bool,address)[])
func (_SpotLive *SpotLiveCallerSession) GetCheckInList0(user common.Address) ([]SpotLiveCheckInInfo, error) {
	return _SpotLive.Contract.GetCheckInList0(&_SpotLive.CallOpts, user)
}

// GetDistance is a free data retrieval call binding the contract method 0x2098a22a.
//
// Solidity: function getDistance(int256 lat1, int256 lon1, int256 lat2, int256 lon2) pure returns(int256)
func (_SpotLive *SpotLiveCaller) GetDistance(opts *bind.CallOpts, lat1 *big.Int, lon1 *big.Int, lat2 *big.Int, lon2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "getDistance", lat1, lon1, lat2, lon2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDistance is a free data retrieval call binding the contract method 0x2098a22a.
//
// Solidity: function getDistance(int256 lat1, int256 lon1, int256 lat2, int256 lon2) pure returns(int256)
func (_SpotLive *SpotLiveSession) GetDistance(lat1 *big.Int, lon1 *big.Int, lat2 *big.Int, lon2 *big.Int) (*big.Int, error) {
	return _SpotLive.Contract.GetDistance(&_SpotLive.CallOpts, lat1, lon1, lat2, lon2)
}

// GetDistance is a free data retrieval call binding the contract method 0x2098a22a.
//
// Solidity: function getDistance(int256 lat1, int256 lon1, int256 lat2, int256 lon2) pure returns(int256)
func (_SpotLive *SpotLiveCallerSession) GetDistance(lat1 *big.Int, lon1 *big.Int, lat2 *big.Int, lon2 *big.Int) (*big.Int, error) {
	return _SpotLive.Contract.GetDistance(&_SpotLive.CallOpts, lat1, lon1, lat2, lon2)
}

// GetOriginList is a free data retrieval call binding the contract method 0xcdaddc5e.
//
// Solidity: function getOriginList() view returns(address[], uint256)
func (_SpotLive *SpotLiveCaller) GetOriginList(opts *bind.CallOpts) ([]common.Address, *big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "getOriginList")

	if err != nil {
		return *new([]common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetOriginList is a free data retrieval call binding the contract method 0xcdaddc5e.
//
// Solidity: function getOriginList() view returns(address[], uint256)
func (_SpotLive *SpotLiveSession) GetOriginList() ([]common.Address, *big.Int, error) {
	return _SpotLive.Contract.GetOriginList(&_SpotLive.CallOpts)
}

// GetOriginList is a free data retrieval call binding the contract method 0xcdaddc5e.
//
// Solidity: function getOriginList() view returns(address[], uint256)
func (_SpotLive *SpotLiveCallerSession) GetOriginList() ([]common.Address, *big.Int, error) {
	return _SpotLive.Contract.GetOriginList(&_SpotLive.CallOpts)
}

// GetPriceByLive is a free data retrieval call binding the contract method 0x7ba91f7e.
//
// Solidity: function getPriceByLive(address live, uint256 amount) view returns(uint256)
func (_SpotLive *SpotLiveCaller) GetPriceByLive(opts *bind.CallOpts, live common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "getPriceByLive", live, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceByLive is a free data retrieval call binding the contract method 0x7ba91f7e.
//
// Solidity: function getPriceByLive(address live, uint256 amount) view returns(uint256)
func (_SpotLive *SpotLiveSession) GetPriceByLive(live common.Address, amount *big.Int) (*big.Int, error) {
	return _SpotLive.Contract.GetPriceByLive(&_SpotLive.CallOpts, live, amount)
}

// GetPriceByLive is a free data retrieval call binding the contract method 0x7ba91f7e.
//
// Solidity: function getPriceByLive(address live, uint256 amount) view returns(uint256)
func (_SpotLive *SpotLiveCallerSession) GetPriceByLive(live common.Address, amount *big.Int) (*big.Int, error) {
	return _SpotLive.Contract.GetPriceByLive(&_SpotLive.CallOpts, live, amount)
}

// GetUserTokenList is a free data retrieval call binding the contract method 0x54098092.
//
// Solidity: function getUserTokenList(address user) view returns(uint256[], uint256)
func (_SpotLive *SpotLiveCaller) GetUserTokenList(opts *bind.CallOpts, user common.Address) ([]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "getUserTokenList", user)

	if err != nil {
		return *new([]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUserTokenList is a free data retrieval call binding the contract method 0x54098092.
//
// Solidity: function getUserTokenList(address user) view returns(uint256[], uint256)
func (_SpotLive *SpotLiveSession) GetUserTokenList(user common.Address) ([]*big.Int, *big.Int, error) {
	return _SpotLive.Contract.GetUserTokenList(&_SpotLive.CallOpts, user)
}

// GetUserTokenList is a free data retrieval call binding the contract method 0x54098092.
//
// Solidity: function getUserTokenList(address user) view returns(uint256[], uint256)
func (_SpotLive *SpotLiveCallerSession) GetUserTokenList(user common.Address) ([]*big.Int, *big.Int, error) {
	return _SpotLive.Contract.GetUserTokenList(&_SpotLive.CallOpts, user)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SpotLive *SpotLiveCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SpotLive *SpotLiveSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SpotLive.Contract.IsApprovedForAll(&_SpotLive.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_SpotLive *SpotLiveCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _SpotLive.Contract.IsApprovedForAll(&_SpotLive.CallOpts, owner, operator)
}

// LiveEvenInfoMap is a free data retrieval call binding the contract method 0xe66ac694.
//
// Solidity: function liveEvenInfoMap(address ) view returns(string name, string symbol, uint256 startTime, uint256 endTime, string description)
func (_SpotLive *SpotLiveCaller) LiveEvenInfoMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name        string
	Symbol      string
	StartTime   *big.Int
	EndTime     *big.Int
	Description string
}, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "liveEvenInfoMap", arg0)

	outstruct := new(struct {
		Name        string
		Symbol      string
		StartTime   *big.Int
		EndTime     *big.Int
		Description string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Description = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// LiveEvenInfoMap is a free data retrieval call binding the contract method 0xe66ac694.
//
// Solidity: function liveEvenInfoMap(address ) view returns(string name, string symbol, uint256 startTime, uint256 endTime, string description)
func (_SpotLive *SpotLiveSession) LiveEvenInfoMap(arg0 common.Address) (struct {
	Name        string
	Symbol      string
	StartTime   *big.Int
	EndTime     *big.Int
	Description string
}, error) {
	return _SpotLive.Contract.LiveEvenInfoMap(&_SpotLive.CallOpts, arg0)
}

// LiveEvenInfoMap is a free data retrieval call binding the contract method 0xe66ac694.
//
// Solidity: function liveEvenInfoMap(address ) view returns(string name, string symbol, uint256 startTime, uint256 endTime, string description)
func (_SpotLive *SpotLiveCallerSession) LiveEvenInfoMap(arg0 common.Address) (struct {
	Name        string
	Symbol      string
	StartTime   *big.Int
	EndTime     *big.Int
	Description string
}, error) {
	return _SpotLive.Contract.LiveEvenInfoMap(&_SpotLive.CallOpts, arg0)
}

// LiveSupplyMap is a free data retrieval call binding the contract method 0xa73b84bf.
//
// Solidity: function liveSupplyMap(address ) view returns(uint256)
func (_SpotLive *SpotLiveCaller) LiveSupplyMap(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "liveSupplyMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiveSupplyMap is a free data retrieval call binding the contract method 0xa73b84bf.
//
// Solidity: function liveSupplyMap(address ) view returns(uint256)
func (_SpotLive *SpotLiveSession) LiveSupplyMap(arg0 common.Address) (*big.Int, error) {
	return _SpotLive.Contract.LiveSupplyMap(&_SpotLive.CallOpts, arg0)
}

// LiveSupplyMap is a free data retrieval call binding the contract method 0xa73b84bf.
//
// Solidity: function liveSupplyMap(address ) view returns(uint256)
func (_SpotLive *SpotLiveCallerSession) LiveSupplyMap(arg0 common.Address) (*big.Int, error) {
	return _SpotLive.Contract.LiveSupplyMap(&_SpotLive.CallOpts, arg0)
}

// MaxPoints is a free data retrieval call binding the contract method 0x9a65e3eb.
//
// Solidity: function maxPoints() view returns(uint256)
func (_SpotLive *SpotLiveCaller) MaxPoints(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "maxPoints")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPoints is a free data retrieval call binding the contract method 0x9a65e3eb.
//
// Solidity: function maxPoints() view returns(uint256)
func (_SpotLive *SpotLiveSession) MaxPoints() (*big.Int, error) {
	return _SpotLive.Contract.MaxPoints(&_SpotLive.CallOpts)
}

// MaxPoints is a free data retrieval call binding the contract method 0x9a65e3eb.
//
// Solidity: function maxPoints() view returns(uint256)
func (_SpotLive *SpotLiveCallerSession) MaxPoints() (*big.Int, error) {
	return _SpotLive.Contract.MaxPoints(&_SpotLive.CallOpts)
}

// MemberCheckList is a free data retrieval call binding the contract method 0x96cd0736.
//
// Solidity: function memberCheckList(address , uint256 ) view returns(uint256 latitude, uint256 longitude, uint256 time, bool origin, address belongToScope)
func (_SpotLive *SpotLiveCaller) MemberCheckList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Latitude      *big.Int
	Longitude     *big.Int
	Time          *big.Int
	Origin        bool
	BelongToScope common.Address
}, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "memberCheckList", arg0, arg1)

	outstruct := new(struct {
		Latitude      *big.Int
		Longitude     *big.Int
		Time          *big.Int
		Origin        bool
		BelongToScope common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Latitude = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Longitude = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Origin = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.BelongToScope = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// MemberCheckList is a free data retrieval call binding the contract method 0x96cd0736.
//
// Solidity: function memberCheckList(address , uint256 ) view returns(uint256 latitude, uint256 longitude, uint256 time, bool origin, address belongToScope)
func (_SpotLive *SpotLiveSession) MemberCheckList(arg0 common.Address, arg1 *big.Int) (struct {
	Latitude      *big.Int
	Longitude     *big.Int
	Time          *big.Int
	Origin        bool
	BelongToScope common.Address
}, error) {
	return _SpotLive.Contract.MemberCheckList(&_SpotLive.CallOpts, arg0, arg1)
}

// MemberCheckList is a free data retrieval call binding the contract method 0x96cd0736.
//
// Solidity: function memberCheckList(address , uint256 ) view returns(uint256 latitude, uint256 longitude, uint256 time, bool origin, address belongToScope)
func (_SpotLive *SpotLiveCallerSession) MemberCheckList(arg0 common.Address, arg1 *big.Int) (struct {
	Latitude      *big.Int
	Longitude     *big.Int
	Time          *big.Int
	Origin        bool
	BelongToScope common.Address
}, error) {
	return _SpotLive.Contract.MemberCheckList(&_SpotLive.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SpotLive *SpotLiveCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SpotLive *SpotLiveSession) Name() (string, error) {
	return _SpotLive.Contract.Name(&_SpotLive.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SpotLive *SpotLiveCallerSession) Name() (string, error) {
	return _SpotLive.Contract.Name(&_SpotLive.CallOpts)
}

// OriginList is a free data retrieval call binding the contract method 0x3f7f68fb.
//
// Solidity: function originList(uint256 ) view returns(address)
func (_SpotLive *SpotLiveCaller) OriginList(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "originList", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OriginList is a free data retrieval call binding the contract method 0x3f7f68fb.
//
// Solidity: function originList(uint256 ) view returns(address)
func (_SpotLive *SpotLiveSession) OriginList(arg0 *big.Int) (common.Address, error) {
	return _SpotLive.Contract.OriginList(&_SpotLive.CallOpts, arg0)
}

// OriginList is a free data retrieval call binding the contract method 0x3f7f68fb.
//
// Solidity: function originList(uint256 ) view returns(address)
func (_SpotLive *SpotLiveCallerSession) OriginList(arg0 *big.Int) (common.Address, error) {
	return _SpotLive.Contract.OriginList(&_SpotLive.CallOpts, arg0)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SpotLive *SpotLiveCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SpotLive *SpotLiveSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SpotLive.Contract.OwnerOf(&_SpotLive.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_SpotLive *SpotLiveCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _SpotLive.Contract.OwnerOf(&_SpotLive.CallOpts, tokenId)
}

// ScopeMap is a free data retrieval call binding the contract method 0xb5809224.
//
// Solidity: function scopeMap(address ) view returns(uint256 latitude, uint256 longitude, uint256 time, bool origin, address belongToScope)
func (_SpotLive *SpotLiveCaller) ScopeMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Latitude      *big.Int
	Longitude     *big.Int
	Time          *big.Int
	Origin        bool
	BelongToScope common.Address
}, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "scopeMap", arg0)

	outstruct := new(struct {
		Latitude      *big.Int
		Longitude     *big.Int
		Time          *big.Int
		Origin        bool
		BelongToScope common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Latitude = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Longitude = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Origin = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.BelongToScope = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// ScopeMap is a free data retrieval call binding the contract method 0xb5809224.
//
// Solidity: function scopeMap(address ) view returns(uint256 latitude, uint256 longitude, uint256 time, bool origin, address belongToScope)
func (_SpotLive *SpotLiveSession) ScopeMap(arg0 common.Address) (struct {
	Latitude      *big.Int
	Longitude     *big.Int
	Time          *big.Int
	Origin        bool
	BelongToScope common.Address
}, error) {
	return _SpotLive.Contract.ScopeMap(&_SpotLive.CallOpts, arg0)
}

// ScopeMap is a free data retrieval call binding the contract method 0xb5809224.
//
// Solidity: function scopeMap(address ) view returns(uint256 latitude, uint256 longitude, uint256 time, bool origin, address belongToScope)
func (_SpotLive *SpotLiveCallerSession) ScopeMap(arg0 common.Address) (struct {
	Latitude      *big.Int
	Longitude     *big.Int
	Time          *big.Int
	Origin        bool
	BelongToScope common.Address
}, error) {
	return _SpotLive.Contract.ScopeMap(&_SpotLive.CallOpts, arg0)
}

// ScopeMembersMap is a free data retrieval call binding the contract method 0x25ce9619.
//
// Solidity: function scopeMembersMap(address , uint256 ) view returns(address)
func (_SpotLive *SpotLiveCaller) ScopeMembersMap(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "scopeMembersMap", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ScopeMembersMap is a free data retrieval call binding the contract method 0x25ce9619.
//
// Solidity: function scopeMembersMap(address , uint256 ) view returns(address)
func (_SpotLive *SpotLiveSession) ScopeMembersMap(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _SpotLive.Contract.ScopeMembersMap(&_SpotLive.CallOpts, arg0, arg1)
}

// ScopeMembersMap is a free data retrieval call binding the contract method 0x25ce9619.
//
// Solidity: function scopeMembersMap(address , uint256 ) view returns(address)
func (_SpotLive *SpotLiveCallerSession) ScopeMembersMap(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _SpotLive.Contract.ScopeMembersMap(&_SpotLive.CallOpts, arg0, arg1)
}

// ScopeTokenList is a free data retrieval call binding the contract method 0xb70b2d29.
//
// Solidity: function scopeTokenList(address , uint256 ) view returns(uint256)
func (_SpotLive *SpotLiveCaller) ScopeTokenList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "scopeTokenList", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScopeTokenList is a free data retrieval call binding the contract method 0xb70b2d29.
//
// Solidity: function scopeTokenList(address , uint256 ) view returns(uint256)
func (_SpotLive *SpotLiveSession) ScopeTokenList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SpotLive.Contract.ScopeTokenList(&_SpotLive.CallOpts, arg0, arg1)
}

// ScopeTokenList is a free data retrieval call binding the contract method 0xb70b2d29.
//
// Solidity: function scopeTokenList(address , uint256 ) view returns(uint256)
func (_SpotLive *SpotLiveCallerSession) ScopeTokenList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SpotLive.Contract.ScopeTokenList(&_SpotLive.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SpotLive *SpotLiveCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SpotLive *SpotLiveSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SpotLive.Contract.SupportsInterface(&_SpotLive.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SpotLive *SpotLiveCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SpotLive.Contract.SupportsInterface(&_SpotLive.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SpotLive *SpotLiveCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SpotLive *SpotLiveSession) Symbol() (string, error) {
	return _SpotLive.Contract.Symbol(&_SpotLive.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SpotLive *SpotLiveCallerSession) Symbol() (string, error) {
	return _SpotLive.Contract.Symbol(&_SpotLive.CallOpts)
}

// TokenIdToScope is a free data retrieval call binding the contract method 0xb76bfc90.
//
// Solidity: function tokenIdToScope(uint256 ) view returns(address)
func (_SpotLive *SpotLiveCaller) TokenIdToScope(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "tokenIdToScope", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenIdToScope is a free data retrieval call binding the contract method 0xb76bfc90.
//
// Solidity: function tokenIdToScope(uint256 ) view returns(address)
func (_SpotLive *SpotLiveSession) TokenIdToScope(arg0 *big.Int) (common.Address, error) {
	return _SpotLive.Contract.TokenIdToScope(&_SpotLive.CallOpts, arg0)
}

// TokenIdToScope is a free data retrieval call binding the contract method 0xb76bfc90.
//
// Solidity: function tokenIdToScope(uint256 ) view returns(address)
func (_SpotLive *SpotLiveCallerSession) TokenIdToScope(arg0 *big.Int) (common.Address, error) {
	return _SpotLive.Contract.TokenIdToScope(&_SpotLive.CallOpts, arg0)
}

// TokenPoints is a free data retrieval call binding the contract method 0x57e97340.
//
// Solidity: function tokenPoints(uint256 ) view returns(uint256)
func (_SpotLive *SpotLiveCaller) TokenPoints(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "tokenPoints", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenPoints is a free data retrieval call binding the contract method 0x57e97340.
//
// Solidity: function tokenPoints(uint256 ) view returns(uint256)
func (_SpotLive *SpotLiveSession) TokenPoints(arg0 *big.Int) (*big.Int, error) {
	return _SpotLive.Contract.TokenPoints(&_SpotLive.CallOpts, arg0)
}

// TokenPoints is a free data retrieval call binding the contract method 0x57e97340.
//
// Solidity: function tokenPoints(uint256 ) view returns(uint256)
func (_SpotLive *SpotLiveCallerSession) TokenPoints(arg0 *big.Int) (*big.Int, error) {
	return _SpotLive.Contract.TokenPoints(&_SpotLive.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SpotLive *SpotLiveCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SpotLive *SpotLiveSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SpotLive.Contract.TokenURI(&_SpotLive.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_SpotLive *SpotLiveCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _SpotLive.Contract.TokenURI(&_SpotLive.CallOpts, tokenId)
}

// UserTokenList is a free data retrieval call binding the contract method 0xd00cd650.
//
// Solidity: function userTokenList(address , uint256 ) view returns(uint256)
func (_SpotLive *SpotLiveCaller) UserTokenList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SpotLive.contract.Call(opts, &out, "userTokenList", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTokenList is a free data retrieval call binding the contract method 0xd00cd650.
//
// Solidity: function userTokenList(address , uint256 ) view returns(uint256)
func (_SpotLive *SpotLiveSession) UserTokenList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SpotLive.Contract.UserTokenList(&_SpotLive.CallOpts, arg0, arg1)
}

// UserTokenList is a free data retrieval call binding the contract method 0xd00cd650.
//
// Solidity: function userTokenList(address , uint256 ) view returns(uint256)
func (_SpotLive *SpotLiveCallerSession) UserTokenList(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _SpotLive.Contract.UserTokenList(&_SpotLive.CallOpts, arg0, arg1)
}

// AddLiveInfo is a paid mutator transaction binding the contract method 0x704cd899.
//
// Solidity: function addLiveInfo(string userName, address scope, string scopename, uint256 startTime, uint256 endTime, string description) returns(bool)
func (_SpotLive *SpotLiveTransactor) AddLiveInfo(opts *bind.TransactOpts, userName string, scope common.Address, scopename string, startTime *big.Int, endTime *big.Int, description string) (*types.Transaction, error) {
	return _SpotLive.contract.Transact(opts, "addLiveInfo", userName, scope, scopename, startTime, endTime, description)
}

// AddLiveInfo is a paid mutator transaction binding the contract method 0x704cd899.
//
// Solidity: function addLiveInfo(string userName, address scope, string scopename, uint256 startTime, uint256 endTime, string description) returns(bool)
func (_SpotLive *SpotLiveSession) AddLiveInfo(userName string, scope common.Address, scopename string, startTime *big.Int, endTime *big.Int, description string) (*types.Transaction, error) {
	return _SpotLive.Contract.AddLiveInfo(&_SpotLive.TransactOpts, userName, scope, scopename, startTime, endTime, description)
}

// AddLiveInfo is a paid mutator transaction binding the contract method 0x704cd899.
//
// Solidity: function addLiveInfo(string userName, address scope, string scopename, uint256 startTime, uint256 endTime, string description) returns(bool)
func (_SpotLive *SpotLiveTransactorSession) AddLiveInfo(userName string, scope common.Address, scopename string, startTime *big.Int, endTime *big.Int, description string) (*types.Transaction, error) {
	return _SpotLive.Contract.AddLiveInfo(&_SpotLive.TransactOpts, userName, scope, scopename, startTime, endTime, description)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SpotLive *SpotLiveTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SpotLive *SpotLiveSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.Contract.Approve(&_SpotLive.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_SpotLive *SpotLiveTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.Contract.Approve(&_SpotLive.TransactOpts, to, tokenId)
}

// BuyPoint is a paid mutator transaction binding the contract method 0xe77c0745.
//
// Solidity: function buyPoint(address scope, uint256 tokenId) payable returns()
func (_SpotLive *SpotLiveTransactor) BuyPoint(opts *bind.TransactOpts, scope common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.contract.Transact(opts, "buyPoint", scope, tokenId)
}

// BuyPoint is a paid mutator transaction binding the contract method 0xe77c0745.
//
// Solidity: function buyPoint(address scope, uint256 tokenId) payable returns()
func (_SpotLive *SpotLiveSession) BuyPoint(scope common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.Contract.BuyPoint(&_SpotLive.TransactOpts, scope, tokenId)
}

// BuyPoint is a paid mutator transaction binding the contract method 0xe77c0745.
//
// Solidity: function buyPoint(address scope, uint256 tokenId) payable returns()
func (_SpotLive *SpotLiveTransactorSession) BuyPoint(scope common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.Contract.BuyPoint(&_SpotLive.TransactOpts, scope, tokenId)
}

// CheckIn is a paid mutator transaction binding the contract method 0x4501fc6e.
//
// Solidity: function checkIn(address scopeAddress, uint256 latitude, uint256 longitude, uint256 time, bool isOrigin) returns()
func (_SpotLive *SpotLiveTransactor) CheckIn(opts *bind.TransactOpts, scopeAddress common.Address, latitude *big.Int, longitude *big.Int, time *big.Int, isOrigin bool) (*types.Transaction, error) {
	return _SpotLive.contract.Transact(opts, "checkIn", scopeAddress, latitude, longitude, time, isOrigin)
}

// CheckIn is a paid mutator transaction binding the contract method 0x4501fc6e.
//
// Solidity: function checkIn(address scopeAddress, uint256 latitude, uint256 longitude, uint256 time, bool isOrigin) returns()
func (_SpotLive *SpotLiveSession) CheckIn(scopeAddress common.Address, latitude *big.Int, longitude *big.Int, time *big.Int, isOrigin bool) (*types.Transaction, error) {
	return _SpotLive.Contract.CheckIn(&_SpotLive.TransactOpts, scopeAddress, latitude, longitude, time, isOrigin)
}

// CheckIn is a paid mutator transaction binding the contract method 0x4501fc6e.
//
// Solidity: function checkIn(address scopeAddress, uint256 latitude, uint256 longitude, uint256 time, bool isOrigin) returns()
func (_SpotLive *SpotLiveTransactorSession) CheckIn(scopeAddress common.Address, latitude *big.Int, longitude *big.Int, time *big.Int, isOrigin bool) (*types.Transaction, error) {
	return _SpotLive.Contract.CheckIn(&_SpotLive.TransactOpts, scopeAddress, latitude, longitude, time, isOrigin)
}

// MintPASS is a paid mutator transaction binding the contract method 0xc0a09560.
//
// Solidity: function mintPASS(address scopeAddress, uint256 amount, string location) returns()
func (_SpotLive *SpotLiveTransactor) MintPASS(opts *bind.TransactOpts, scopeAddress common.Address, amount *big.Int, location string) (*types.Transaction, error) {
	return _SpotLive.contract.Transact(opts, "mintPASS", scopeAddress, amount, location)
}

// MintPASS is a paid mutator transaction binding the contract method 0xc0a09560.
//
// Solidity: function mintPASS(address scopeAddress, uint256 amount, string location) returns()
func (_SpotLive *SpotLiveSession) MintPASS(scopeAddress common.Address, amount *big.Int, location string) (*types.Transaction, error) {
	return _SpotLive.Contract.MintPASS(&_SpotLive.TransactOpts, scopeAddress, amount, location)
}

// MintPASS is a paid mutator transaction binding the contract method 0xc0a09560.
//
// Solidity: function mintPASS(address scopeAddress, uint256 amount, string location) returns()
func (_SpotLive *SpotLiveTransactorSession) MintPASS(scopeAddress common.Address, amount *big.Int, location string) (*types.Transaction, error) {
	return _SpotLive.Contract.MintPASS(&_SpotLive.TransactOpts, scopeAddress, amount, location)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SpotLive *SpotLiveTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SpotLive *SpotLiveSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.Contract.SafeTransferFrom(&_SpotLive.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_SpotLive *SpotLiveTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.Contract.SafeTransferFrom(&_SpotLive.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SpotLive *SpotLiveTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SpotLive.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SpotLive *SpotLiveSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SpotLive.Contract.SafeTransferFrom0(&_SpotLive.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_SpotLive *SpotLiveTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _SpotLive.Contract.SafeTransferFrom0(&_SpotLive.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SpotLive *SpotLiveTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _SpotLive.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SpotLive *SpotLiveSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SpotLive.Contract.SetApprovalForAll(&_SpotLive.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_SpotLive *SpotLiveTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _SpotLive.Contract.SetApprovalForAll(&_SpotLive.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SpotLive *SpotLiveTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SpotLive *SpotLiveSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.Contract.TransferFrom(&_SpotLive.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_SpotLive *SpotLiveTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _SpotLive.Contract.TransferFrom(&_SpotLive.TransactOpts, from, to, tokenId)
}

// SpotLiveApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SpotLive contract.
type SpotLiveApprovalIterator struct {
	Event *SpotLiveApproval // Event containing the contract specifics and raw log

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
func (it *SpotLiveApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotLiveApproval)
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
		it.Event = new(SpotLiveApproval)
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
func (it *SpotLiveApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotLiveApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotLiveApproval represents a Approval event raised by the SpotLive contract.
type SpotLiveApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SpotLive *SpotLiveFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*SpotLiveApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SpotLive.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotLiveApprovalIterator{contract: _SpotLive.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SpotLive *SpotLiveFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SpotLiveApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SpotLive.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotLiveApproval)
				if err := _SpotLive.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_SpotLive *SpotLiveFilterer) ParseApproval(log types.Log) (*SpotLiveApproval, error) {
	event := new(SpotLiveApproval)
	if err := _SpotLive.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotLiveApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the SpotLive contract.
type SpotLiveApprovalForAllIterator struct {
	Event *SpotLiveApprovalForAll // Event containing the contract specifics and raw log

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
func (it *SpotLiveApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotLiveApprovalForAll)
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
		it.Event = new(SpotLiveApprovalForAll)
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
func (it *SpotLiveApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotLiveApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotLiveApprovalForAll represents a ApprovalForAll event raised by the SpotLive contract.
type SpotLiveApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SpotLive *SpotLiveFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*SpotLiveApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SpotLive.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &SpotLiveApprovalForAllIterator{contract: _SpotLive.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SpotLive *SpotLiveFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *SpotLiveApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _SpotLive.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotLiveApprovalForAll)
				if err := _SpotLive.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_SpotLive *SpotLiveFilterer) ParseApprovalForAll(log types.Log) (*SpotLiveApprovalForAll, error) {
	event := new(SpotLiveApprovalForAll)
	if err := _SpotLive.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotLiveBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the SpotLive contract.
type SpotLiveBatchMetadataUpdateIterator struct {
	Event *SpotLiveBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *SpotLiveBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotLiveBatchMetadataUpdate)
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
		it.Event = new(SpotLiveBatchMetadataUpdate)
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
func (it *SpotLiveBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotLiveBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotLiveBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the SpotLive contract.
type SpotLiveBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_SpotLive *SpotLiveFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*SpotLiveBatchMetadataUpdateIterator, error) {

	logs, sub, err := _SpotLive.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &SpotLiveBatchMetadataUpdateIterator{contract: _SpotLive.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_SpotLive *SpotLiveFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *SpotLiveBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _SpotLive.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotLiveBatchMetadataUpdate)
				if err := _SpotLive.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_SpotLive *SpotLiveFilterer) ParseBatchMetadataUpdate(log types.Log) (*SpotLiveBatchMetadataUpdate, error) {
	event := new(SpotLiveBatchMetadataUpdate)
	if err := _SpotLive.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotLiveMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the SpotLive contract.
type SpotLiveMetadataUpdateIterator struct {
	Event *SpotLiveMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *SpotLiveMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotLiveMetadataUpdate)
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
		it.Event = new(SpotLiveMetadataUpdate)
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
func (it *SpotLiveMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotLiveMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotLiveMetadataUpdate represents a MetadataUpdate event raised by the SpotLive contract.
type SpotLiveMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_SpotLive *SpotLiveFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*SpotLiveMetadataUpdateIterator, error) {

	logs, sub, err := _SpotLive.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &SpotLiveMetadataUpdateIterator{contract: _SpotLive.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_SpotLive *SpotLiveFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *SpotLiveMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _SpotLive.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotLiveMetadataUpdate)
				if err := _SpotLive.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_SpotLive *SpotLiveFilterer) ParseMetadataUpdate(log types.Log) (*SpotLiveMetadataUpdate, error) {
	event := new(SpotLiveMetadataUpdate)
	if err := _SpotLive.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SpotLiveTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SpotLive contract.
type SpotLiveTransferIterator struct {
	Event *SpotLiveTransfer // Event containing the contract specifics and raw log

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
func (it *SpotLiveTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SpotLiveTransfer)
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
		it.Event = new(SpotLiveTransfer)
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
func (it *SpotLiveTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SpotLiveTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SpotLiveTransfer represents a Transfer event raised by the SpotLive contract.
type SpotLiveTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SpotLive *SpotLiveFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*SpotLiveTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SpotLive.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &SpotLiveTransferIterator{contract: _SpotLive.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SpotLive *SpotLiveFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SpotLiveTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _SpotLive.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SpotLiveTransfer)
				if err := _SpotLive.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_SpotLive *SpotLiveFilterer) ParseTransfer(log types.Log) (*SpotLiveTransfer, error) {
	event := new(SpotLiveTransfer)
	if err := _SpotLive.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
