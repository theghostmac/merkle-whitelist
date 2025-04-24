#!/bin/bash
set -e

# This script generates Go bindings for the OGWhitelist Solidity contract
# It handles compilation with remappings for OpenZeppelin

# Define paths
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" &>/dev/null && pwd)"
PROJECT_ROOT="$(dirname "$(dirname "$SCRIPT_DIR")")" # Go up two levels from scripts dir
CONTRACT_PATH="$PROJECT_ROOT/src/OGWhitelist.sol"
OUTPUT_DIR="$PROJECT_ROOT/offchain/pkg/contracts/ogwhitelist"
TEMP_DIR=$(mktemp -d)
FIXED_CONTRACT="$TEMP_DIR/OGWhitelist_fixed.sol"

echo "Project root: $PROJECT_ROOT"
echo "Contract path: $CONTRACT_PATH"
echo "Output directory: $OUTPUT_DIR"
echo "Temp directory: $TEMP_DIR"

# Create output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# First, fix the visibility issue and create a modified contract
echo "Creating fixed contract version..."
cat "$CONTRACT_PATH" | sed 's/external public/public/g' > "$FIXED_CONTRACT"

# Check if OpenZeppelin is installed via npm
if [ ! -d "node_modules/@openzeppelin" ]; then
    echo "OpenZeppelin contracts not found, installing..."
    npm init -y
    npm install @openzeppelin/contracts
fi

# Compile the contract with remappings
echo "Compiling contract..."
solc --abi --bin \
    --overwrite \
    --optimize \
    --output-dir "$TEMP_DIR" \
    --base-path . \
    --include-path node_modules/ \
    --include-path lib/ \
    --via-ir \
    "$FIXED_CONTRACT"

# Find the compiled files
CONTRACT_NAME=$(basename "$FIXED_CONTRACT" .sol)
ABI_FILE="$TEMP_DIR/$CONTRACT_NAME.abi"
BIN_FILE="$TEMP_DIR/$CONTRACT_NAME.bin"

if [ ! -f "$ABI_FILE" ] || [ ! -f "$BIN_FILE" ]; then
    echo "Error: Compilation failed or output files not found"
    echo "ABI should be at: $ABI_FILE"
    echo "BIN should be at: $BIN_FILE"
    exit 1
fi

echo "Found ABI file: $ABI_FILE"
echo "Found BIN file: $BIN_FILE"

# Generate Go bindings
echo "Generating Go bindings..."
abigen --abi="$ABI_FILE" \
       --bin="$BIN_FILE" \
       --pkg=ogwhitelist \
       --type=Ogwhitelist \
       --out="$OUTPUT_DIR/ogwhitelist.go"

echo "Go bindings generated successfully!"
echo "Output: $OUTPUT_DIR/ogwhitelist.go"

# Clean up temporary directory
rm -rf "$TEMP_DIR"