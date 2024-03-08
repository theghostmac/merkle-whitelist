const { ethers } = require("ethers");

async function main() {
    // The address of the deployed OGWhitelist contract
    const contractAddress = "0x5FbDB2315678afecb367f032d93F642f64180aa3";

    // The address we're testing with
    const testAddress = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266";

    // The Merkle proof for the address
    const proof = [
        "417ebaa8de17c9a73acf7cf3f6c998ca2dd1e993276cc9dd37dfb7fb0cb197a4",
        "417ebaa8de17c9a73acf7cf3f6c998ca2dd1e993276cc9dd37dfb7fb0cb197a4",
        "417ebaa8de17c9a73acf7cf3f6c998ca2dd1e993276cc9dd37dfb7fb0cb197a4",
        "417ebaa8de17c9a73acf7cf3f6c998ca2dd1e993276cc9dd37dfb7fb0cb197a4"
    ];

    // Quantity to mint
    const quantity = 1; // Adjust based on what your contract allows

    // The cost to mint (ensure this matches the contract requirement for ogMint)
    const value = ethers.parseEther("0.0001");

    // Setup provider and signer
    const [signer] = await ethers.getSigners();

    // Connect to your deployed contract
    const contract = await ethers.getContractAt("OGWhitelist", contractAddress, signer);

    // Call the ogMint function with the Merkle proof, quantity, and value
    const tx = await contract.ogMint(proof, quantity, { value });

    // Wait for the transaction to be mined
    await tx.wait();

    console.log(`Minted ${quantity} NFT(s) for address ${testAddress}`);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
