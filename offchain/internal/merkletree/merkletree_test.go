package merkletree

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMerkleTree(t *testing.T) {
	// Test with valid addresses
	addresses := []common.Address{
		common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
		common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"),
	}

	tree, err := NewMerkleTree(addresses)
	require.NoError(t, err)
	assert.NotNil(t, tree)
	assert.NotEmpty(t, tree.Root)
	assert.Len(t, tree.Leafs, 3)

	// Test with empty list
	_, err = NewMerkleTree([]common.Address{})
	assert.Error(t, err)

	// Test with single address
	singleAddress := []common.Address{
		common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
	}
	tree, err = NewMerkleTree(singleAddress)
	require.NoError(t, err)
	assert.NotNil(t, tree)
	assert.NotEmpty(t, tree.Root)
	assert.Len(t, tree.Leafs, 1)
}

func TestGetProof(t *testing.T) {
	// Create a tree with known addresses
	addresses := []common.Address{
		common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
		common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"),
		common.HexToAddress("0x90F79bf6EB2c4f870365E785982E1f101E93b906"),
	}

	tree, err := NewMerkleTree(addresses)
	require.NoError(t, err)

	// Get proof for an address that exists
	proof, err := tree.GetProof(addresses[1])
	require.NoError(t, err)
	assert.NotEmpty(t, proof)

	// Verify the proof works
	isValid := VerifyProof(tree.Root, convertProof(proof), addresses[1])
	assert.True(t, isValid)

	// Get proof for address that doesn't exist
	nonExistentAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	_, err = tree.GetProof(nonExistentAddr)
	assert.Error(t, err)
}

func TestVerifyProof(t *testing.T) {
	// Create a tree with several addresses
	addresses := []common.Address{
		common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
		common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"),
		common.HexToAddress("0x90F79bf6EB2c4f870365E785982E1f101E93b906"),
		common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"),
	}

	tree, err := NewMerkleTree(addresses)
	require.NoError(t, err)

	// Test verification for each address
	for _, addr := range addresses {
		proof, err := tree.GetProof(addr)
		require.NoError(t, err)

		// Convert proof format for verification
		proofBytes := convertProof(proof)

		// Verify the correct address succeeds
		isValid := VerifyProof(tree.Root, proofBytes, addr)
		assert.True(t, isValid, "Proof should be valid for address %s", addr.Hex())

		// Verify a different address fails
		wrongAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
		isValid = VerifyProof(tree.Root, proofBytes, wrongAddr)
		assert.False(t, isValid, "Proof should be invalid for wrong address")
	}
}

func TestRootFormats(t *testing.T) {
	addresses := []common.Address{
		common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"),
		common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
	}

	tree, err := NewMerkleTree(addresses)
	require.NoError(t, err)

	// Test hex format
	hexRoot := tree.RootHex()
	assert.NotEmpty(t, hexRoot)
	assert.Len(t, hexRoot, 64) // 32 bytes = 64 hex chars

	// Test bytes32 format
	bytes32Root := tree.RootBytes32()
	assert.Len(t, bytes32Root[:], 32)
}

// Helper function to convert [][32]byte to [][]byte for testing
func convertProof(proof [][32]byte) [][]byte {
	result := make([][]byte, len(proof))
	for i, p := range proof {
		result[i] = p[:]
	}
	return result
}
