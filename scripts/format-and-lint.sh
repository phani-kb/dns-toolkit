#!/bin/sh

set -e

echo "Running golines..."
find . -name '*.go' | xargs golines --max-len=120 --base-formatter=gofumpt -w

echo "Running golangci-lint..."
golangci-lint run
