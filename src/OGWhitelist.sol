// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract OGWhitelist {
    bytes public merkleRoot;

    // === Event for logging the whitelisted address access ===
    event AccessGranted(address addr);

    constructor(bytes32 _merkleRoot){
        merkleRoot = _merkleRoot;
    }

    // === Function to update the Merkle root, only callable by owner ===
    function updateMerkleRoot(bytes32 _merkleRoot) external onlyOwner {
        merkleRoot = _merkleRoot;
    }

    // === Function to verify the caller's address against the Merkle tree ===
    function verifyAddress(bytes32[] calldata _merkleProof) public view returns (bool) {
        bytes32 leaf = keccak256(abi.encodePacked(msg.sender));
        return MerkleProof.verify(_merkleProof, merkleRoot, leaf);
    }

    // === Function for only whitelisted addresses to call ===
    function whitelistedFunction(bytes32[] calldata _merkleProof) external {
        require(verifyAddress(_merkleProof), "OGWhitelist: Caller's address is not whitelisted");

        // Emit an event for successful access.
        emit AccessGranted(msg.sender);

        // Function body goes here.
    }
}
