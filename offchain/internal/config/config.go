// Package config provides configuration management for the application
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

// Config holds the application configuration
type Config struct {
	// Ethereum connection
	RPCEndpoint string         `json:"rpc_endpoint"`
	ChainID     int64          `json:"chain_id"`
	Contracts   ContractConfig `json:"contracts"`

	// Authentication
	PrivateKey string `json:"private_key,omitempty"`

	// Paths
	AddressesFile string `json:"addresses_file"`
}

// ContractConfig holds contract addresses
type ContractConfig struct {
	OGWhitelist string `json:"og_whitelist"`
}

// DefaultConfig returns a default configuration
func DefaultConfig() *Config {
	return &Config{
		RPCEndpoint:   "http://localhost:8545",
		ChainID:       1337, // Default to local development chain
		AddressesFile: "addresses.json",
		Contracts: ContractConfig{
			OGWhitelist: "",
		},
	}
}

// LoadConfig loads configuration from file and environment variables
func LoadConfig(configPath string) (*Config, error) {
	// Start with default config
	config := DefaultConfig()

	// Load from config file if it exists
	if configPath != "" {
		if err := loadFromFile(config, configPath); err != nil {
			return nil, err
		}
	}

	// Override with environment variables
	loadFromEnv(config)

	// Validate configuration
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	return config, nil
}

// loadFromFile loads configuration from a JSON file
func loadFromFile(config *Config, configPath string) error {
	file, err := os.ReadFile(configPath)
	if err != nil {
		// If file doesn't exist, just use defaults
		if os.IsNotExist(err) {
			log.Info().Str("path", configPath).Msg("Configuration file not found, using defaults")
			return nil
		}
		return fmt.Errorf("error reading config file: %w", err)
	}

	if err := json.Unmarshal(file, config); err != nil {
		return fmt.Errorf("error parsing config file: %w", err)
	}

	log.Debug().Str("path", configPath).Msg("Loaded configuration from file")
	return nil
}

// loadFromEnv loads configuration from environment variables and .env file
func loadFromEnv(config *Config) {
	// Load .env file if it exists
	_ = godotenv.Load()

	// RPC endpoint
	if endpoint := os.Getenv("ETH_RPC_ENDPOINT"); endpoint != "" {
		config.RPCEndpoint = endpoint
	}

	// Chain ID
	if chainIDStr := os.Getenv("ETH_CHAIN_ID"); chainIDStr != "" {
		var chainID int64
		_, err := fmt.Sscanf(chainIDStr, "%d", &chainID)
		if err == nil {
			config.ChainID = chainID
		}
	}

	// Contract addresses
	if ogWhitelist := os.Getenv("OG_WHITELIST_CONTRACT"); ogWhitelist != "" {
		config.Contracts.OGWhitelist = ogWhitelist
	}

	// Private key
	if privateKey := os.Getenv("PRIVATE_KEY_HEX"); privateKey != "" {
		config.PrivateKey = privateKey
	}

	// Addresses file
	if addressesFile := os.Getenv("ADDRESSES_FILE"); addressesFile != "" {
		config.AddressesFile = addressesFile
	}
}

// validateConfig validates the configuration
func validateConfig(config *Config) error {
	// Convert paths to absolute if relative
	config.AddressesFile = expandPath(config.AddressesFile)

	return nil
}

// expandPath expands a relative path to an absolute path
func expandPath(path string) string {
	// If it starts with ~, replace with home directory
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err == nil {
			path = filepath.Join(home, path[1:])
		}
	}

	// If path is not absolute, make it absolute
	if !filepath.IsAbs(path) {
		absPath, err := filepath.Abs(path)
		if err == nil {
			path = absPath
		}
	}

	return path
}

// GetOGWhitelistAddress returns the OGWhitelist contract address
func (c *Config) GetOGWhitelistAddress() common.Address {
	return common.HexToAddress(c.Contracts.OGWhitelist)
}

// SaveConfig saves the configuration to a file
func (c *Config) SaveConfig(configPath string) error {
	// Create JSON
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling config: %w", err)
	}

	// Create directory if it doesn't exist
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	// Write to file
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	log.Info().Str("path", configPath).Msg("Configuration saved")
	return nil
}
