// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ogwhitelist

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

// OgwhitelistMetaData contains all meta data concerning the Ogwhitelist contract.
var OgwhitelistMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MAX_SUPPLY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OG_MINT_PRICE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUBLIC_MINT_PRICE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addressMintedBalance\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getApproved\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isApprovedForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"merkleRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ogMint\",\"inputs\":[{\"name\":\"_merkleProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownerOf\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publicMint\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"remainingSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"safeTransferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setApprovalForAll\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenURI\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalMinted\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMerkleRoot\",\"inputs\":[{\"name\":\"_newMerkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verifyWhitelistStatus\",\"inputs\":[{\"name\":\"_merkleProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ApprovalForAll\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"approved\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BatchMetadataUpdate\",\"inputs\":[{\"name\":\"_fromTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_toTokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerkleRootUpdated\",\"inputs\":[{\"name\":\"newMerkleRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MetadataUpdate\",\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTMinted\",\"inputs\":[{\"name\":\"minter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawal\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC721IncorrectOwner\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InsufficientApproval\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOperator\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC721NonexistentToken\",\"inputs\":[{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// OgwhitelistABI is the input ABI used to generate the binding from.
// Deprecated: Use OgwhitelistMetaData.ABI instead.
var OgwhitelistABI = OgwhitelistMetaData.ABI

// Ogwhitelist is an auto generated Go binding around an Ethereum contract.
type Ogwhitelist struct {
	OgwhitelistCaller     // Read-only binding to the contract
	OgwhitelistTransactor // Write-only binding to the contract
	OgwhitelistFilterer   // Log filterer for contract events
}

// OgwhitelistCaller is an auto generated read-only Go binding around an Ethereum contract.
type OgwhitelistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OgwhitelistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OgwhitelistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OgwhitelistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OgwhitelistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OgwhitelistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OgwhitelistSession struct {
	Contract     *Ogwhitelist      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OgwhitelistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OgwhitelistCallerSession struct {
	Contract *OgwhitelistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OgwhitelistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OgwhitelistTransactorSession struct {
	Contract     *OgwhitelistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OgwhitelistRaw is an auto generated low-level Go binding around an Ethereum contract.
type OgwhitelistRaw struct {
	Contract *Ogwhitelist // Generic contract binding to access the raw methods on
}

// OgwhitelistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OgwhitelistCallerRaw struct {
	Contract *OgwhitelistCaller // Generic read-only contract binding to access the raw methods on
}

// OgwhitelistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OgwhitelistTransactorRaw struct {
	Contract *OgwhitelistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOgwhitelist creates a new instance of Ogwhitelist, bound to a specific deployed contract.
func NewOgwhitelist(address common.Address, backend bind.ContractBackend) (*Ogwhitelist, error) {
	contract, err := bindOgwhitelist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ogwhitelist{OgwhitelistCaller: OgwhitelistCaller{contract: contract}, OgwhitelistTransactor: OgwhitelistTransactor{contract: contract}, OgwhitelistFilterer: OgwhitelistFilterer{contract: contract}}, nil
}

// NewOgwhitelistCaller creates a new read-only instance of Ogwhitelist, bound to a specific deployed contract.
func NewOgwhitelistCaller(address common.Address, caller bind.ContractCaller) (*OgwhitelistCaller, error) {
	contract, err := bindOgwhitelist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OgwhitelistCaller{contract: contract}, nil
}

// NewOgwhitelistTransactor creates a new write-only instance of Ogwhitelist, bound to a specific deployed contract.
func NewOgwhitelistTransactor(address common.Address, transactor bind.ContractTransactor) (*OgwhitelistTransactor, error) {
	contract, err := bindOgwhitelist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OgwhitelistTransactor{contract: contract}, nil
}

// NewOgwhitelistFilterer creates a new log filterer instance of Ogwhitelist, bound to a specific deployed contract.
func NewOgwhitelistFilterer(address common.Address, filterer bind.ContractFilterer) (*OgwhitelistFilterer, error) {
	contract, err := bindOgwhitelist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OgwhitelistFilterer{contract: contract}, nil
}

// bindOgwhitelist binds a generic wrapper to an already deployed contract.
func bindOgwhitelist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OgwhitelistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ogwhitelist *OgwhitelistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ogwhitelist.Contract.OgwhitelistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ogwhitelist *OgwhitelistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.OgwhitelistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ogwhitelist *OgwhitelistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.OgwhitelistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ogwhitelist *OgwhitelistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ogwhitelist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ogwhitelist *OgwhitelistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ogwhitelist *OgwhitelistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.contract.Transact(opts, method, params...)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Ogwhitelist *OgwhitelistCaller) MAXSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "MAX_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Ogwhitelist *OgwhitelistSession) MAXSUPPLY() (*big.Int, error) {
	return _Ogwhitelist.Contract.MAXSUPPLY(&_Ogwhitelist.CallOpts)
}

// MAXSUPPLY is a free data retrieval call binding the contract method 0x32cb6b0c.
//
// Solidity: function MAX_SUPPLY() view returns(uint256)
func (_Ogwhitelist *OgwhitelistCallerSession) MAXSUPPLY() (*big.Int, error) {
	return _Ogwhitelist.Contract.MAXSUPPLY(&_Ogwhitelist.CallOpts)
}

// OGMINTPRICE is a free data retrieval call binding the contract method 0xf48c7673.
//
// Solidity: function OG_MINT_PRICE() view returns(uint256)
func (_Ogwhitelist *OgwhitelistCaller) OGMINTPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "OG_MINT_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OGMINTPRICE is a free data retrieval call binding the contract method 0xf48c7673.
//
// Solidity: function OG_MINT_PRICE() view returns(uint256)
func (_Ogwhitelist *OgwhitelistSession) OGMINTPRICE() (*big.Int, error) {
	return _Ogwhitelist.Contract.OGMINTPRICE(&_Ogwhitelist.CallOpts)
}

// OGMINTPRICE is a free data retrieval call binding the contract method 0xf48c7673.
//
// Solidity: function OG_MINT_PRICE() view returns(uint256)
func (_Ogwhitelist *OgwhitelistCallerSession) OGMINTPRICE() (*big.Int, error) {
	return _Ogwhitelist.Contract.OGMINTPRICE(&_Ogwhitelist.CallOpts)
}

// PUBLICMINTPRICE is a free data retrieval call binding the contract method 0x6bde2627.
//
// Solidity: function PUBLIC_MINT_PRICE() view returns(uint256)
func (_Ogwhitelist *OgwhitelistCaller) PUBLICMINTPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "PUBLIC_MINT_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBLICMINTPRICE is a free data retrieval call binding the contract method 0x6bde2627.
//
// Solidity: function PUBLIC_MINT_PRICE() view returns(uint256)
func (_Ogwhitelist *OgwhitelistSession) PUBLICMINTPRICE() (*big.Int, error) {
	return _Ogwhitelist.Contract.PUBLICMINTPRICE(&_Ogwhitelist.CallOpts)
}

// PUBLICMINTPRICE is a free data retrieval call binding the contract method 0x6bde2627.
//
// Solidity: function PUBLIC_MINT_PRICE() view returns(uint256)
func (_Ogwhitelist *OgwhitelistCallerSession) PUBLICMINTPRICE() (*big.Int, error) {
	return _Ogwhitelist.Contract.PUBLICMINTPRICE(&_Ogwhitelist.CallOpts)
}

// AddressMintedBalance is a free data retrieval call binding the contract method 0x18cae269.
//
// Solidity: function addressMintedBalance(address ) view returns(uint256)
func (_Ogwhitelist *OgwhitelistCaller) AddressMintedBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "addressMintedBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddressMintedBalance is a free data retrieval call binding the contract method 0x18cae269.
//
// Solidity: function addressMintedBalance(address ) view returns(uint256)
func (_Ogwhitelist *OgwhitelistSession) AddressMintedBalance(arg0 common.Address) (*big.Int, error) {
	return _Ogwhitelist.Contract.AddressMintedBalance(&_Ogwhitelist.CallOpts, arg0)
}

// AddressMintedBalance is a free data retrieval call binding the contract method 0x18cae269.
//
// Solidity: function addressMintedBalance(address ) view returns(uint256)
func (_Ogwhitelist *OgwhitelistCallerSession) AddressMintedBalance(arg0 common.Address) (*big.Int, error) {
	return _Ogwhitelist.Contract.AddressMintedBalance(&_Ogwhitelist.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ogwhitelist *OgwhitelistCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ogwhitelist *OgwhitelistSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ogwhitelist.Contract.BalanceOf(&_Ogwhitelist.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ogwhitelist *OgwhitelistCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ogwhitelist.Contract.BalanceOf(&_Ogwhitelist.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ogwhitelist *OgwhitelistCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ogwhitelist *OgwhitelistSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Ogwhitelist.Contract.GetApproved(&_Ogwhitelist.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ogwhitelist *OgwhitelistCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Ogwhitelist.Contract.GetApproved(&_Ogwhitelist.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ogwhitelist *OgwhitelistCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ogwhitelist *OgwhitelistSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Ogwhitelist.Contract.IsApprovedForAll(&_Ogwhitelist.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ogwhitelist *OgwhitelistCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Ogwhitelist.Contract.IsApprovedForAll(&_Ogwhitelist.CallOpts, owner, operator)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Ogwhitelist *OgwhitelistCaller) MerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "merkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Ogwhitelist *OgwhitelistSession) MerkleRoot() ([32]byte, error) {
	return _Ogwhitelist.Contract.MerkleRoot(&_Ogwhitelist.CallOpts)
}

// MerkleRoot is a free data retrieval call binding the contract method 0x2eb4a7ab.
//
// Solidity: function merkleRoot() view returns(bytes32)
func (_Ogwhitelist *OgwhitelistCallerSession) MerkleRoot() ([32]byte, error) {
	return _Ogwhitelist.Contract.MerkleRoot(&_Ogwhitelist.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ogwhitelist *OgwhitelistCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ogwhitelist *OgwhitelistSession) Name() (string, error) {
	return _Ogwhitelist.Contract.Name(&_Ogwhitelist.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ogwhitelist *OgwhitelistCallerSession) Name() (string, error) {
	return _Ogwhitelist.Contract.Name(&_Ogwhitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ogwhitelist *OgwhitelistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ogwhitelist *OgwhitelistSession) Owner() (common.Address, error) {
	return _Ogwhitelist.Contract.Owner(&_Ogwhitelist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ogwhitelist *OgwhitelistCallerSession) Owner() (common.Address, error) {
	return _Ogwhitelist.Contract.Owner(&_Ogwhitelist.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ogwhitelist *OgwhitelistCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ogwhitelist *OgwhitelistSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Ogwhitelist.Contract.OwnerOf(&_Ogwhitelist.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Ogwhitelist *OgwhitelistCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Ogwhitelist.Contract.OwnerOf(&_Ogwhitelist.CallOpts, tokenId)
}

// RemainingSupply is a free data retrieval call binding the contract method 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (_Ogwhitelist *OgwhitelistCaller) RemainingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "remainingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingSupply is a free data retrieval call binding the contract method 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (_Ogwhitelist *OgwhitelistSession) RemainingSupply() (*big.Int, error) {
	return _Ogwhitelist.Contract.RemainingSupply(&_Ogwhitelist.CallOpts)
}

// RemainingSupply is a free data retrieval call binding the contract method 0xda0239a6.
//
// Solidity: function remainingSupply() view returns(uint256)
func (_Ogwhitelist *OgwhitelistCallerSession) RemainingSupply() (*big.Int, error) {
	return _Ogwhitelist.Contract.RemainingSupply(&_Ogwhitelist.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ogwhitelist *OgwhitelistCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ogwhitelist *OgwhitelistSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ogwhitelist.Contract.SupportsInterface(&_Ogwhitelist.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ogwhitelist *OgwhitelistCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ogwhitelist.Contract.SupportsInterface(&_Ogwhitelist.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ogwhitelist *OgwhitelistCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ogwhitelist *OgwhitelistSession) Symbol() (string, error) {
	return _Ogwhitelist.Contract.Symbol(&_Ogwhitelist.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ogwhitelist *OgwhitelistCallerSession) Symbol() (string, error) {
	return _Ogwhitelist.Contract.Symbol(&_Ogwhitelist.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ogwhitelist *OgwhitelistCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ogwhitelist *OgwhitelistSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Ogwhitelist.Contract.TokenURI(&_Ogwhitelist.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ogwhitelist *OgwhitelistCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Ogwhitelist.Contract.TokenURI(&_Ogwhitelist.CallOpts, tokenId)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Ogwhitelist *OgwhitelistCaller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Ogwhitelist *OgwhitelistSession) TotalMinted() (*big.Int, error) {
	return _Ogwhitelist.Contract.TotalMinted(&_Ogwhitelist.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_Ogwhitelist *OgwhitelistCallerSession) TotalMinted() (*big.Int, error) {
	return _Ogwhitelist.Contract.TotalMinted(&_Ogwhitelist.CallOpts)
}

// VerifyWhitelistStatus is a free data retrieval call binding the contract method 0xd236dc78.
//
// Solidity: function verifyWhitelistStatus(bytes32[] _merkleProof, address _address) view returns(bool)
func (_Ogwhitelist *OgwhitelistCaller) VerifyWhitelistStatus(opts *bind.CallOpts, _merkleProof [][32]byte, _address common.Address) (bool, error) {
	var out []interface{}
	err := _Ogwhitelist.contract.Call(opts, &out, "verifyWhitelistStatus", _merkleProof, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyWhitelistStatus is a free data retrieval call binding the contract method 0xd236dc78.
//
// Solidity: function verifyWhitelistStatus(bytes32[] _merkleProof, address _address) view returns(bool)
func (_Ogwhitelist *OgwhitelistSession) VerifyWhitelistStatus(_merkleProof [][32]byte, _address common.Address) (bool, error) {
	return _Ogwhitelist.Contract.VerifyWhitelistStatus(&_Ogwhitelist.CallOpts, _merkleProof, _address)
}

// VerifyWhitelistStatus is a free data retrieval call binding the contract method 0xd236dc78.
//
// Solidity: function verifyWhitelistStatus(bytes32[] _merkleProof, address _address) view returns(bool)
func (_Ogwhitelist *OgwhitelistCallerSession) VerifyWhitelistStatus(_merkleProof [][32]byte, _address common.Address) (bool, error) {
	return _Ogwhitelist.Contract.VerifyWhitelistStatus(&_Ogwhitelist.CallOpts, _merkleProof, _address)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ogwhitelist *OgwhitelistTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ogwhitelist.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ogwhitelist *OgwhitelistSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.Approve(&_Ogwhitelist.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ogwhitelist *OgwhitelistTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.Approve(&_Ogwhitelist.TransactOpts, to, tokenId)
}

// OgMint is a paid mutator transaction binding the contract method 0x918ed5d5.
//
// Solidity: function ogMint(bytes32[] _merkleProof) payable returns()
func (_Ogwhitelist *OgwhitelistTransactor) OgMint(opts *bind.TransactOpts, _merkleProof [][32]byte) (*types.Transaction, error) {
	return _Ogwhitelist.contract.Transact(opts, "ogMint", _merkleProof)
}

// OgMint is a paid mutator transaction binding the contract method 0x918ed5d5.
//
// Solidity: function ogMint(bytes32[] _merkleProof) payable returns()
func (_Ogwhitelist *OgwhitelistSession) OgMint(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.OgMint(&_Ogwhitelist.TransactOpts, _merkleProof)
}

// OgMint is a paid mutator transaction binding the contract method 0x918ed5d5.
//
// Solidity: function ogMint(bytes32[] _merkleProof) payable returns()
func (_Ogwhitelist *OgwhitelistTransactorSession) OgMint(_merkleProof [][32]byte) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.OgMint(&_Ogwhitelist.TransactOpts, _merkleProof)
}

// PublicMint is a paid mutator transaction binding the contract method 0x26092b83.
//
// Solidity: function publicMint() payable returns()
func (_Ogwhitelist *OgwhitelistTransactor) PublicMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ogwhitelist.contract.Transact(opts, "publicMint")
}

// PublicMint is a paid mutator transaction binding the contract method 0x26092b83.
//
// Solidity: function publicMint() payable returns()
func (_Ogwhitelist *OgwhitelistSession) PublicMint() (*types.Transaction, error) {
	return _Ogwhitelist.Contract.PublicMint(&_Ogwhitelist.TransactOpts)
}

// PublicMint is a paid mutator transaction binding the contract method 0x26092b83.
//
// Solidity: function publicMint() payable returns()
func (_Ogwhitelist *OgwhitelistTransactorSession) PublicMint() (*types.Transaction, error) {
	return _Ogwhitelist.Contract.PublicMint(&_Ogwhitelist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ogwhitelist *OgwhitelistTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ogwhitelist.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ogwhitelist *OgwhitelistSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ogwhitelist.Contract.RenounceOwnership(&_Ogwhitelist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ogwhitelist *OgwhitelistTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ogwhitelist.Contract.RenounceOwnership(&_Ogwhitelist.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ogwhitelist *OgwhitelistTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ogwhitelist.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ogwhitelist *OgwhitelistSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.SafeTransferFrom(&_Ogwhitelist.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ogwhitelist *OgwhitelistTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.SafeTransferFrom(&_Ogwhitelist.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Ogwhitelist *OgwhitelistTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Ogwhitelist.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Ogwhitelist *OgwhitelistSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.SafeTransferFrom0(&_Ogwhitelist.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Ogwhitelist *OgwhitelistTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.SafeTransferFrom0(&_Ogwhitelist.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ogwhitelist *OgwhitelistTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ogwhitelist.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ogwhitelist *OgwhitelistSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.SetApprovalForAll(&_Ogwhitelist.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ogwhitelist *OgwhitelistTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.SetApprovalForAll(&_Ogwhitelist.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ogwhitelist *OgwhitelistTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ogwhitelist.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ogwhitelist *OgwhitelistSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.TransferFrom(&_Ogwhitelist.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ogwhitelist *OgwhitelistTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.TransferFrom(&_Ogwhitelist.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ogwhitelist *OgwhitelistTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ogwhitelist.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ogwhitelist *OgwhitelistSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.TransferOwnership(&_Ogwhitelist.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ogwhitelist *OgwhitelistTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.TransferOwnership(&_Ogwhitelist.TransactOpts, newOwner)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 _newMerkleRoot) returns()
func (_Ogwhitelist *OgwhitelistTransactor) UpdateMerkleRoot(opts *bind.TransactOpts, _newMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Ogwhitelist.contract.Transact(opts, "updateMerkleRoot", _newMerkleRoot)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 _newMerkleRoot) returns()
func (_Ogwhitelist *OgwhitelistSession) UpdateMerkleRoot(_newMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.UpdateMerkleRoot(&_Ogwhitelist.TransactOpts, _newMerkleRoot)
}

// UpdateMerkleRoot is a paid mutator transaction binding the contract method 0x4783f0ef.
//
// Solidity: function updateMerkleRoot(bytes32 _newMerkleRoot) returns()
func (_Ogwhitelist *OgwhitelistTransactorSession) UpdateMerkleRoot(_newMerkleRoot [32]byte) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.UpdateMerkleRoot(&_Ogwhitelist.TransactOpts, _newMerkleRoot)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Ogwhitelist *OgwhitelistTransactor) Withdraw(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Ogwhitelist.contract.Transact(opts, "withdraw", _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Ogwhitelist *OgwhitelistSession) Withdraw(_to common.Address) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.Withdraw(&_Ogwhitelist.TransactOpts, _to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address _to) returns()
func (_Ogwhitelist *OgwhitelistTransactorSession) Withdraw(_to common.Address) (*types.Transaction, error) {
	return _Ogwhitelist.Contract.Withdraw(&_Ogwhitelist.TransactOpts, _to)
}

// OgwhitelistApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ogwhitelist contract.
type OgwhitelistApprovalIterator struct {
	Event *OgwhitelistApproval // Event containing the contract specifics and raw log

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
func (it *OgwhitelistApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OgwhitelistApproval)
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
		it.Event = new(OgwhitelistApproval)
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
func (it *OgwhitelistApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OgwhitelistApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OgwhitelistApproval represents a Approval event raised by the Ogwhitelist contract.
type OgwhitelistApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ogwhitelist *OgwhitelistFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*OgwhitelistApprovalIterator, error) {

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

	logs, sub, err := _Ogwhitelist.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &OgwhitelistApprovalIterator{contract: _Ogwhitelist.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ogwhitelist *OgwhitelistFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OgwhitelistApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Ogwhitelist.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OgwhitelistApproval)
				if err := _Ogwhitelist.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Ogwhitelist *OgwhitelistFilterer) ParseApproval(log types.Log) (*OgwhitelistApproval, error) {
	event := new(OgwhitelistApproval)
	if err := _Ogwhitelist.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OgwhitelistApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Ogwhitelist contract.
type OgwhitelistApprovalForAllIterator struct {
	Event *OgwhitelistApprovalForAll // Event containing the contract specifics and raw log

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
func (it *OgwhitelistApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OgwhitelistApprovalForAll)
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
		it.Event = new(OgwhitelistApprovalForAll)
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
func (it *OgwhitelistApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OgwhitelistApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OgwhitelistApprovalForAll represents a ApprovalForAll event raised by the Ogwhitelist contract.
type OgwhitelistApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ogwhitelist *OgwhitelistFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*OgwhitelistApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ogwhitelist.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &OgwhitelistApprovalForAllIterator{contract: _Ogwhitelist.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Ogwhitelist *OgwhitelistFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *OgwhitelistApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ogwhitelist.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OgwhitelistApprovalForAll)
				if err := _Ogwhitelist.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Ogwhitelist *OgwhitelistFilterer) ParseApprovalForAll(log types.Log) (*OgwhitelistApprovalForAll, error) {
	event := new(OgwhitelistApprovalForAll)
	if err := _Ogwhitelist.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OgwhitelistBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Ogwhitelist contract.
type OgwhitelistBatchMetadataUpdateIterator struct {
	Event *OgwhitelistBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *OgwhitelistBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OgwhitelistBatchMetadataUpdate)
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
		it.Event = new(OgwhitelistBatchMetadataUpdate)
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
func (it *OgwhitelistBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OgwhitelistBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OgwhitelistBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Ogwhitelist contract.
type OgwhitelistBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Ogwhitelist *OgwhitelistFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*OgwhitelistBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Ogwhitelist.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &OgwhitelistBatchMetadataUpdateIterator{contract: _Ogwhitelist.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Ogwhitelist *OgwhitelistFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *OgwhitelistBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Ogwhitelist.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OgwhitelistBatchMetadataUpdate)
				if err := _Ogwhitelist.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_Ogwhitelist *OgwhitelistFilterer) ParseBatchMetadataUpdate(log types.Log) (*OgwhitelistBatchMetadataUpdate, error) {
	event := new(OgwhitelistBatchMetadataUpdate)
	if err := _Ogwhitelist.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OgwhitelistMerkleRootUpdatedIterator is returned from FilterMerkleRootUpdated and is used to iterate over the raw logs and unpacked data for MerkleRootUpdated events raised by the Ogwhitelist contract.
type OgwhitelistMerkleRootUpdatedIterator struct {
	Event *OgwhitelistMerkleRootUpdated // Event containing the contract specifics and raw log

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
func (it *OgwhitelistMerkleRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OgwhitelistMerkleRootUpdated)
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
		it.Event = new(OgwhitelistMerkleRootUpdated)
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
func (it *OgwhitelistMerkleRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OgwhitelistMerkleRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OgwhitelistMerkleRootUpdated represents a MerkleRootUpdated event raised by the Ogwhitelist contract.
type OgwhitelistMerkleRootUpdated struct {
	NewMerkleRoot [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootUpdated is a free log retrieval operation binding the contract event 0x90004c04698bc3322499a575ed3752dd4abf33e0a7294c06a787a0fe01bea941.
//
// Solidity: event MerkleRootUpdated(bytes32 newMerkleRoot)
func (_Ogwhitelist *OgwhitelistFilterer) FilterMerkleRootUpdated(opts *bind.FilterOpts) (*OgwhitelistMerkleRootUpdatedIterator, error) {

	logs, sub, err := _Ogwhitelist.contract.FilterLogs(opts, "MerkleRootUpdated")
	if err != nil {
		return nil, err
	}
	return &OgwhitelistMerkleRootUpdatedIterator{contract: _Ogwhitelist.contract, event: "MerkleRootUpdated", logs: logs, sub: sub}, nil
}

// WatchMerkleRootUpdated is a free log subscription operation binding the contract event 0x90004c04698bc3322499a575ed3752dd4abf33e0a7294c06a787a0fe01bea941.
//
// Solidity: event MerkleRootUpdated(bytes32 newMerkleRoot)
func (_Ogwhitelist *OgwhitelistFilterer) WatchMerkleRootUpdated(opts *bind.WatchOpts, sink chan<- *OgwhitelistMerkleRootUpdated) (event.Subscription, error) {

	logs, sub, err := _Ogwhitelist.contract.WatchLogs(opts, "MerkleRootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OgwhitelistMerkleRootUpdated)
				if err := _Ogwhitelist.contract.UnpackLog(event, "MerkleRootUpdated", log); err != nil {
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

// ParseMerkleRootUpdated is a log parse operation binding the contract event 0x90004c04698bc3322499a575ed3752dd4abf33e0a7294c06a787a0fe01bea941.
//
// Solidity: event MerkleRootUpdated(bytes32 newMerkleRoot)
func (_Ogwhitelist *OgwhitelistFilterer) ParseMerkleRootUpdated(log types.Log) (*OgwhitelistMerkleRootUpdated, error) {
	event := new(OgwhitelistMerkleRootUpdated)
	if err := _Ogwhitelist.contract.UnpackLog(event, "MerkleRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OgwhitelistMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Ogwhitelist contract.
type OgwhitelistMetadataUpdateIterator struct {
	Event *OgwhitelistMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *OgwhitelistMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OgwhitelistMetadataUpdate)
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
		it.Event = new(OgwhitelistMetadataUpdate)
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
func (it *OgwhitelistMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OgwhitelistMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OgwhitelistMetadataUpdate represents a MetadataUpdate event raised by the Ogwhitelist contract.
type OgwhitelistMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Ogwhitelist *OgwhitelistFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*OgwhitelistMetadataUpdateIterator, error) {

	logs, sub, err := _Ogwhitelist.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &OgwhitelistMetadataUpdateIterator{contract: _Ogwhitelist.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Ogwhitelist *OgwhitelistFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *OgwhitelistMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Ogwhitelist.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OgwhitelistMetadataUpdate)
				if err := _Ogwhitelist.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_Ogwhitelist *OgwhitelistFilterer) ParseMetadataUpdate(log types.Log) (*OgwhitelistMetadataUpdate, error) {
	event := new(OgwhitelistMetadataUpdate)
	if err := _Ogwhitelist.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OgwhitelistNFTMintedIterator is returned from FilterNFTMinted and is used to iterate over the raw logs and unpacked data for NFTMinted events raised by the Ogwhitelist contract.
type OgwhitelistNFTMintedIterator struct {
	Event *OgwhitelistNFTMinted // Event containing the contract specifics and raw log

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
func (it *OgwhitelistNFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OgwhitelistNFTMinted)
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
		it.Event = new(OgwhitelistNFTMinted)
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
func (it *OgwhitelistNFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OgwhitelistNFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OgwhitelistNFTMinted represents a NFTMinted event raised by the Ogwhitelist contract.
type OgwhitelistNFTMinted struct {
	Minter  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNFTMinted is a free log retrieval operation binding the contract event 0x4cc0a9c4a99ddc700de1af2c9f916a7cbfdb71f14801ccff94061ad1ef8a8040.
//
// Solidity: event NFTMinted(address minter, uint256 tokenId)
func (_Ogwhitelist *OgwhitelistFilterer) FilterNFTMinted(opts *bind.FilterOpts) (*OgwhitelistNFTMintedIterator, error) {

	logs, sub, err := _Ogwhitelist.contract.FilterLogs(opts, "NFTMinted")
	if err != nil {
		return nil, err
	}
	return &OgwhitelistNFTMintedIterator{contract: _Ogwhitelist.contract, event: "NFTMinted", logs: logs, sub: sub}, nil
}

// WatchNFTMinted is a free log subscription operation binding the contract event 0x4cc0a9c4a99ddc700de1af2c9f916a7cbfdb71f14801ccff94061ad1ef8a8040.
//
// Solidity: event NFTMinted(address minter, uint256 tokenId)
func (_Ogwhitelist *OgwhitelistFilterer) WatchNFTMinted(opts *bind.WatchOpts, sink chan<- *OgwhitelistNFTMinted) (event.Subscription, error) {

	logs, sub, err := _Ogwhitelist.contract.WatchLogs(opts, "NFTMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OgwhitelistNFTMinted)
				if err := _Ogwhitelist.contract.UnpackLog(event, "NFTMinted", log); err != nil {
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

// ParseNFTMinted is a log parse operation binding the contract event 0x4cc0a9c4a99ddc700de1af2c9f916a7cbfdb71f14801ccff94061ad1ef8a8040.
//
// Solidity: event NFTMinted(address minter, uint256 tokenId)
func (_Ogwhitelist *OgwhitelistFilterer) ParseNFTMinted(log types.Log) (*OgwhitelistNFTMinted, error) {
	event := new(OgwhitelistNFTMinted)
	if err := _Ogwhitelist.contract.UnpackLog(event, "NFTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OgwhitelistOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ogwhitelist contract.
type OgwhitelistOwnershipTransferredIterator struct {
	Event *OgwhitelistOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OgwhitelistOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OgwhitelistOwnershipTransferred)
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
		it.Event = new(OgwhitelistOwnershipTransferred)
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
func (it *OgwhitelistOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OgwhitelistOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OgwhitelistOwnershipTransferred represents a OwnershipTransferred event raised by the Ogwhitelist contract.
type OgwhitelistOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ogwhitelist *OgwhitelistFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OgwhitelistOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ogwhitelist.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OgwhitelistOwnershipTransferredIterator{contract: _Ogwhitelist.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ogwhitelist *OgwhitelistFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OgwhitelistOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ogwhitelist.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OgwhitelistOwnershipTransferred)
				if err := _Ogwhitelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ogwhitelist *OgwhitelistFilterer) ParseOwnershipTransferred(log types.Log) (*OgwhitelistOwnershipTransferred, error) {
	event := new(OgwhitelistOwnershipTransferred)
	if err := _Ogwhitelist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OgwhitelistTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ogwhitelist contract.
type OgwhitelistTransferIterator struct {
	Event *OgwhitelistTransfer // Event containing the contract specifics and raw log

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
func (it *OgwhitelistTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OgwhitelistTransfer)
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
		it.Event = new(OgwhitelistTransfer)
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
func (it *OgwhitelistTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OgwhitelistTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OgwhitelistTransfer represents a Transfer event raised by the Ogwhitelist contract.
type OgwhitelistTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ogwhitelist *OgwhitelistFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*OgwhitelistTransferIterator, error) {

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

	logs, sub, err := _Ogwhitelist.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &OgwhitelistTransferIterator{contract: _Ogwhitelist.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ogwhitelist *OgwhitelistFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OgwhitelistTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Ogwhitelist.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OgwhitelistTransfer)
				if err := _Ogwhitelist.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Ogwhitelist *OgwhitelistFilterer) ParseTransfer(log types.Log) (*OgwhitelistTransfer, error) {
	event := new(OgwhitelistTransfer)
	if err := _Ogwhitelist.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OgwhitelistWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Ogwhitelist contract.
type OgwhitelistWithdrawalIterator struct {
	Event *OgwhitelistWithdrawal // Event containing the contract specifics and raw log

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
func (it *OgwhitelistWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OgwhitelistWithdrawal)
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
		it.Event = new(OgwhitelistWithdrawal)
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
func (it *OgwhitelistWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OgwhitelistWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OgwhitelistWithdrawal represents a Withdrawal event raised by the Ogwhitelist contract.
type OgwhitelistWithdrawal struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address to, uint256 amount)
func (_Ogwhitelist *OgwhitelistFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*OgwhitelistWithdrawalIterator, error) {

	logs, sub, err := _Ogwhitelist.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &OgwhitelistWithdrawalIterator{contract: _Ogwhitelist.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address to, uint256 amount)
func (_Ogwhitelist *OgwhitelistFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *OgwhitelistWithdrawal) (event.Subscription, error) {

	logs, sub, err := _Ogwhitelist.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OgwhitelistWithdrawal)
				if err := _Ogwhitelist.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address to, uint256 amount)
func (_Ogwhitelist *OgwhitelistFilterer) ParseWithdrawal(log types.Log) (*OgwhitelistWithdrawal, error) {
	event := new(OgwhitelistWithdrawal)
	if err := _Ogwhitelist.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
