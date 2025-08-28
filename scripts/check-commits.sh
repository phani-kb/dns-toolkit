#!/bin/bash

# Script to validate commit messages locally before pushing
# Based on the same validation logic as .github/workflows/validate-commits.yml

set -e

show_usage() {
    cat << EOF
Usage: $0 [OPTIONS]

Options:
  -h, --help      Show this help message
  -r, --range     Check commits in specific range (e.g., HEAD~5..HEAD)
  -b, --branch    Check commits between current branch and specified branch (default: release/1.0.0)
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
    if echo "$message" | grep -q "^Merge "; then
        return 0
    fi

    if echo "$message" | grep -qE '#[0-9]+|fix(es)?\s*#[0-9]+|close(s)?\s*#[0-9]+|resolve(s)?\s*#[0-9]+'; then
        return 0
    else
        return 1
    fi
}

show_fix_instructions() {
    cat << EOF

HOW TO FIX INVALID COMMITS:

1. Fix most recent commit:
   git commit --amend -m "#123: Your new message"

2. Fix older commits (interactive rebase):
   git rebase -i <commit-hash>^
   # Change 'pick' to 'reword' for invalid commits
   # Update commit message to include issue reference

3. Fix multiple commits with soft reset:
   git reset --soft HEAD~N  # N = number of commits to undo
   git commit -m "#123: Combined fix message"

4. If commits are already pushed:
   git push --force-with-lease origin <branch-name>

EOF
}

RANGE=""
BRANCH="release/1.0.0"
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
    if [[ -z "$COMMITS" ]]; then
        echo "No commits found between $BRANCH and current branch"
        echo "Current branch is up to date with $BRANCH"
        exit 0
    fi
else
    echo "Checking last $NUMBER commits"
    COMMITS=$(git log --format="%h %s" -n "$NUMBER")
fi

if [[ -z "$COMMITS" ]]; then
    echo "No commits to validate"
    exit 0
fi

echo
echo "Found $(echo "$COMMITS" | wc -l) commit(s) to analyze..."
echo

INVALID_COMMITS=""
VALID_COUNT=0
SKIPPED_COUNT=0
INVALID_COUNT=0

while IFS= read -r line; do
    [[ -z "$line" ]] && continue
    
    HASH="${line%% *}"
    MESSAGE="${line#* }"
    
    if echo "$MESSAGE" | grep -q "^Merge "; then
        echo "MERGE (skipped): $HASH $MESSAGE"
        SKIPPED_COUNT=$((SKIPPED_COUNT + 1))
    elif validate_commit "$MESSAGE"; then
        echo "VALID: $HASH $MESSAGE"
        VALID_COUNT=$((VALID_COUNT + 1))
    else
        echo "INVALID: $HASH $MESSAGE"
        if [[ -n "$INVALID_COMMITS" ]]; then
            INVALID_COMMITS+=$'\n'
        fi
        INVALID_COMMITS+="$HASH: $MESSAGE"
        INVALID_COUNT=$((INVALID_COUNT + 1))
    fi
done <<< "$COMMITS"

echo
echo "Valid commits: $VALID_COUNT"
echo "Skipped commits (merges): $SKIPPED_COUNT"
echo "Invalid commits: $INVALID_COUNT"
echo "Total commits analyzed: $((VALID_COUNT + SKIPPED_COUNT + INVALID_COUNT))"
echo

if [[ $INVALID_COUNT -eq 0 ]]; then
    echo "All commit messages are valid! The validate-commits workflow will pass."
    exit 0
else
    echo "Found $INVALID_COUNT invalid commit(s) that will fail the validate-commits workflow:"
    echo
    echo "$INVALID_COMMITS"
    echo
    
    if [[ "$SHOW_FIX" == true ]]; then
        show_fix_instructions
    else
        echo "Run '$0 --fix' to see detailed fix instructions."
    fi
    
    echo
    echo "These commits will fail the validate-commits workflow on GitHub."
    exit 1
fi