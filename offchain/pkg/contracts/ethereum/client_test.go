package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
	"time"
)

// TestNewClient tests creating a new client
// Note: This test requires a running Ethereum node
// Skip it if not running an integration test
func TestNewClient(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx := context.Background()
	cfg := DefaultConfig()
	client, err := NewClient(ctx, cfg)
	require.NoError(t, err)
	defer client.Close()

	assert.NotNil(t, client.ec)
	assert.NotNil(t, client.chainID)
}

// TestCreateTransactOpts tests creating transaction options
func TestCreateTransactOpts(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx := context.Background()
	cfg := DefaultConfig()
	client, err := NewClient(ctx, cfg)
	require.NoError(t, err)
	defer client.Close()

	// Test private key (only use for testing!)
	privateKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	value := big.NewInt(1000000000000000000) // 1 ETH

	auth, err := client.CreateTransactOpts(ctx, privateKey, value)
	require.NoError(t, err)
	assert.NotNil(t, auth)
	assert.Equal(t, value, auth.Value)
}

// TestGetBalance tests getting the balance of an address
func TestGetBalance(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test in short mode")
	}

	ctx := context.Background()
	cfg := DefaultConfig()
	client, err := NewClient(ctx, cfg)
	require.NoError(t, err)
	defer client.Close()

	// This address should have some ETH in the test network
	address := common.HexToAddress("0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")

	balance, err := client.GetBalance(ctx, address)
	require.NoError(t, err)
	assert.NotNil(t, balance)
}

// TestNewClientWithTimeout tests creating a client with a timeout
func TestNewClientWithTimeout(t *testing.T) {
	ctx := context.Background()

	// Use a non-existent URL to force a timeout
	cfg := ClientConfig{
		RPCURL:  "http://nonexistenturl:8545",
		Timeout: 1 * time.Second,
	}

	_, err := NewClient(ctx, cfg)
	assert.Error(t, err, "Expected error due to timeout/connection failure")
}
