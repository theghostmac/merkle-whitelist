// merkleAlgo.js
const { MerkleTree } = require('merkletreejs');
const { keccak256 } = require('js-sha3');
const fs = require('fs/promises'); // Use fs/promises for async file read

// Function to asynchronously load addresses from JSON file
async function loadAddresses() {
    try {
        const data = await fs.readFile("./addresses.json", { encoding: 'utf8' });
        return JSON.parse(data).addresses;
    } catch (error) {
        console.error('Error loading addresses:', error);
        process.exit(1);
    }
}

// Function to create a Merkle Tree and find the root hash
const createMerkleTree = (addressList) => {
    const leaves = addressList.map(addr => keccak256(addr));
    return new MerkleTree(leaves, keccak256, { sortPairs: true });
};

// Function to find a hex proof for a specific address
const findHexProofForAddress = async (address) => {
    const addresses = await loadAddresses(); // Load addresses asynchronously

    const tree = createMerkleTree(addresses);
    const hashedAddress = keccak256(address);
    const proof = tree.getHexProof(hashedAddress);

    const isValid = tree.verify(proof, hashedAddress, tree.getRoot());

    if (isValid) {
        console.log("Address is in the list", { proof: proof, isValid: isValid });
        return { tree: "Valid", proof: proof, isValid: isValid };
    } else {
        console.log("Address is not in the list");
        return { tree: "Invalid", proof: [], isValid: false };
    }
};
