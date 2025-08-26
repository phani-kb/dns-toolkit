#!/bin/bash

set -euo pipefail

# Usage: ./push-to-main.sh <files-pattern> <commit-message> [branch-name]

if [[ $# -lt 2 ]]; then
    echo "Usage: $0 <files-pattern> <commit-message> [branch-name]"
    exit 1
fi

FILES_PATTERN="$1"
COMMIT_MESSAGE="$2"
BRANCH_NAME="${3:-update-main}"  # default update-main branch

if [[ -z "${PUSH_TOKEN}" ]]; then
    echo "Error: PUSH_TOKEN environment variable is not set"
    exit 1
fi

if [[ -z "${GITHUB_REPOSITORY}" ]]; then
    echo "Error: GITHUB_REPOSITORY environment variable is not set"
    exit 1
fi

if [[ -z "${GITHUB_RUN_NUMBER}" ]]; then
    echo "Warning: GITHUB_RUN_NUMBER not set, using 'manual'"
    GITHUB_RUN_NUMBER="manual"
fi

echo "Files Pattern: ${FILES_PATTERN}"
echo "Commit Message: ${COMMIT_MESSAGE}"
echo "Branch Name: ${BRANCH_NAME}"
echo "Repository: ${GITHUB_REPOSITORY}"

echo "Checking for changes in: ${FILES_PATTERN}"
if ! git diff --quiet -- ${FILES_PATTERN} 2>/dev/null || ! git diff --cached --quiet -- ${FILES_PATTERN} 2>/dev/null; then
    echo "Changes detected"
else
    UNTRACKED_FILES=$(git ls-files --others --exclude-standard -- ${FILES_PATTERN} 2>/dev/null || true)
    if [[ -z "${UNTRACKED_FILES}" ]]; then
        echo "No changes detected in ${FILES_PATTERN}, skipping."
        exit 0
    fi
    echo "New untracked files detected"
fi

echo "Configuring git..."
git config user.name "GitHub Actions Bot"
git config user.email "actions@github.com"

CURRENT_BRANCH=$(git rev-parse --abbrev-ref HEAD)
echo "Current branch: ${CURRENT_BRANCH}"

git remote set-url origin "https://x-access-token:${PUSH_TOKEN}@github.com/${GITHUB_REPOSITORY}.git"

echo "Fetching latest from origin..."
git fetch origin

if git ls-remote --heads origin "${BRANCH_NAME}" | grep -q "${BRANCH_NAME}"; then
    echo "Branch '${BRANCH_NAME}' exists remotely, checking out..."
    
    if [[ "${CURRENT_BRANCH}" == "${BRANCH_NAME}" ]]; then
        echo "Already on branch ${BRANCH_NAME}"
        git pull origin "${BRANCH_NAME}" --rebase || {
            echo "Pull failed, attempting to resolve..."
            git rebase --abort 2>/dev/null || true
            git pull origin "${BRANCH_NAME}" --no-rebase
        }
    else
        git stash push -m "Temporary stash for branch switch" || true
        
        git checkout "${BRANCH_NAME}" 2>/dev/null || {
            echo "Local branch doesn't exist, creating from remote..."
            git checkout -b "${BRANCH_NAME}" "origin/${BRANCH_NAME}"
        }
        
        git pull origin "${BRANCH_NAME}" --rebase || {
            echo "Pull failed, attempting to resolve..."
            git rebase --abort 2>/dev/null || true
            git pull origin "${BRANCH_NAME}" --no-rebase
        }
        
        git stash pop 2>/dev/null || true
    fi
else
    echo "Branch '${BRANCH_NAME}' doesn't exist, creating new branch..."
    
    if [[ "${CURRENT_BRANCH}" != "${BRANCH_NAME}" ]]; then
        git stash push -m "Temporary stash for branch creation" || true
        
        git checkout -b "${BRANCH_NAME}" origin/main
        
        git stash pop 2>/dev/null || true
    fi
fi

echo "Adding files matching pattern: ${FILES_PATTERN}"
git add ${FILES_PATTERN}

if git diff --cached --quiet; then
    echo "No staged changes to commit after adding files."
    exit 0
fi

FULL_COMMIT_MESSAGE="${COMMIT_MESSAGE} (CI run #${GITHUB_RUN_NUMBER})"
echo "Committing with message: ${FULL_COMMIT_MESSAGE}"
git commit -m "${FULL_COMMIT_MESSAGE}" || {
    echo "Commit failed or no changes to commit"
    exit 0
}

echo "Pushing branch '${BRANCH_NAME}' to origin..."
git push origin "${BRANCH_NAME}" || {
    echo "Push failed, attempting with force-with-lease..."
    git push origin "${BRANCH_NAME}" --force-with-lease
}

echo "Changes pushed to branch '${BRANCH_NAME}'"

if command -v gh &> /dev/null; then
    echo ""
    echo "Checking for existing PR..."
    export GITHUB_TOKEN="${PUSH_TOKEN}"
    
    EXISTING_PR=$(gh pr list --head "${BRANCH_NAME}" --base main --json number --jq '.[0].number' 2>/dev/null || echo "")
    
    if [[ -z "${EXISTING_PR}" ]]; then
        echo "Creating new PR to main branch..."
        
        PR_BODY="## Automated Update

**Branch**: \`${BRANCH_NAME}\`
**CI Run**: #${GITHUB_RUN_NUMBER}
**Files Changed**: \`${FILES_PATTERN}\`

### Changes
${COMMIT_MESSAGE}

---
*This PR was automatically created by the CI pipeline.*"
        
        gh pr create \
            --title "${FULL_COMMIT_MESSAGE}" \
            --body "${PR_BODY}" \
            --head "${BRANCH_NAME}" \
            --base main \
            --repo "${GITHUB_REPOSITORY}" || {
            echo "Warning: Failed to create PR. Please create it manually."
        }
        
        echo "PR created successfully"
    else
        echo "PR #${EXISTING_PR} already exists and was updated with new commit"

        gh pr comment "${EXISTING_PR}" \
            --body "Updated with new changes from CI run #${GITHUB_RUN_NUMBER}" \
            --repo "${GITHUB_REPOSITORY}" 2>/dev/null || true
    fi
else
    echo ""
    echo "⚠️  GitHub CLI not available. Please create PR manually:"
    echo "   From: ${BRANCH_NAME}"
    echo "   To: main"
    echo "   Title: ${FULL_COMMIT_MESSAGE}"
fi

if [[ "${CURRENT_BRANCH}" != "${BRANCH_NAME}" ]] && [[ -n "${CURRENT_BRANCH}" ]]; then
    echo ""
    echo "Returning to original branch: ${CURRENT_BRANCH}"
    git checkout "${CURRENT_BRANCH}" 2>/dev/null || {
        echo "Note: Could not return to ${CURRENT_BRANCH}, staying on ${BRANCH_NAME}"
    }
fi

echo "Push completed successfully"
