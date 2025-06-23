#!/bin/bash

# go clean -modcache

# Install Go dependencies
echo "Installing Go dependencies..."
go mod tidy

# go get .

# Build and check all packages
echo "Checking all packages..."
if ! go build ./...; then
    echo "Build failed - there are compilation errors"
    exit 1
fi

# Build the executable
echo "Building the project..."
go build -o bin/ ./...

chmod +x bin/dns-toolkit

echo "All dependencies installed and build completed successfully."
