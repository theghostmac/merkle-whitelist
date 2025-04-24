package merkletree

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rs/zerolog/log"
	"sort"
)

// MerkleTree represents a Merkle tree
type MerkleTree struct {
	Root  []byte
	Leafs [][]byte
}

// NewMerkleTree creates a new MerkleTree from a list of addresses.
func NewMerkleTree(addresses []common.Address) (*MerkleTree, error) {
	if len(addresses) == 0 {
		return nil, fmt.Errorf("cannot create MerkleTree with no addresses")
	}

	// Sort addresses to ensure consistent tree structure
	sortedAddresses := make([]common.Address, len(addresses))
	copy(sortedAddresses, addresses)
	sort.Slice(sortedAddresses, func(i, j int) bool {
		return bytes.Compare(sortedAddresses[i].Bytes(), sortedAddresses[j].Bytes()) < 0
	})

	// Create leaf nodes by hashing each address
	leafs := make([][]byte, len(sortedAddresses))
	for i, addr := range sortedAddresses {
		leafs[i] = crypto.Keccak256(addr.Bytes()) // leafs[i] = leaf
	}

	// Calculate the root.
	root := calculateRoot(leafs)

	log.Debug().
		Int("address_count", len(addresses)).
		Str("root", hex.EncodeToString(root)).
		Msg("Created MerkleTree")

	return &MerkleTree{
		Root:  root,
		Leafs: leafs,
	}, nil
}

// calculateRoot calculates the Merkle root from a list of leaf nodes
func calculateRoot(leafs [][]byte) []byte {
	if len(leafs) == 0 {
		return nil
	}

	if len(leafs) == 1 {
		return leafs[0]
	}

	var nextLevel [][]byte

	// Handle odd number of leafs by duplicating the last one
	if len(leafs)%2 != 0 {
		leafs = append(leafs, leafs[len(leafs)-1])
	}

	// Calculate parent nodes
	for i := 0; i < len(leafs); i += 2 {
		node := hashPair(leafs[i], leafs[i+1])
		nextLevel = append(nextLevel, node)
	}

	// Recurse with the new level
	return calculateRoot(nextLevel)
}

// hashPair hashes two nodes together
func hashPair(left, right []byte) []byte {
	// Ensure left < right for consistent ordering
	if bytes.Compare(left, right) > 0 {
		left, right = right, left
	}

	// Concatenate and hash
	data := append(left, right...)
	return crypto.Keccak256(data)
}

// GetProof generates a Merkle proof for a specific address
func (mt *MerkleTree) GetProof(address common.Address) ([][32]byte, error) {
	// Hash the address to get the leaf node
	leaf := crypto.Keccak256(address.Bytes())

	// Find the leaf in the tree
	index := -1
	for i, l := range mt.Leafs {
		if bytes.Equal(l, leaf) {
			index = i
			break
		}
	}

	if index == -1 {
		return nil, fmt.Errorf("address not found in Merkle tree")
	}

	// Generate the proof
	proof := generateProof(mt.Leafs, index)

	// Convert to the required format
	result := make([][32]byte, len(proof))
	for i, p := range proof {
		var fixed [32]byte
		copy(fixed[:], p)
		result[i] = fixed
	}

	log.Debug().
		Str("address", address.Hex()).
		Int("proof_length", len(proof)).
		Msg("Generated Merkle proof")

	return result, nil
}

// generateProof creates a Merkle proof for a leaf at the given index
func generateProof(leafs [][]byte, index int) [][]byte {
	if len(leafs) == 1 {
		return [][]byte{}
	}

	var proof [][]byte
	var nextLevel [][]byte

	// Handle odd number of leafs by duplicating the last one
	if len(leafs)%2 != 0 {
		leafs = append(leafs, leafs[len(leafs)-1])
	}

	// Process current level
	for i := 0; i < len(leafs); i += 2 {
		if i == index || i+1 == index {
			// Add sibling to proof
			if i == index {
				proof = append(proof, leafs[i+1])
			} else {
				proof = append(proof, leafs[i])
			}
		}

		// Create parent node for next level
		parentNode := hashPair(leafs[i], leafs[i+1])
		nextLevel = append(nextLevel, parentNode)
	}

	// Continue with next level if there are more nodes
	if len(nextLevel) > 1 {
		// Calculate new index for the next level
		newIndex := index / 2
		subProof := generateProof(nextLevel, newIndex)
		proof = append(proof, subProof...)
	}

	return proof
}

// VerifyProof verifies a Merkle proof for a specific address
func VerifyProof(root []byte, proof [][]byte, address common.Address) bool {
	// Hash the address to get the leaf node
	leaf := crypto.Keccak256(address.Bytes())
	computedHash := leaf

	// Verify the proof
	for _, proofElement := range proof {
		computedHash = hashPair(computedHash, proofElement)
	}

	return bytes.Equal(computedHash, root)
}

// RootHex returns the Merkle root as a hex string
func (mt *MerkleTree) RootHex() string {
	return hex.EncodeToString(mt.Root)
}

// RootBytes32 returns the Merkle root as a [32]byte
func (mt *MerkleTree) RootBytes32() [32]byte {
	var result [32]byte
	copy(result[:], mt.Root)
	return result
}
