// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ERC721} from "../lib/openzeppelin-contracts/contracts/token/ERC721/ERC721.sol";
import {Ownable} from "../lib/openzeppelin-contracts/contracts/access/Ownable.sol";
import {MerkleProof} from "../lib/openzeppelin-contracts/contracts/utils/cryptography/MerkleProof.sol";

contract TieredOGWhitelist is ERC721, Ownable {
    uint256 private _tokenIdCounter = 0;
    bytes32 public merkleRoot;

    // **OG Supporter Tiers and Benefits**
    enum OGTier {
        None,
        Bronze,
        Silver,
        Gold
    }

    mapping(address => OGTier) public ogTiers;
    mapping(OGTier => uint256) public tierMintPrice;
    mapping(OGTier => uint256) public tierMintLimit;

    // **Anti-Sybil Mechanism**
    mapping(address => bool) public minted; // Tracks addresses that have already minted

    constructor(bytes32 _merkleRoot) ERC721("OGWhitelistNFT", "OGNFT") {
        merkleRoot = _merkleRoot;

        // **Set initial tier prices and limits**
        tierMintPrice[OGTier.Bronze] = 0.0005 ether;
        tierMintLimit[OGTier.Bronze] = 2;
        tierMintPrice[OGTier.Silver] = 0.0002 ether;
        tierMintLimit[OGTier.Silver] = 5;
        tierMintPrice[OGTier.Gold] = 0.0001 ether;
        tierMintLimit[OGTier.Gold] = 10;
    }

    function updateMerkleRoot(bytes32 _newMerkleRoot) public onlyOwner {
        merkleRoot = _newMerkleRoot;
    }

    function setOGTier(address[] calldata _addresses, OGTier _tier) public onlyOwner {
        for (uint256 i = 0; i < _addresses.length; i++) {
            ogTiers[_addresses[i]] = _tier;
        }
    }

    function publicMint() public payable {
        require(msg.value >= PUBLIC_MINT_PRICE, "Insufficient funds for public mint");
        _mintNFT();
    }

    function ogMint(bytes32[] calldata _merkleProof) public payable {
        require(msg.value >= tierMintPrice[ogTiers[msg.sender]], "Insufficient funds for OG mint");
        require(!minted[msg.sender], "Address already minted");

        bytes32 leaf = keccak256(abi.encodePacked(msg.sender));
        require(MerkleProof.verify(_merkleProof, merkleRoot, leaf), "Invalid proof");

        // Ensure mint limit for the tier is not exceeded.
        require(_balanceOf[msg.sender] < tierMintLimit[ogTiers[msg.sender]], "Mint limit for tier reached");

        _mintNFT();
        minted[msg.sender] = true;
    }

    function _mintNFT() private {
        _safeMint(msg.sender, _tokenIdCounter);
        _tokenIdCounter++; // increment the counter after minting.
    }
}
