#!/bin/bash

# Script to validate commit messages locally before pushing

set -e

show_usage() {
    cat << EOF
Usage: $0 [OPTIONS]

Options:
  -h, --help      Show this help message
  -r, --range     Check commits in specific range (e.g., HEAD~5..HEAD)
  -b, --branch    Check commits between current branch and specified branch (default: main)
  -n, --number    Check last N commits (default: 10)
  --fix          Show fix instructions for invalid commits

Examples:
  $0              # Check last 10 commits
  $0 -n 5         # Check last 5 commits
  $0 -b main      # Check commits between current branch and main
  $0 -r HEAD~3..HEAD  # Check specific range
  $0 --fix        # Show fix instructions
EOF
}

validate_commit() {
    local message="$1"
    [[ "$message" =~ ^Merge ]] && return 0
    [[ "$message" =~ (#[0-9]+|fix(es)?[[:space:]]*#[0-9]+|close(s)?[[:space:]]*#[0-9]+|resolve(s)?[[:space:]]*#[0-9]+) ]]
}

show_fix_instructions() {
    cat << EOF

HOW TO FIX INVALID COMMITS:

1. Fix most recent commit:
   git commit --amend -m "Your new message #123"

2. Fix older commits:
   git rebase -i <commit-hash>^
   Change 'pick' to 'reword' for invalid commits
   Update commit message to include issue reference

3. If commits are already pushed:
   git push --force-with-lease origin <branch-name>

EOF
}

RANGE=""
BRANCH="main"
NUMBER=10
SHOW_FIX=false

while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help) show_usage; exit 0 ;;
        -r|--range) RANGE="$2"; shift 2 ;;
        -b|--branch) BRANCH="$2"; shift 2 ;;
        -n|--number) NUMBER="$2"; shift 2 ;;
        --fix) SHOW_FIX=true; shift ;;
        *) echo "Unknown option: $1"; show_usage; exit 1 ;;
    esac
done


if [[ -n "$RANGE" ]]; then
    echo "Checking commits in range: $RANGE"
    COMMITS=$(git log --format="%h %s" "$RANGE" 2>/dev/null) || {
        echo "ERROR: Invalid commit range: $RANGE"
        exit 1
    }
elif git merge-base --is-ancestor "$BRANCH" HEAD 2>/dev/null; then
    RANGE="$BRANCH..HEAD"
    echo "Checking commits between $BRANCH and current branch"
    COMMITS=$(git log --format="%h %s" "$RANGE")
else
    echo "Checking last $NUMBER commits"
    COMMITS=$(git log --format="%h %s" -n "$NUMBER")
fi

[[ -z "$COMMITS" ]] && { echo "No commits to validate"; exit 0; }

echo
echo "Analyzing commits..."
echo

INVALID_COMMITS=""
VALID_COUNT=0
SKIPPED_COUNT=0
INVALID_COUNT=0

while IFS= read -r line; do
    [[ -z "$line" ]] && continue
    
    HASH="${line%% *}"
    MESSAGE="${line#* }"
    
    if [[ "$MESSAGE" =~ ^Merge ]]; then
        echo "MERGE (skipped): $HASH $MESSAGE"
        ((SKIPPED_COUNT++))
    elif validate_commit "$MESSAGE"; then
        echo "VALID: $HASH $MESSAGE"
        ((VALID_COUNT++))
    else
        echo "INVALID: $HASH $MESSAGE"
        INVALID_COMMITS+="$HASH: $MESSAGE"$'\n'
        ((INVALID_COUNT++))
    fi
done <<< "$COMMITS"

echo ""
print_header "VALIDATION SUMMARY"
echo "Valid commits: $VALID_COUNT"
echo "Skipped commits (merges): $SKIPPED_COUNT"
echo "Invalid commits: $INVALID_COUNT"
echo

if [[ $INVALID_COUNT -eq 0 ]]; then
    echo "All commit messages are valid!"
    exit 0
else
    echo "Found $INVALID_COUNT invalid commit(s)"
    echo
    echo "The following commits are missing issue references:"
    echo "$INVALID_COMMITS"
    
    [[ "$SHOW_FIX" == true ]] && show_fix_instructions || echo "Run '$0 --fix' to see detailed fix instructions."
    
    echo
    echo "These commits will fail the validate-commits workflow."
    exit 1
fi