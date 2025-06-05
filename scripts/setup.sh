#!/bin/bash

set -e

echo "Setting up DNS Toolkit..."

echo "Building dns-toolkit..."

mkdir -p bin

go build -o bin/dns-toolkit .
chmod +x bin/dns-toolkit