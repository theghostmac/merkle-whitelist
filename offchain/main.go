package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"log"
	"merkle-whitelist/offchain/pkg/contracts/ogwhitelist"
	"os"
)

type AddressList struct {
	Addresses []string `json:"addresses"`
}

func main() {
	// To get the Merkle root:
	merkleRoot, err := getMerkleRoot("addresses.json")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Println("Merkle Root: ", hex.EncodeToString(merkleRoot))

	// Replace with the actual address and proof
	testAddress := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	proof, err := findHexProofForAnAddress("addresses.json", testAddress)
	if err != nil {
		log.Fatalf("Error getting proof for address: %v", err)
	}

	log.Println("Proof for the address: ", proof)

	// Establish a connection to the Ethereum client
	client, err := ethclient.Dial("http://localhost:8545") // Or an Ethereum RPC URL
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	} else {
		log.Print("Successfully dialed the ethereum client...")
	}

	// Load the contract
	contractAddress := common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")
	ogWhitelistInstance, err := ogwhitelist.NewOgwhitelist(contractAddress, client)
	if err != nil {
		log.Fatalf("Failed to instantiate a OGWhitelist contract: %v", err)
	}

	err = updateMerkleRoot(client, contractAddress, "f1d131581e485ec20d2172a7e3a2893fbe62a760a2aa36d8d602e5f8bd0c75e7")
	if err != nil {
		log.Fatalf("Failed to update the merkle root: %e", err)
	}

	price, err := ogWhitelistInstance.OGMINTPRICE(nil)
	if err != nil {
		log.Fatalf("Failed to get the OG Minting Price: %v", err)
	} else {
		log.Printf("OG Minting Price: %v", price)
	}

	// Simulate minting an NFT
	if err := MintNFT(ogWhitelistInstance, client, testAddress); err != nil {
		log.Println("Proof found: ", proof)
		log.Fatalf("Minting NFT failed: %v", err)
	} else {
		log.Printf("Successfully minted an NFT for %s, at price %v", testAddress, price)
	}
}

func MintNFT(contract *ogwhitelist.Ogwhitelist, client *ethclient.Client, testAddress string) error {
	// Extract the privateKey from environment variable
	privateKeyHex := os.Getenv("PRIVATE_KEY_HEX")
	if privateKeyHex == "" {
		return fmt.Errorf("PRIVATE_KEY_HEX environment variable not found")
	}

	// Convert the private key from a hex string to an *ecdsa.PrivateKey
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	// Fetch the network ID dynamically
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to fetch network ID: %v", err)
	}
	log.Println("---> Chain ID is: \n\n", chainID)

	// Prepare the authorized transaction options
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create keyed transactor: %v", err)
	}

	// Retrieve the OG mint price from the contract
	ogMintPrice, err := contract.OGMINTPRICE(nil)
	if err != nil {
		return fmt.Errorf("failed to fetch OG mint price: %v", err)
	}
	auth.Value = ogMintPrice // Set the transaction value

	// Dynamically calculate the proof for the provided testAddress
	proof, err := findHexProofForAnAddress("addresses.json", testAddress)
	if err != nil {
		return fmt.Errorf("failed to find proof for address %s: %v", testAddress, err)
	}

	log.Println("Proof from MintNFT func is: ", proof)

	// // Manually specify the correct proofs
	// proofs := []string{
	// 	"1eceaf9a464e556e95ce7556ff21729f1594ffefe67c7b814005588104635631",
	// 	"34b107911b4069ea53ad445e68d65ee199d64a549f75e7e0728725279bb2fbc4",
	// 	"db2ec52069cba65d62a560cad9088e7a709d09e7c14b33d740b428038e383e34",
	// 	"417ebaa8de17c9a73acf7cf3f6c998ca2dd1e993276cc9dd37dfb7fb0cb197a4",
	// }

	// Convert proof strings to [][32]byte format required by the smart contract
	var merkleProof [][32]byte
	for _, p := range proof {
		proofBytes, err := hex.DecodeString(p)
		if err != nil {
			return fmt.Errorf("invalid proof hex string: %v", err)
		}
		var proofElement [32]byte
		copy(proofElement[:], proofBytes[:32])
		merkleProof = append(merkleProof, proofElement)
	}

	// // Convert proof strings to [][32]byte format required by the smart contract
	// var merkleProof [][32]byte
	// for _, p := range proofs {
	// 	proofBytes, err := hex.DecodeString(p)
	// 	if err != nil {
	// 		return fmt.Errorf("invalid proof hex string: %v", err)
	// 	}
	// 	var proofElement [32]byte
	// 	copy(proofElement[:], proofBytes[:32])
	// 	merkleProof = append(merkleProof, proofElement)
	// }

	// Submit the mint transaction with the prepared proofs
	tx, err := contract.OgMint(auth, merkleProof)
	if err != nil {
		return fmt.Errorf("failed to submit mint transaction: %v", err)
	}
	fmt.Printf("Mint transaction submitted, Tx hash: %s\n", tx.Hash().Hex())

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for the mint transaction to be mined: %v", err)
	}

	// Check the transaction status
	if receipt.Status != 1 {
		return fmt.Errorf("mint transaction failed, status: %v", receipt.Status)
	}

	fmt.Println("Mint transaction successful. NFT minted.")
	fmt.Printf("Transaction hash: %s, Gas used: %d\n", receipt.TxHash.Hex(), receipt.GasUsed)
	return nil
}

func updateMerkleRoot(client *ethclient.Client, contractAddress common.Address, newMerkleRootHex string) error {
	privateKeyHex := os.Getenv("PRIVATE_KEY_HEX")
	if privateKeyHex == "" {
		return fmt.Errorf("private key not found in environment variables")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	chainID, err := client.NetworkID(context.Background()) // Alternatively, set manually: big.NewInt(31337)
	if err != nil {
		return fmt.Errorf("failed to get network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return fmt.Errorf("failed to create authorized transactor: %v", err)
	}

	contract, err := ogwhitelist.NewOgwhitelist(contractAddress, client)
	if err != nil {
		return fmt.Errorf("failed to instantiate the contract: %v", err)
	}

	// Convert newMerkleRootHex to [32]byte
	newMerkleRoot, err := hex.DecodeString(newMerkleRootHex)
	if err != nil {
		return fmt.Errorf("failed to decode new merkle root hex string: %v", err)
	}
	var newMerkleRootFixed [32]byte
	copy(newMerkleRootFixed[:], newMerkleRoot)

	tx, err := contract.UpdateMerkleRoot(auth, newMerkleRootFixed)
	if err != nil {
		return fmt.Errorf("failed to submit transaction to update merkle root: %v", err)
	}

	fmt.Printf("Submitted transaction to update merkle root: %s\n", tx.Hash().Hex())
	return nil
}

func getMerkleRoot(filePath string) ([]byte, error) {
	addressList, err := readAddresses(filePath)
	if err != nil {
		return nil, err
	}

	leaves := hashAddresses(addressList.Addresses)
	return calculateMerkleRoot(leaves), nil
}

func readAddresses(filePath string) (*AddressList, error) {
	file, err := os.ReadFile("addresses.json")
	if err != nil {
		return nil, fmt.Errorf("error reading %s file: %w", filePath, err)
	}

	var addressList AddressList
	if err := json.Unmarshal(file, &addressList); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return &addressList, nil
}

func hashAddresses(addresses []string) [][]byte {
	leaves := make([][]byte, len(addresses))
	for i, address := range addresses {
		hash := sha3.NewLegacyKeccak256()
		hash.Write([]byte(address))
		leaves[i] = hash.Sum(nil)
	}
	return leaves
}

func calculateMerkleRoot(leaves [][]byte) []byte {
	if len(leaves) == 1 {
		return leaves[0]
	}

	var newLevel [][]byte
	for i := 0; i < len(leaves); i += 2 {
		if i+1 < len(leaves) {
			newLevel = append(newLevel, hashPair(leaves[i], leaves[i+1]))
		} else {
			// For an odd number of leaves, duplicate the last leaf
			newLevel = append(newLevel, leaves[i])
		}
	}

	return calculateMerkleRoot(newLevel)
}

func hashPair(left, right []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(left)
	hash.Write(right)
	return hash.Sum(nil)
}

// findHexProofForAnAddress finds the Merkle proof for a given address.
func findHexProofForAnAddress(filePath, address string) ([]string, error) {
	addressList, err := readAddresses(filePath)
	if err != nil {
		return nil, err
	}

	// Hash the target address to match the format in the leaves.
	hash := sha3.NewLegacyKeccak256()
	hash.Write([]byte(address))
	targetHash := hash.Sum(nil)

	// Hash all addresses to create leaves.
	leaves := hashAddresses(addressList.Addresses)

	// Find index of the target hash in leaves.
	index := -1
	for i, leaf := range leaves {
		if hex.EncodeToString(leaf) == hex.EncodeToString(targetHash) {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, fmt.Errorf("address not found in the list")
	}

	// Generate proof for the target address.
	proof, err := generateProof(leaves, index)
	if err != nil {
		return nil, err
	}

	return proof, nil
}

// generateProof generates the Merkle proof for a leaf at a given index.
func generateProof(leaves [][]byte, index int) ([]string, error) {
	// Example proof generation. Replace with actual logic to traverse the tree and collect proof.
	var proof []string

	// Simplified: just indicating positions without actual hashing logic.
	for level := len(leaves); level > 1; level = (level + 1) / 2 {
		if index%2 == 0 && index+1 < level {
			// Example: add right sibling hash.
			proof = append(proof, hex.EncodeToString(leaves[index+1]))
		} else if index > 0 {
			// Example: add left sibling hash.
			proof = append(proof, hex.EncodeToString(leaves[index-1]))
		}
		index /= 2 // Move to the next level.
	}

	return proof, nil
}
