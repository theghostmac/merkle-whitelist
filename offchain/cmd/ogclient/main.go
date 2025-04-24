// OG Whitelist Client is a command-line tool for interacting with the OGWhitelist contract
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"merkle-whitelist/offchain/internal/merkletree"
	"merkle-whitelist/offchain/pkg/contracts/ethereum"
	"merkle-whitelist/offchain/pkg/contracts/ogwhitelist"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	defaultRPCURL        = "http://localhost:8545"
	defaultAddressesFile = "addresses.json"
)

// AddressList represents the whitelist addresses JSON structure
type AddressList struct {
	Addresses []string `json:"addresses"`
}

func main() {
	// Set up logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Parse command line arguments
	rpcURL := flag.String("rpc", defaultRPCURL, "Ethereum RPC URL")
	contractAddrStr := flag.String("contract", "", "OGWhitelist contract address")
	privateKey := flag.String("key", "", "Private key for transactions (or use PRIVATE_KEY_HEX env var)")
	addressesFile := flag.String("addresses", defaultAddressesFile, "JSON file containing whitelisted addresses")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")

	// Actions
	generateMerkleRoot := flag.Bool("generate-root", false, "Generate Merkle root from addresses file")
	updateMerkleRoot := flag.Bool("update-root", false, "Update Merkle root in contract")
	verifyAddress := flag.String("verify", "", "Verify if address is whitelisted")
	mint := flag.String("mint", "", "Mint an NFT for the specified address (or 'public' for public mint)")
	status := flag.Bool("status", false, "Get contract status")

	flag.Parse()

	// Set log level
	if *verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Get private key from flag or environment
	keyToUse := *privateKey
	if keyToUse == "" {
		keyToUse = os.Getenv("PRIVATE_KEY_HEX")
		if keyToUse == "" && (*updateMerkleRoot || *mint != "") {
			log.Fatal().Msg("Private key required for transactions. Use -key flag or PRIVATE_KEY_HEX env var")
		}
	}

	// Create context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown
	setupSignalHandler(cancel)

	// Process commands
	if *generateMerkleRoot {
		handleGenerateMerkleRoot(*addressesFile)
		return
	}

	// Commands requiring contract interaction
	if *contractAddrStr == "" && (*updateMerkleRoot || *verifyAddress != "" || *mint != "" || *status) {
		log.Fatal().Msg("Contract address required. Use -contract flag")
	}

	if *updateMerkleRoot || *verifyAddress != "" || *mint != "" || *status {
		contractAddr := common.HexToAddress(*contractAddrStr)
		client, wrapper, err := setupContractWrapper(ctx, *rpcURL, contractAddr)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to set up contract wrapper")
		}
		defer client.Close()

		if *status {
			handleStatus(ctx, wrapper)
			return
		}

		if *updateMerkleRoot {
			handleUpdateMerkleRoot(ctx, wrapper, client, keyToUse, *addressesFile)
			return
		}

		if *verifyAddress != "" {
			handleVerifyAddress(ctx, wrapper, *verifyAddress, *addressesFile)
			return
		}

		if *mint != "" {
			handleMint(ctx, wrapper, client, keyToUse, *mint, *addressesFile)
			return
		}
	}

	// If no command provided, show usage
	flag.Usage()
}

// setupSignalHandler sets up a handler for interrupt signals
func setupSignalHandler(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Info().Msg("Received interrupt signal, shutting down...")
		cancel()
		os.Exit(0)
	}()
}

// setupContractWrapper creates the Ethereum client and contract wrapper
func setupContractWrapper(ctx context.Context, rpcURL string, contractAddr common.Address) (*ethereum.Client, *ogwhitelist.Wrapper, error) {
	// Create Ethereum client
	cfg := ethereum.ClientConfig{
		RPCURL:  rpcURL,
		Timeout: 30 * time.Second,
	}

	client, err := ethereum.NewClient(ctx, cfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create Ethereum client: %w", err)
	}

	// Create contract wrapper
	wrapper, err := ogwhitelist.NewWrapper(client.EthClient(), contractAddr)
	if err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("failed to create contract wrapper: %w", err)
	}

	return client, wrapper, nil
}

// handleGenerateMerkleRoot generates a Merkle root from the addresses file
func handleGenerateMerkleRoot(addressesFile string) {
	addresses, err := loadAddresses(addressesFile)
	if err != nil {
		log.Fatal().Err(err).Str("file", addressesFile).Msg("Failed to load addresses")
	}

	// Convert string addresses to common.Address
	ethAddresses := make([]common.Address, len(addresses))
	for i, addr := range addresses {
		ethAddresses[i] = common.HexToAddress(addr)
	}

	// Create Merkle tree
	tree, err := merkletree.NewMerkleTree(ethAddresses)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Merkle tree")
	}

	// Output the Merkle root
	fmt.Printf("Merkle Root: 0x%s\n", tree.RootHex())
}

// handleStatus gets the current status of the contract
func handleStatus(ctx context.Context, wrapper *ogwhitelist.Wrapper) {
	// Get Merkle root
	merkleRoot, err := wrapper.GetMerkleRoot(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get Merkle root")
	} else {
		fmt.Printf("Current Merkle Root: 0x%x\n", merkleRoot)
	}

	// Get supply information
	maxSupply, err := wrapper.GetMaxSupply(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get max supply")
	}

	totalMinted, err := wrapper.GetTotalMinted(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get total minted")
	}

	remainingSupply, err := wrapper.GetRemainingSupply(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get remaining supply")
	}

	// Get pricing information
	ogPrice, err := wrapper.GetOGMintPrice(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get OG mint price")
	}

	publicPrice, err := wrapper.GetPublicMintPrice(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get public mint price")
	}

	// Print the information
	fmt.Println("\nContract Status:")
	fmt.Printf("Max Supply: %s\n", maxSupply)
	fmt.Printf("Total Minted: %s\n", totalMinted)
	fmt.Printf("Remaining Supply: %s\n", remainingSupply)
	fmt.Printf("OG Mint Price: %s Ether\n", formatEther(ogPrice))
	fmt.Printf("Public Mint Price: %s Ether\n", formatEther(publicPrice))
}

// handleUpdateMerkleRoot updates the Merkle root in the contract
func handleUpdateMerkleRoot(ctx context.Context, wrapper *ogwhitelist.Wrapper, client *ethereum.Client, privateKey, addressesFile string) {
	addresses, err := loadAddresses(addressesFile)
	if err != nil {
		log.Fatal().Err(err).Str("file", addressesFile).Msg("Failed to load addresses")
	}

	// Convert string addresses to common.Address
	ethAddresses := make([]common.Address, len(addresses))
	for i, addr := range addresses {
		ethAddresses[i] = common.HexToAddress(addr)
	}

	// Create Merkle tree
	tree, err := merkletree.NewMerkleTree(ethAddresses)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Merkle tree")
	}

	newRoot := tree.RootBytes32()
	fmt.Printf("Updating Merkle Root to: 0x%s\n", tree.RootHex())

	// Create transaction options
	auth, err := client.CreateTransactOpts(ctx, privateKey, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create transaction options")
	}

	// Submit the transaction
	tx, err := wrapper.UpdateMerkleRoot(ctx, auth, newRoot)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to update Merkle root")
	}

	// Wait for transaction to be mined
	fmt.Printf("Transaction submitted: %s\n", tx.Hash().Hex())
	receipt, err := client.WaitForTx(ctx, tx)
	if err != nil {
		log.Fatal().Err(err).Msg("Transaction failed")
	}

	fmt.Printf("Merkle root updated successfully. Gas used: %d\n", receipt.GasUsed)
}

// handleVerifyAddress checks if an address is whitelisted
func handleVerifyAddress(ctx context.Context, wrapper *ogwhitelist.Wrapper, addressToVerify, addressesFile string) {
	addresses, err := loadAddresses(addressesFile)
	if err != nil {
		log.Fatal().Err(err).Str("file", addressesFile).Msg("Failed to load addresses")
	}

	// Convert string addresses to common.Address
	ethAddresses := make([]common.Address, len(addresses))
	for i, addr := range addresses {
		ethAddresses[i] = common.HexToAddress(addr)
	}

	// Create Merkle tree
	tree, err := merkletree.NewMerkleTree(ethAddresses)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Merkle tree")
	}

	// Get the address to verify
	addr := common.HexToAddress(addressToVerify)

	// Generate proof
	proof, err := tree.GetProof(addr)
	if err != nil {
		log.Info().Str("address", addr.Hex()).Msg("Address is not in the whitelist")
		return
	}

	// Verify on-chain
	isWhitelisted, err := wrapper.VerifyWhitelistStatus(ctx, proof, addr)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to verify whitelist status")
	}

	if isWhitelisted {
		fmt.Printf("Address %s is whitelisted\n", addr.Hex())
	} else {
		fmt.Printf("Address %s is NOT whitelisted (proof invalid on-chain)\n", addr.Hex())
	}
}

// handleMint performs minting an NFT
func handleMint(ctx context.Context, wrapper *ogwhitelist.Wrapper, client *ethereum.Client, privateKey, mintType, addressesFile string) {
	// Get mint prices
	ogPrice, err := wrapper.GetOGMintPrice(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get OG mint price")
	}

	publicPrice, err := wrapper.GetPublicMintPrice(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get public mint price")
	}

	// Handle public mint
	if strings.ToLower(mintType) == "public" {
		// Create transaction options with public mint price
		auth, err := client.CreateTransactOpts(ctx, privateKey, publicPrice)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create transaction options")
		}

		// Submit mint transaction
		fmt.Println("Submitting public mint transaction...")
		tx, err := wrapper.PublicMint(ctx, auth)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to submit public mint transaction")
		}

		// Wait for transaction to be mined
		fmt.Printf("Transaction submitted: %s\n", tx.Hash().Hex())
		receipt, err := client.WaitForTx(ctx, tx)
		if err != nil {
			log.Fatal().Err(err).Msg("Transaction failed")
		}

		fmt.Printf("NFT minted successfully. Gas used: %d\n", receipt.GasUsed)
		return
	}

	// Handle OG mint (address provided)
	addresses, err := loadAddresses(addressesFile)
	if err != nil {
		log.Fatal().Err(err).Str("file", addressesFile).Msg("Failed to load addresses")
	}

	// Convert string addresses to common.Address
	ethAddresses := make([]common.Address, len(addresses))
	for i, addr := range addresses {
		ethAddresses[i] = common.HexToAddress(addr)
	}

	// Create Merkle tree
	tree, err := merkletree.NewMerkleTree(ethAddresses)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Merkle tree")
	}

	// Get the address to mint for
	addr := common.HexToAddress(mintType)

	// Generate proof
	proof, err := tree.GetProof(addr)
	if err != nil {
		log.Fatal().Err(err).Str("address", addr.Hex()).Msg("Address is not in the whitelist")
	}

	// Create transaction options with OG mint price
	auth, err := client.CreateTransactOpts(ctx, privateKey, ogPrice)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create transaction options")
	}

	// Submit mint transaction
	fmt.Printf("Submitting OG mint transaction for address %s...\n", addr.Hex())
	tx, err := wrapper.OGMint(ctx, auth, proof)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to submit OG mint transaction")
	}

	// Wait for transaction to be mined
	fmt.Printf("Transaction submitted: %s\n", tx.Hash().Hex())
	receipt, err := client.WaitForTx(ctx, tx)
	if err != nil {
		log.Fatal().Err(err).Msg("Transaction failed")
	}

	fmt.Printf("NFT minted successfully for address %s. Gas used: %d\n", addr.Hex(), receipt.GasUsed)
}

// loadAddresses loads addresses from a JSON file
func loadAddresses(filePath string) ([]string, error) {
	// Read file content
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	// Parse JSON
	var addressList AddressList
	if err := json.Unmarshal(fileContent, &addressList); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return addressList.Addresses, nil
}

// formatEther formats a wei amount as ether
func formatEther(wei *big.Int) string {
	if wei == nil {
		return "unknown"
	}

	// Convert wei to a decimal representation of ether
	// This is a simplified conversion - for production code,
	// consider using a dedicated package for precise math
	f := new(big.Float).SetInt(wei)
	f.Quo(f, new(big.Float).SetInt(big.NewInt(1e18))) // Divide by 10^18

	return f.Text('f', 18)
}
