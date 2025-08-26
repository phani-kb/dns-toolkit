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
            echo "Found PR #$PR_NUMBER, checking merge readiness..."
            
            # First try auto-merge (will queue for when checks pass)
            if gh pr merge "$PR_NUMBER" --auto --squash --delete-branch; then
                echo "✅ PR #$PR_NUMBER auto-merge enabled (will merge when checks pass)"
            else
                echo "⚠️ Auto-merge failed, checking if we can merge immediately..."
                
                # Check if PR is ready to merge (all checks passed)
                PR_STATUS=$(gh pr view "$PR_NUMBER" --json mergeable,statusCheckRollupState --jq '.mergeable + ":" + .statusCheckRollupState' 2>/dev/null || echo "unknown:unknown")
                MERGEABLE=$(echo "$PR_STATUS" | cut -d: -f1)
                CHECK_STATE=$(echo "$PR_STATUS" | cut -d: -f2)
                
                echo "PR Status: mergeable=$MERGEABLE, checks=$CHECK_STATE"
                
                if [[ "$MERGEABLE" == "MERGEABLE" && "$CHECK_STATE" == "SUCCESS" ]]; then
                    echo "✅ All checks passed, attempting direct merge..."
                    if gh pr merge "$PR_NUMBER" --squash --delete-branch; then
                        echo "✅ PR #$PR_NUMBER merged directly and branch deleted"
                    else
                        echo "⚠️ Direct merge failed despite checks passing"
                    fi
                elif [[ "$CHECK_STATE" == "PENDING" || "$CHECK_STATE" == "EXPECTED" ]]; then
                    echo "⏳ Status checks are still running, waiting briefly..."
                    
                    for i in {1..8}; do
                        sleep 15
                        NEW_STATUS=$(gh pr view "$PR_NUMBER" --json mergeable,statusCheckRollupState --jq '.mergeable + ":" + .statusCheckRollupState' 2>/dev/null || echo "unknown:unknown")
                        NEW_MERGEABLE=$(echo "$NEW_STATUS" | cut -d: -f1)
                        NEW_CHECK_STATE=$(echo "$NEW_STATUS" | cut -d: -f2)
                        
                        echo "  Check $i/8: mergeable=$NEW_MERGEABLE, checks=$NEW_CHECK_STATE"
                        
                        if [[ "$NEW_MERGEABLE" == "MERGEABLE" && "$NEW_CHECK_STATE" == "SUCCESS" ]]; then
                            echo "✅ All checks passed after waiting, attempting merge..."
                            if gh pr merge "$PR_NUMBER" --squash --delete-branch; then
                                echo "✅ PR #$PR_NUMBER merged successfully after waiting"
                                break
                            else
                                echo "⚠️ Merge failed despite checks passing"
                                break
                            fi
                        elif [[ "$NEW_CHECK_STATE" == "FAILURE" ]]; then
                            echo "❌ Status checks failed during wait"
                            break
                        fi
                    done
                    
                    # If we get here and haven't merged, provide final status
                    if [[ "$NEW_CHECK_STATE" == "PENDING" || "$NEW_CHECK_STATE" == "EXPECTED" ]]; then
                        echo "⏳ Checks still running after 2 minutes, PR will auto-merge when ready"
                        echo "   Monitor at: https://github.com/$GITHUB_REPOSITORY/pull/$PR_NUMBER"
                    fi
                elif [[ "$CHECK_STATE" == "FAILURE" ]]; then
                    echo "❌ Status checks failed, PR needs manual review"
                    echo "   Check failures at: https://github.com/$GITHUB_REPOSITORY/pull/$PR_NUMBER"
                else
                    echo "⚠️ PR #$PR_NUMBER needs manual review (status: $PR_STATUS)"
                    echo "   Manual merge may be required due to branch protection rules"
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
