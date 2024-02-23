// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import { ERC721 } from "../lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol";
import { Ownable } from "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";
import { MerkleProof } from "../lib/openzeppelin-contracts/contracts/utils/cryptography/MerkleProof.sol";

contract OGWhitelist is ERC721, Ownable {
    uint256 private _tokenIdCounter = 0;
    bytes public merkleRoot;
    uint256 public constant MAX_SUPPLY = 4000;
    uint256 public constant OG_MINT_PRICE = 0.001 ether;
    uint256 public constant PUBLIC_MINT_PRICE = 0.005 ether;

    mapping(address => uint) public addressMintedBalance;

   constructor(bytes32 _merkleRoot) ERC721("OGWhitelistNFT", "OGNFT") {
       merkleRoot = _merkleRoot;
   }

    function updateMerkleRoot(bytes32 _newMerkleRoot) public onlyOwner {
        merkleRoot = _newMerkleRoot;
    }

    function publicMint() public payable {
        require(_tokenIdCounter < MAX_SUPPLY, "Max supply reached");
        require(msg.value >= PUBLIC_MINT_PRICE, "Insufficient funds for public mint");
        require(addressMintedBalance[msg.sender] < 3, "Address has already minted");

        _mintNFT();

        addressMintedBalance[msg.sender]++;
    }

    function ogMint(bytes32[] calldata _merkleProof) public payable {
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

    function withdraw(address payable _to) public onlyOwner {
        require(_to != address(0), "Invalid address");
        require(address(this).balance > 0, "No funds to withdraw");

        (bool success, ) = _to.call{value: address(this).balance}("");
        require(success, "Transfer failed");
    }

    function _mintNFT() private {
        require(_tokenIdCounter < MAX_SUPPLY, "Max NFT supply reached");
        _safeMint(msg.sender, _tokenIdCounter);
        _tokenIdCounter++;
    }
}
