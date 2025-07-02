#!/bin/bash
set -e

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

THRESHOLD=80

# Timestamp for archiving reports
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")

# Create coverage folder if it doesn't exist
echo -e "${GREEN}Creating coverage directory...${NC}"
mkdir -p coverage
mkdir -p coverage/archive

echo -e "${GREEN}Running tests with coverage...${NC}"

export DNS_TOOLKIT_TEST_MODE=true
export DNS_TOOLKIT_TEST_CONFIG_PATH=$(pwd)/testdata/config.yml

PACKAGES=$(go list ./... | grep -v "/mocks" | grep -v "constants")
if ! go test -coverprofile=coverage/coverage.out -covermode=atomic $PACKAGES; then
    echo -e "${RED}Tests failed! See output above for details.${NC}"
    exit 1
fi

unset DNS_TOOLKIT_TEST_CONFIG_PATH
unset DNS_TOOLKIT_TEST_MODE

echo -e "${GREEN}Filtering coverage report...${NC}"
grep -v "test_helpers.go" coverage/coverage.out > coverage/filtered_coverage.out
mv coverage/filtered_coverage.out coverage/coverage.out

# Generate HTML report
echo -e "${GREEN}Generating HTML coverage report at coverage/coverage.html...${NC}"
go tool cover -html=coverage/coverage.out -o coverage/coverage.html

# Archive coverage report
cp coverage/coverage.html coverage/archive/coverage_${TIMESTAMP}.html
cp coverage/coverage.out coverage/archive/coverage_${TIMESTAMP}.out

# Generate a summary of the coverage
#echo -e "${YELLOW}Coverage Summary (sorted by coverage):${NC}"
#go tool cover -func=coverage/coverage.out | grep -v "^total:" | sort -k3 -r | head -n 20

echo -e "${YELLOW}Packages with lowest coverage:${NC}"
go tool cover -func=coverage/coverage.out | grep -v "^total:" | sort -k3 | head -n 10

# Calculate total coverage
TOTAL_COVERAGE=$(go tool cover -func=coverage/coverage.out | grep total: | awk '{print $3}')
COVERAGE_NUM=$(echo "$TOTAL_COVERAGE" | sed 's/%//')

echo ""
if awk "BEGIN {exit ($COVERAGE_NUM < $THRESHOLD) ? 0 : 1}"; then
    echo -e "${RED}Total coverage (${TOTAL_COVERAGE}) is below the threshold (${THRESHOLD}%)!${NC}"
    exit 1
else
    echo -e "${GREEN}Total coverage: ${TOTAL_COVERAGE} (threshold: ${THRESHOLD}%)${NC}"
fi
