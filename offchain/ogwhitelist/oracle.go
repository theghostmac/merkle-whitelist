package ogwhitelist

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	privateKey      = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	rpcURL          = "http://localhost:8545"
	contractAddress = common.HexToAddress("0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512")
)

// Oracle contains the methods to interact with the smart contract
type Oracle struct {
	client       *ethclient.Client
	whitelist    *Ogwhitelist
	transactOpts *bind.TransactOpts
}

// NewOracle creates a new oracle with a connected Ethereum client
func NewOracle() (*Oracle, error) {
	// Connect to an Ethereum node
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}

	// Instantiate the contract
	whitelist, err := NewOgwhitelist(contractAddress, client)
	if err != nil {
		return nil, err
	}

	// Set up authentication (transactOpts) here using the private key
	// ...

	return &Oracle{
		client:    client,
		whitelist: whitelist,
		//transactOpts: transactOpts,
	}, nil
}

// OgMint wraps the ogMint function of the smart contract
func (o *Oracle) OgMint(proof [][32]byte) error {
	tx, err := o.whitelist.OgMint(o.transactOpts, proof)
	if err != nil {
		return err
	}

	// Wait for the transaction to be mined
	_, err = bind.WaitMined(context.Background(), o.client, tx)
	if err != nil {
		return err
	}

	return nil
}

// PublicMint wraps the publicMint function of the smart contract
func (o *Oracle) PublicMint() error {
	tx, err := o.whitelist.PublicMint(o.transactOpts)
	if err != nil {
		return err
	}

	_, err = bind.WaitMined(context.Background(), o.client, tx)
	if err != nil {
		return err
	}

	return nil
}
