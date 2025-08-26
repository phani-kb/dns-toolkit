#!/bin/bash

# Usage: ./commit-and-auto-merge.sh <files-pattern> <commit-message> <branch-name> <description>

set -euo pipefail

if [[ $# -lt 3 ]]; then
    echo "Usage: $0 <files-pattern> <commit-message> <branch-name> [description]"
    exit 1
fi

FILES_PATTERN="$1"
COMMIT_MESSAGE="$2"
BRANCH_NAME="$3"
DESCRIPTION="${4:-$COMMIT_MESSAGE}"

echo "Files: $FILES_PATTERN"
echo "Message: $COMMIT_MESSAGE"
echo "Branch: $BRANCH_NAME"

if git diff --quiet -- $FILES_PATTERN 2>/dev/null; then
    echo "No changes detected in $FILES_PATTERN, skipping."
    exit 0
fi

echo "Changes detected in $FILES_PATTERN"

echo "Step 1: Committing changes locally..."
git add $FILES_PATTERN
git commit -m "$COMMIT_MESSAGE" || {
    echo "Local commit failed"
    exit 1
}
echo "✅ Changes committed locally"

if [[ -z "${PUSH_TOKEN:-}" ]]; then
    echo "No PUSH_TOKEN available, changes only committed locally"
    exit 0
fi

echo "Step 2: Creating PR with auto-merge..."

# Reset to get the changes back for pushing
git reset --soft HEAD~1

export GITHUB_TOKEN="${PUSH_TOKEN}"

if ./.github/scripts/push-to-main.sh "$FILES_PATTERN" "$DESCRIPTION" "$BRANCH_NAME"; then
    echo "✅ PR created successfully"
    
    if command -v gh &> /dev/null; then
        echo "Step 3: Attempting auto-merge..."
        
        PR_NUMBER=$(gh pr list --head "$BRANCH_NAME" --base main --json number --jq '.[0].number' 2>/dev/null || echo "")
        
        if [[ -n "$PR_NUMBER" ]]; then
            echo "Found PR #$PR_NUMBER, attempting auto-merge..."
            
            if gh pr merge "$PR_NUMBER" --auto --squash --delete-branch; then
                echo "✅ PR #$PR_NUMBER auto-merged successfully"
            else
                echo "⚠️ Auto-merge failed, attempting direct merge..."
                
                if gh pr merge "$PR_NUMBER" --squash --delete-branch; then
                    echo "✅ PR #$PR_NUMBER merged directly and branch deleted"
                else
                    echo "⚠️ Direct merge also failed, PR #$PR_NUMBER needs manual review"
                fi
            fi
        else
            echo "⚠️ Could not find PR number for auto-merge"
        fi
    else
        echo "⚠️ GitHub CLI not available for auto-merge"
    fi
else
    echo "❌ PR creation failed, committing locally"
    git add $FILES_PATTERN
    git commit -m "$COMMIT_MESSAGE" || true
fi

echo "Commit and Auto-Merge completed"
