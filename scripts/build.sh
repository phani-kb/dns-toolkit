#!/bin/bash

# go clean -modcache

# Install Go dependencies
echo "Installing Go dependencies..."
go mod tidy

# go get .

# Build the project
echo "Building the project..."
go build -o bin/dns-toolkit

chmod +x bin/dns-toolkit

echo "All dependencies installed and build completed successfully."
