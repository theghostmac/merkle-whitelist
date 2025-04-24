// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";


contract OGWhitelist is ERC721URIStorage, Ownable, ReentrancyGuard {
    uint256 private _tokenIdCounter = 0;
    bytes32 public merkleRoot;
    uint256 public constant MAX_SUPPLY = 1000;
    uint256 public constant OG_MINT_PRICE = 0.0000000000001 ether;
    uint256 public constant PUBLIC_MINT_PRICE = 0.0000000000005 ether;

    mapping(address => uint256) public addressMintedBalance;

    // Events for transparency
    event MerkleRootUpdated(bytes32 newMerkleRoot);
    event NFTMinted(address minter, uint256 tokenId);
    event Withdrawal(address to, uint256 amount);

    constructor(bytes32 _merkleRoot, address initialOwner) ERC721("OGWhitelistNFT", "OGNFT") Ownable(initialOwner) {
        merkleRoot = _merkleRoot;
    }

    function updateMerkleRoot(bytes32 _newMerkleRoot) public onlyOwner {
        merkleRoot = _newMerkleRoot;
        emit MerkleRootUpdated(_newMerkleRoot);
    }

    function publicMint() public payable nonReentrant {
        require(_tokenIdCounter < MAX_SUPPLY, "Max supply reached");
        require(msg.value >= PUBLIC_MINT_PRICE, "Insufficient funds for public mint");
        require(addressMintedBalance[msg.sender] < 3, "Address has already minted the maximum allowed");

        _mintNFT();

        addressMintedBalance[msg.sender]++;
    }

    function ogMint(bytes32[] memory _merkleProof) public payable nonReentrant {
        require(_tokenIdCounter < MAX_SUPPLY, "Max supply reached");
        require(msg.value >= OG_MINT_PRICE, "Insufficient funds for OG mint");
        require(addressMintedBalance[msg.sender] < 1, "Address has already minted");

        bytes32 leaf = keccak256(abi.encodePacked(msg.sender));
        require(MerkleProof.verify(_merkleProof, merkleRoot, leaf), "Invalid proof");

        _mintNFT();

        addressMintedBalance[msg.sender]++;
    }

    function verifyWhitelistStatus(bytes32[] calldata _merkleProof, address _address) public view returns (bool) {
        bytes32 leaf = keccak256(abi.encodePacked(_address));
        return MerkleProof.verify(_merkleProof, merkleRoot, leaf);
    }

    function totalMinted() public view returns (uint256) {
        return _tokenIdCounter;
    }

    function remainingSupply() public view returns (uint256) {
        return MAX_SUPPLY - _tokenIdCounter;
    }

    function withdraw(address payable _to) public onlyOwner nonReentrant {
        require(_to != address(0), "Invalid address");
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");

        (bool success,) = _to.call{value: balance}("");
        require(success, "Transfer failed");

        emit Withdrawal(_to, balance);
    }

    function _mintNFT() private {
        require(_tokenIdCounter < MAX_SUPPLY, "Max NFT supply reached");
        _safeMint(msg.sender, _tokenIdCounter);
        emit NFTMinted(msg.sender, _tokenIdCounter);
        _tokenIdCounter++;
    }
}