package ethereum

import "C"
import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"math/big"
	"time"
)

// Client represents an enhanced Ethereum client with common operations
type Client struct {
	ec      *ethclient.Client
	chainID *big.Int
}

// ClientConfig holds configuration for creating a new Ethereum client
type ClientConfig struct {
	RPCURL  string
	ChainID *big.Int
	Timeout time.Duration
}

// DefaultConfig returns a default configuration for the client
func DefaultConfig() ClientConfig {
	return ClientConfig{
		RPCURL:  "http://localhost:8545",
		Timeout: 10 * time.Second,
	}
}

// NewClient creates a new Ethereum client with the given configuration
func NewClient(ctx context.Context, cfg ClientConfig) (*Client, error) {
	if cfg.Timeout == 0 {
		cfg.Timeout = 10 * time.Second
	}

	dialCtx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	ec, err := ethclient.DialContext(dialCtx, cfg.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: %w", err)
	}

	// If chainID wasn't provided, retrieve it from the network.
	if cfg.ChainID == nil {
		networkCtx, networkCancel := context.WithTimeout(ctx, cfg.Timeout)
		defer networkCancel()

		chainID, err := ec.NetworkID(networkCtx)
		if err != nil {
			return nil, fmt.Errorf("failed to retrieve chain ID: %w", err)
		}
		cfg.ChainID = chainID
	}

	log.Debug().
		Str("rpc_url", cfg.RPCURL).
		Str("chain_id", cfg.ChainID.String()).
		Msg("Connected to Ethereum client")

	return &Client{
		ec:      ec,
		chainID: cfg.ChainID,
	}, nil
}

// CreateTransactOpts creates authorized transaction options from a private key.
func (c *Client) CreateTransactOpts(ctx context.Context, privateKeyHex string, value *big.Int) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create keyed transactor: %w", err)
	}

	if value != nil {
		auth.Value = value
	}

	return auth, nil
}

// WaitForTx waits for a transaction to be mined and returns the receipt.
func (c *Client) WaitForTx(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(ctx, c.ec, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for transaction to be mined: %w", err)
	}

	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("transaction failed with status: %d", receipt.Status)
	}

	return receipt, nil
}

// GetBalance returns the balance of an address.
func (c *Client) GetBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	balance, err := c.ec.BalanceAt(ctx, address, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %w", err)
	}

	return balance, nil
}

// EthClient returns the underlying ethclient.Client
func (c *Client) EthClient() *ethclient.Client {
	return c.ec
}

// Close closes the connection to the Ethereum node
func (c *Client) Close() {
	c.ec.Close()
}
