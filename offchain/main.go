package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/crypto/sha3"
)

type AddressList struct {
	Addresses []string `json:"addresses"`
}

func main() {
	// Read the JSON file containing the list of addresses.
	file, err := os.ReadFile("addresses.json")
	if err != nil {
		fmt.Println("Error reading addresses.json file: ", err)
		os.Exit(1)
	}

	var addressList AddressList
	err = json.Unmarshal(file, &addressList)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		os.Exit(1)
	}

	// Hash the addresses using keccak256 to create the leaves.
	leaves := make([][]byte, len(addressList.Addresses))
	for i, address := range addressList.Addresses {
		hash := sha3.NewLegacyKeccak256()
		hash.Write([]byte(address))
		leaves[i] = hash.Sum(nil)
	}

	// Construct the Merkle tree and calculate the root.
	merkleRoot := calculateMerkleRoot(leaves)

	fmt.Println(
		"Merkle Root: ", hex.EncodeToString(merkleRoot),
	)
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