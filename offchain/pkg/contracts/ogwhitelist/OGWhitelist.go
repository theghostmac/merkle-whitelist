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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newMerkleRoot\",\"type\":\"bytes32\"}],\"name\":\"MerkleRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"NFTMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OG_MINT_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLIC_MINT_PRICE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressMintedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"merkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"}],\"name\":\"ogMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_newMerkleRoot\",\"type\":\"bytes32\"}],\"name\":\"updateMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_merkleProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"verifyWhitelistStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040525f60095534801562000014575f80fd5b5060405162001f3b38038062001f3b833981016040819052620000379162000139565b806040518060400160405280600e81526020016d13d1d5da1a5d195b1a5cdd13919560921b8152506040518060400160405280600581526020016413d1d3919560da1b815250815f90816200008d919062000213565b5060016200009c828262000213565b5050506001600160a01b038116620000cd57604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b620000d881620000e8565b50506001600855600a55620002df565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f80604083850312156200014b575f80fd5b825160208401519092506001600160a01b03811681146200016a575f80fd5b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200019e57607f821691505b602082108103620001bd57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f8211156200020e57805f5260205f20601f840160051c81016020851015620001ea5750805b601f840160051c820191505b818110156200020b575f8155600101620001f6565b50505b505050565b81516001600160401b038111156200022f576200022f62000175565b620002478162000240845462000189565b84620001c3565b602080601f8311600181146200027d575f8415620002655750858301515b5f19600386901b1c1916600185901b178555620002d7565b5f85815260208120601f198616915b82811015620002ad578886015182559484019460019091019084016200028c565b5085821015620002cb57878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b611c4e80620002ed5f395ff3fe6080604052600436106101af575f3560e01c806370a08231116100e7578063b88d4fde11610087578063da0239a611610062578063da0239a614610483578063e985e9c514610497578063f2fde38b146104b6578063f48c7673146104d5575f80fd5b8063b88d4fde14610426578063c87b56dd14610445578063d236dc7814610464575f80fd5b8063918ed5d5116100c2578063918ed5d5146103cc57806395d89b41146103df578063a22cb465146103f3578063a2309ff814610412575f80fd5b806370a082311461037c578063715018a61461039b5780638da5cb5b146103af575f80fd5b80632eb4a7ab116101525780634783f0ef1161012d5780634783f0ef1461030957806351cff8d9146103285780636352211e146103475780636bde262714610366575f80fd5b80632eb4a7ab146102c057806332cb6b0c146102d557806342842e0e146102ea575f80fd5b8063095ea7b31161018d578063095ea7b31461023f57806318cae2691461026057806323b872dd1461029957806326092b83146102b8575f80fd5b806301ffc9a7146101b357806306fdde03146101e7578063081812fc14610208575b5f80fd5b3480156101be575f80fd5b506101d26101cd366004611757565b6104eb565b60405190151581526020015b60405180910390f35b3480156101f2575f80fd5b506101fb610515565b6040516101de91906117bf565b348015610213575f80fd5b506102276102223660046117d1565b6105a4565b6040516001600160a01b0390911681526020016101de565b34801561024a575f80fd5b5061025e6102593660046117fc565b6105cb565b005b34801561026b575f80fd5b5061028b61027a366004611826565b600b6020525f908152604090205481565b6040519081526020016101de565b3480156102a4575f80fd5b5061025e6102b3366004611841565b6105da565b61025e610668565b3480156102cb575f80fd5b5061028b600a5481565b3480156102e0575f80fd5b5061028b6103e881565b3480156102f5575f80fd5b5061025e610304366004611841565b6107be565b348015610314575f80fd5b5061025e6103233660046117d1565b6107dd565b348015610333575f80fd5b5061025e610342366004611826565b610820565b348015610352575f80fd5b506102276103613660046117d1565b61099f565b348015610371575f80fd5b5061028b6207a12081565b348015610387575f80fd5b5061028b610396366004611826565b6109a9565b3480156103a6575f80fd5b5061025e6109ee565b3480156103ba575f80fd5b506007546001600160a01b0316610227565b61025e6103da3660046118c4565b6109ff565b3480156103ea575f80fd5b506101fb610bb3565b3480156103fe575f80fd5b5061025e61040d366004611965565b610bc2565b34801561041d575f80fd5b5060095461028b565b348015610431575f80fd5b5061025e6104403660046119a0565b610bcd565b348015610450575f80fd5b506101fb61045f3660046117d1565b610be5565b34801561046f575f80fd5b506101d261047e366004611a5d565b610cf0565b34801561048e575f80fd5b5061028b610d74565b3480156104a2575f80fd5b506101d26104b1366004611adc565b610d8a565b3480156104c1575f80fd5b5061025e6104d0366004611826565b610db7565b3480156104e0575f80fd5b5061028b620186a081565b5f6001600160e01b03198216632483248360e11b148061050f575061050f82610df1565b92915050565b60605f805461052390611b08565b80601f016020809104026020016040519081016040528092919081815260200182805461054f90611b08565b801561059a5780601f106105715761010080835404028352916020019161059a565b820191905f5260205f20905b81548152906001019060200180831161057d57829003601f168201915b5050505050905090565b5f6105ae82610e40565b505f828152600460205260409020546001600160a01b031661050f565b6105d6828233610e78565b5050565b6001600160a01b03821661060857604051633250574960e11b81525f60048201526024015b60405180910390fd5b5f610614838333610e85565b9050836001600160a01b0316816001600160a01b031614610662576040516364283d7b60e01b81526001600160a01b03808616600483015260248201849052821660448201526064016105ff565b50505050565b610670610f77565b6103e8600954106106b85760405162461bcd60e51b815260206004820152601260248201527113585e081cdd5c1c1b1e481c995858da195960721b60448201526064016105ff565b6207a1203410156107165760405162461bcd60e51b815260206004820152602260248201527f496e73756666696369656e742066756e647320666f72207075626c6963206d696044820152611b9d60f21b60648201526084016105ff565b335f908152600b602052604090205460031161078b5760405162461bcd60e51b815260206004820152602e60248201527f416464726573732068617320616c7265616479206d696e74656420746865206d60448201526d185e1a5b5d5b48185b1b1bddd95960921b60648201526084016105ff565b610793610fa1565b335f908152600b602052604081208054916107ad83611b54565b91905055506107bc6001600855565b565b6107d883838360405180602001604052805f815250610bcd565b505050565b6107e561104c565b600a8190556040518181527f90004c04698bc3322499a575ed3752dd4abf33e0a7294c06a787a0fe01bea9419060200160405180910390a150565b61082861104c565b610830610f77565b6001600160a01b0381166108785760405162461bcd60e51b815260206004820152600f60248201526e496e76616c6964206164647265737360881b60448201526064016105ff565b47806108bd5760405162461bcd60e51b81526020600482015260146024820152734e6f2066756e647320746f20776974686472617760601b60448201526064016105ff565b5f826001600160a01b0316826040515f6040518083038185875af1925050503d805f8114610906576040519150601f19603f3d011682016040523d82523d5f602084013e61090b565b606091505b505090508061094e5760405162461bcd60e51b815260206004820152600f60248201526e151c985b9cd9995c8819985a5b1959608a1b60448201526064016105ff565b604080516001600160a01b0385168152602081018490527f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65910160405180910390a1505061099c6001600855565b50565b5f61050f82610e40565b5f6001600160a01b0382166109d3576040516322718ad960e21b81525f60048201526024016105ff565b506001600160a01b03165f9081526003602052604090205490565b6109f661104c565b6107bc5f611079565b610a07610f77565b6103e860095410610a4f5760405162461bcd60e51b815260206004820152601260248201527113585e081cdd5c1c1b1e481c995858da195960721b60448201526064016105ff565b620186a0341015610aa25760405162461bcd60e51b815260206004820152601e60248201527f496e73756666696369656e742066756e647320666f72204f47206d696e74000060448201526064016105ff565b335f908152600b6020526040902054600111610b005760405162461bcd60e51b815260206004820152601a60248201527f416464726573732068617320616c7265616479206d696e74656400000000000060448201526064016105ff565b6040516bffffffffffffffffffffffff193360601b1660208201525f90603401604051602081830303815290604052805190602001209050610b4582600a54836110ca565b610b815760405162461bcd60e51b815260206004820152600d60248201526c24b73b30b634b210383937b7b360991b60448201526064016105ff565b610b89610fa1565b335f908152600b60205260408120805491610ba383611b54565b91905055505061099c6001600855565b60606001805461052390611b08565b6105d63383836110df565b610bd88484846105da565b610662338585858561117d565b6060610bf082610e40565b505f8281526006602052604081208054610c0990611b08565b80601f0160208091040260200160405190810160405280929190818152602001828054610c3590611b08565b8015610c805780601f10610c5757610100808354040283529160200191610c80565b820191905f5260205f20905b815481529060010190602001808311610c6357829003601f168201915b505050505090505f610c9c60408051602081019091525f815290565b905080515f03610cad575092915050565b815115610cdf578082604051602001610cc7929190611b6c565b60405160208183030381529060405292505050919050565b610ce8846112a5565b949350505050565b6040516bffffffffffffffffffffffff19606083901b1660208201525f908190603401604051602081830303815290604052805190602001209050610d6b8585808060200260200160405190810160405280939291908181526020018383602002808284375f9201919091525050600a5491508490506110ca565b95945050505050565b5f6009546103e8610d859190611b9a565b905090565b6001600160a01b039182165f90815260056020908152604080832093909416825291909152205460ff1690565b610dbf61104c565b6001600160a01b038116610de857604051631e4fbdf760e01b81525f60048201526024016105ff565b61099c81611079565b5f6001600160e01b031982166380ac58cd60e01b1480610e2157506001600160e01b03198216635b5e139f60e01b145b8061050f57506301ffc9a760e01b6001600160e01b031983161461050f565b5f818152600260205260408120546001600160a01b03168061050f57604051637e27328960e01b8152600481018490526024016105ff565b6107d88383836001611316565b5f828152600260205260408120546001600160a01b0390811690831615610eb157610eb181848661141a565b6001600160a01b03811615610eeb57610ecc5f855f80611316565b6001600160a01b0381165f90815260036020526040902080545f190190555b6001600160a01b03851615610f19576001600160a01b0385165f908152600360205260409020805460010190555b5f8481526002602052604080822080546001600160a01b0319166001600160a01b0389811691821790925591518793918516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4949350505050565b600260085403610f9a57604051633ee5aeb560e01b815260040160405180910390fd5b6002600855565b6103e860095410610fed5760405162461bcd60e51b815260206004820152601660248201527513585e08139195081cdd5c1c1b1e481c995858da195960521b60448201526064016105ff565b610ff93360095461147e565b6009546040805133815260208101929092527f4cc0a9c4a99ddc700de1af2c9f916a7cbfdb71f14801ccff94061ad1ef8a8040910160405180910390a160098054905f61104583611b54565b9190505550565b6007546001600160a01b031633146107bc5760405163118cdaa760e01b81523360048201526024016105ff565b600780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f826110d68584611497565b14949350505050565b6001600160a01b03821661111157604051630b61174360e31b81526001600160a01b03831660048201526024016105ff565b6001600160a01b038381165f81815260056020908152604080832094871680845294825291829020805460ff191686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6001600160a01b0383163b1561129e57604051630a85bd0160e11b81526001600160a01b0384169063150b7a02906111bf908890889087908790600401611bad565b6020604051808303815f875af19250505080156111f9575060408051601f3d908101601f191682019092526111f691810190611be9565b60015b611260573d808015611226576040519150601f19603f3d011682016040523d82523d5f602084013e61122b565b606091505b5080515f0361125857604051633250574960e11b81526001600160a01b03851660048201526024016105ff565b805181602001fd5b6001600160e01b03198116630a85bd0160e11b1461129c57604051633250574960e11b81526001600160a01b03851660048201526024016105ff565b505b5050505050565b60606112b082610e40565b505f6112c660408051602081019091525f815290565b90505f8151116112e45760405180602001604052805f81525061130f565b806112ee846114d9565b6040516020016112ff929190611b6c565b6040516020818303038152906040525b9392505050565b808061132a57506001600160a01b03821615155b156113eb575f61133984610e40565b90506001600160a01b038316158015906113655750826001600160a01b0316816001600160a01b031614155b801561137857506113768184610d8a565b155b156113a15760405163a9fbf51f60e01b81526001600160a01b03841660048201526024016105ff565b81156113e95783856001600160a01b0316826001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b50505f90815260046020526040902080546001600160a01b0319166001600160a01b0392909216919091179055565b611425838383611569565b6107d8576001600160a01b03831661145357604051637e27328960e01b8152600481018290526024016105ff565b60405163177e802f60e01b81526001600160a01b0383166004820152602481018290526044016105ff565b6105d6828260405180602001604052805f8152506115ca565b5f81815b84518110156114d1576114c7828683815181106114ba576114ba611c04565b60200260200101516115e1565b915060010161149b565b509392505050565b60605f6114e58361160a565b60010190505f8167ffffffffffffffff8111156115045761150461187f565b6040519080825280601f01601f19166020018201604052801561152e576020820181803683370190505b5090508181016020015b5f19016f181899199a1a9b1b9c1cb0b131b232b360811b600a86061a8153600a850494508461153857509392505050565b5f6001600160a01b03831615801590610ce85750826001600160a01b0316846001600160a01b031614806115a257506115a28484610d8a565b80610ce85750505f908152600460205260409020546001600160a01b03908116911614919050565b6115d483836116e1565b6107d8335f85858561117d565b5f8183106115fb575f82815260208490526040902061130f565b505f9182526020526040902090565b5f8072184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b83106116485772184f03e93ff9f4daa797ed6e38ed64bf6a1f0160401b830492506040015b6d04ee2d6d415b85acef81000000008310611674576d04ee2d6d415b85acef8100000000830492506020015b662386f26fc10000831061169257662386f26fc10000830492506010015b6305f5e10083106116aa576305f5e100830492506008015b61271083106116be57612710830492506004015b606483106116d0576064830492506002015b600a831061050f5760010192915050565b6001600160a01b03821661170a57604051633250574960e11b81525f60048201526024016105ff565b5f61171683835f610e85565b90506001600160a01b038116156107d8576040516339e3563760e11b81525f60048201526024016105ff565b6001600160e01b03198116811461099c575f80fd5b5f60208284031215611767575f80fd5b813561130f81611742565b5f5b8381101561178c578181015183820152602001611774565b50505f910152565b5f81518084526117ab816020860160208601611772565b601f01601f19169290920160200192915050565b602081525f61130f6020830184611794565b5f602082840312156117e1575f80fd5b5035919050565b6001600160a01b038116811461099c575f80fd5b5f806040838503121561180d575f80fd5b8235611818816117e8565b946020939093013593505050565b5f60208284031215611836575f80fd5b813561130f816117e8565b5f805f60608486031215611853575f80fd5b833561185e816117e8565b9250602084013561186e816117e8565b929592945050506040919091013590565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156118bc576118bc61187f565b604052919050565b5f60208083850312156118d5575f80fd5b823567ffffffffffffffff808211156118ec575f80fd5b818501915085601f8301126118ff575f80fd5b8135818111156119115761191161187f565b8060051b9150611922848301611893565b818152918301840191848101908884111561193b575f80fd5b938501935b8385101561195957843582529385019390850190611940565b98975050505050505050565b5f8060408385031215611976575f80fd5b8235611981816117e8565b915060208301358015158114611995575f80fd5b809150509250929050565b5f805f80608085870312156119b3575f80fd5b84356119be816117e8565b93506020858101356119cf816117e8565b935060408601359250606086013567ffffffffffffffff808211156119f2575f80fd5b818801915088601f830112611a05575f80fd5b813581811115611a1757611a1761187f565b611a29601f8201601f19168501611893565b91508082528984828501011115611a3e575f80fd5b80848401858401375f8482840101525080935050505092959194509250565b5f805f60408486031215611a6f575f80fd5b833567ffffffffffffffff80821115611a86575f80fd5b818601915086601f830112611a99575f80fd5b813581811115611aa7575f80fd5b8760208260051b8501011115611abb575f80fd5b60209283019550935050840135611ad1816117e8565b809150509250925092565b5f8060408385031215611aed575f80fd5b8235611af8816117e8565b91506020830135611995816117e8565b600181811c90821680611b1c57607f821691505b602082108103611b3a57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b5f60018201611b6557611b65611b40565b5060010190565b5f8351611b7d818460208801611772565b835190830190611b91818360208801611772565b01949350505050565b8181038181111561050f5761050f611b40565b6001600160a01b03858116825284166020820152604081018390526080606082018190525f90611bdf90830184611794565b9695505050505050565b5f60208284031215611bf9575f80fd5b815161130f81611742565b634e487b7160e01b5f52603260045260245ffdfea26469706673582212201cfd10875ed0626cd4eca7e28d0cde5651d3f000fd24adedbbbe4ebde3df7a6364736f6c63430008180033",
}

// OgwhitelistABI is the input ABI used to generate the binding from.
// Deprecated: Use OgwhitelistMetaData.ABI instead.
var OgwhitelistABI = OgwhitelistMetaData.ABI

// OgwhitelistBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OgwhitelistMetaData.Bin instead.
var OgwhitelistBin = OgwhitelistMetaData.Bin

// DeployOgwhitelist deploys a new Ethereum contract, binding an instance of Ogwhitelist to it.
func DeployOgwhitelist(auth *bind.TransactOpts, backend bind.ContractBackend, _merkleRoot [32]byte, initialOwner common.Address) (common.Address, *types.Transaction, *Ogwhitelist, error) {
	parsed, err := OgwhitelistMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OgwhitelistBin), backend, _merkleRoot, initialOwner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ogwhitelist{OgwhitelistCaller: OgwhitelistCaller{contract: contract}, OgwhitelistTransactor: OgwhitelistTransactor{contract: contract}, OgwhitelistFilterer: OgwhitelistFilterer{contract: contract}}, nil
}

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
