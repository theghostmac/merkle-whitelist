package ogwhitelist

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// deployContract deploys the OGWhitelist contract to a simulated blockchain
func deployContract(t *testing.T) (*Wrapper, *bind.TransactOpts, *backends.SimulatedBackend) {
	// Generate a test private key
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	// Create auth for transactions
	chainID := big.NewInt(1337) // Simulated chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	require.NoError(t, err)

	// Create simulated blockchain with initial balance
	balance := new(big.Int)
	balance.SetString("10000000000000000000", 10) // 10 ETH
	address := auth.From
	alloc := make(core.GenesisAlloc)
	alloc[address] = core.GenesisAccount{Balance: balance}
	blockchain := backends.NewSimulatedBackend(alloc, 10000000)

	// Set up initial merkle root
	initialMerkleRoot := [32]byte{}

	// Deploy contract
	contractAddress, _, _, err := DeployOgwhitelist(
		auth,
		blockchain,
		initialMerkleRoot,
		auth.From, // Initial owner is the deployer
	)
	require.NoError(t, err)
	blockchain.Commit()

	// Create wrapper
	wrapper, err := NewWrapper(blockchain, contractAddress)
	require.NoError(t, err)

	return wrapper, auth, blockchain
}

func TestNewWrapper(t *testing.T) {
	wrapper, _, _ := deployContract(t)
	assert.NotNil(t, wrapper)
	assert.NotEqual(t, common.Address{}, wrapper.ContractAddress())
}

func TestGetMerkleRoot(t *testing.T) {
	wrapper, auth, blockchain := deployContract(t)
	ctx := context.Background()

	// Check initial merkle root (should be zero)
	initialRoot, err := wrapper.GetMerkleRoot(ctx)
	require.NoError(t, err)
	assert.Equal(t, [32]byte{}, initialRoot)

	// Update the merkle root
	newRoot := [32]byte{1, 2, 3, 4}
	tx, err := wrapper.UpdateMerkleRoot(ctx, auth, newRoot)
	require.NoError(t, err)
	blockchain.Commit()

	// Check that the updated root is returned
	updatedRoot, err := wrapper.GetMerkleRoot(ctx)
	require.NoError(t, err)
	assert.Equal(t, newRoot, updatedRoot)
	assert.NotNil(t, tx.Hash())
}

func TestGetMaxSupply(t *testing.T) {
	wrapper, _, _ := deployContract(t)
	ctx := context.Background()

	// Get max supply (contract constant)
	maxSupply, err := wrapper.GetMaxSupply(ctx)
	require.NoError(t, err)

	// From the Solidity contract we know MAX_SUPPLY = 1000
	expectedMaxSupply := big.NewInt(1000)
	assert.Equal(t, expectedMaxSupply, maxSupply)
}

func TestGetMintPrices(t *testing.T) {
	wrapper, _, _ := deployContract(t)
	ctx := context.Background()

	// Get OG mint price
	ogPrice, err := wrapper.GetOGMintPrice(ctx)
	require.NoError(t, err)

	// Get public mint price
	publicPrice, err := wrapper.GetPublicMintPrice(ctx)
	require.NoError(t, err)

	// From the contract, OG_MINT_PRICE = 0.0000000000001 ether and PUBLIC_MINT_PRICE = 0.0000000000005 ether
	expectedOGPrice, _ := new(big.Int).SetString("100000000000", 10)     // 0.0000000000001 ether in wei
	expectedPublicPrice, _ := new(big.Int).SetString("500000000000", 10) // 0.0000000000005 ether in wei

	assert.Equal(t, expectedOGPrice, ogPrice)
	assert.Equal(t, expectedPublicPrice, publicPrice)
}

func TestPublicMint(t *testing.T) {
	wrapper, auth, blockchain := deployContract(t)
	ctx := context.Background()

	// Get the public mint price
	publicPrice, err := wrapper.GetPublicMintPrice(ctx)
	require.NoError(t, err)

	// Set value for transaction
	auth.Value = publicPrice

	// Check initial minted count
	initialMinted, err := wrapper.GetTotalMinted(ctx)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(0), initialMinted)

	// Perform public mint
	tx, err := wrapper.PublicMint(ctx, auth)
	require.NoError(t, err)
	blockchain.Commit()

	// Verify the mint was successful
	totalMinted, err := wrapper.GetTotalMinted(ctx)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1), totalMinted)

	// Check remaining supply
	remainingSupply, err := wrapper.GetRemainingSupply(ctx)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(999), remainingSupply)

	// Check minted balance for the address
	balance, err := wrapper.GetAddressMintedBalance(ctx, auth.From)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1), balance)

	assert.NotNil(t, tx.Hash())
}

func TestOGMint(t *testing.T) {
	wrapper, auth, blockchain := deployContract(t)
	ctx := context.Background()

	// Get the OG mint price
	ogPrice, err := wrapper.GetOGMintPrice(ctx)
	require.NoError(t, err)

	// Set value for transaction
	auth.Value = ogPrice

	// For OG mint, we need a valid merkle proof
	// In a real test, we would create a merkle tree with the address included
	// and generate a valid proof
	// For this test, we'll set a mock merkle root that allows our address

	// First, calculate the leaf for our address
	leaf := crypto.Keccak256(auth.From.Bytes())
	mockMerkleRoot := [32]byte{}
	copy(mockMerkleRoot[:], leaf)

	// Update the merkle root to match our calculated leaf
	_, err = wrapper.UpdateMerkleRoot(ctx, auth, mockMerkleRoot)
	require.NoError(t, err)
	blockchain.Commit()

	// Create an empty proof (works in this case since the root == leaf)
	emptyProof := [][32]byte{}

	// Perform OG mint
	tx, err := wrapper.OGMint(ctx, auth, emptyProof)
	require.NoError(t, err)
	blockchain.Commit()

	// Verify the mint was successful
	totalMinted, err := wrapper.GetTotalMinted(ctx)
	require.NoError(t, err)
	assert.Equal(t, big.NewInt(1), totalMinted)

	assert.NotNil(t, tx.Hash())
}

func TestVerifyWhitelistStatus(t *testing.T) {
	wrapper, auth, blockchain := deployContract(t)
	ctx := context.Background()

	// For testing whitelist verification, we'll set a specific merkle root
	// and then verify a valid and invalid address

	// First, calculate the leaf for our address
	leaf := crypto.Keccak256(auth.From.Bytes())
	mockMerkleRoot := [32]byte{}
	copy(mockMerkleRoot[:], leaf)

	// Update the merkle root
	_, err := wrapper.UpdateMerkleRoot(ctx, auth, mockMerkleRoot)
	require.NoError(t, err)
	blockchain.Commit()

	// Create an empty proof (works in this case since the root == leaf)
	emptyProof := [][32]byte{}

	// Verify the address that should be whitelisted
	isWhitelisted, err := wrapper.VerifyWhitelistStatus(ctx, emptyProof, auth.From)
	require.NoError(t, err)
	assert.True(t, isWhitelisted)

	// Verify an address that should not be whitelisted
	otherAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")
	isWhitelisted, err = wrapper.VerifyWhitelistStatus(ctx, emptyProof, otherAddress)
	require.NoError(t, err)
	assert.False(t, isWhitelisted)
}

func TestWithdraw(t *testing.T) {
	wrapper, auth, blockchain := deployContract(t)
	ctx := context.Background()

	// First, we need to add some funds to the contract
	// Let's do a public mint first
	publicPrice, err := wrapper.GetPublicMintPrice(ctx)
	require.NoError(t, err)
	auth.Value = publicPrice
	_, err = wrapper.PublicMint(ctx, auth)
	require.NoError(t, err)
	blockchain.Commit()

	// Reset the value field
	auth.Value = nil

	// Create a new address to withdraw to
	withdrawAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")

	// Call withdraw
	tx, err := wrapper.Withdraw(ctx, auth, withdrawAddress)
	require.NoError(t, err)
	blockchain.Commit()

	assert.NotNil(t, tx.Hash())
}

func TestContractAddress(t *testing.T) {
	wrapper, _, _ := deployContract(t)
	address := wrapper.ContractAddress()
	assert.NotEqual(t, common.Address{}, address)
}
