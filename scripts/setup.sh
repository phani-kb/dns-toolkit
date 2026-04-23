#!/bin/bash

set -e

SCRIPT_DIR=$(dirname -- "$0")
SCRIPT_DIR=$(cd -- "$SCRIPT_DIR" && pwd)

echo "Setting up DNS Toolkit..."

echo "Installing development tools..."
"$SCRIPT_DIR/tools.sh" install-tools

echo "Building dns-toolkit..."

mkdir -p bin

go build -o bin/dns-toolkit .
chmod +x bin/dns-toolkit