// Package ogwhitelist provides a wrapper for interacting with the OGWhitelist contract
package ogwhitelist

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rs/zerolog/log"
)

// Wrapper provides a convenient interface to interact with the OGWhitelist contract
type Wrapper struct {
	contract *Ogwhitelist
	backend  bind.ContractBackend
	address  common.Address
}

// NewWrapper creates a new instance of the OGWhitelist contract wrapper
func NewWrapper(backend bind.ContractBackend, address common.Address) (*Wrapper, error) {
	contract, err := NewOgwhitelist(address, backend)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate OGWhitelist contract: %w", err)
	}

	return &Wrapper{
		contract: contract,
		backend:  backend,
		address:  address,
	}, nil
}

// GetMerkleRoot returns the current merkle root stored in the contract
func (w *Wrapper) GetMerkleRoot(ctx context.Context) ([32]byte, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	return w.contract.MerkleRoot(callOpts)
}

// GetMaxSupply returns the maximum supply of NFTs
func (w *Wrapper) GetMaxSupply(ctx context.Context) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	return w.contract.MAXSUPPLY(callOpts)
}

// GetTotalMinted returns the total number of NFTs minted
func (w *Wrapper) GetTotalMinted(ctx context.Context) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	return w.contract.TotalMinted(callOpts)
}

// GetRemainingSupply returns the number of NFTs that can still be minted
func (w *Wrapper) GetRemainingSupply(ctx context.Context) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	return w.contract.RemainingSupply(callOpts)
}

// GetOGMintPrice returns the price for OG minting
func (w *Wrapper) GetOGMintPrice(ctx context.Context) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	return w.contract.OGMINTPRICE(callOpts)
}

// GetPublicMintPrice returns the price for public minting
func (w *Wrapper) GetPublicMintPrice(ctx context.Context) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	return w.contract.PUBLICMINTPRICE(callOpts)
}

// VerifyWhitelistStatus checks if an address is whitelisted
func (w *Wrapper) VerifyWhitelistStatus(ctx context.Context, merkleProof [][32]byte, address common.Address) (bool, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	return w.contract.VerifyWhitelistStatus(callOpts, merkleProof, address)
}

// GetAddressMintedBalance returns the number of NFTs minted by an address
func (w *Wrapper) GetAddressMintedBalance(ctx context.Context, address common.Address) (*big.Int, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	return w.contract.AddressMintedBalance(callOpts, address)
}

// UpdateMerkleRoot updates the merkle root in the contract
func (w *Wrapper) UpdateMerkleRoot(ctx context.Context, auth *bind.TransactOpts, newMerkleRoot [32]byte) (*types.Transaction, error) {
	auth.Context = ctx

	tx, err := w.contract.UpdateMerkleRoot(auth, newMerkleRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to update merkle root: %w", err)
	}

	log.Info().
		Str("tx_hash", tx.Hash().Hex()).
		Str("merkle_root", fmt.Sprintf("%x", newMerkleRoot)).
		Msg("Submitted transaction to update merkle root")

	return tx, nil
}

// OGMint mints an NFT using the OG whitelist
func (w *Wrapper) OGMint(ctx context.Context, auth *bind.TransactOpts, merkleProof [][32]byte) (*types.Transaction, error) {
	auth.Context = ctx

	tx, err := w.contract.OgMint(auth, merkleProof)
	if err != nil {
		return nil, fmt.Errorf("failed to submit OG mint transaction: %w", err)
	}

	log.Info().
		Str("tx_hash", tx.Hash().Hex()).
		Str("from", auth.From.Hex()).
		Msg("Submitted OG mint transaction")

	return tx, nil
}

// PublicMint mints an NFT using the public sale
func (w *Wrapper) PublicMint(ctx context.Context, auth *bind.TransactOpts) (*types.Transaction, error) {
	auth.Context = ctx

	tx, err := w.contract.PublicMint(auth)
	if err != nil {
		return nil, fmt.Errorf("failed to submit public mint transaction: %w", err)
	}

	log.Info().
		Str("tx_hash", tx.Hash().Hex()).
		Str("from", auth.From.Hex()).
		Msg("Submitted public mint transaction")

	return tx, nil
}

// Withdraw withdraws the contract's funds to the specified address
func (w *Wrapper) Withdraw(ctx context.Context, auth *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	auth.Context = ctx

	tx, err := w.contract.Withdraw(auth, to)
	if err != nil {
		return nil, fmt.Errorf("failed to submit withdraw transaction: %w", err)
	}

	log.Info().
		Str("tx_hash", tx.Hash().Hex()).
		Str("to", to.Hex()).
		Msg("Submitted withdraw transaction")

	return tx, nil
}

// ContractAddress returns the address of the contract
func (w *Wrapper) ContractAddress() common.Address {
	return w.address
}
