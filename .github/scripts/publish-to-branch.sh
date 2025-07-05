#!/bin/bash
# Usage: ./publish-to-branch.sh <branch> <source-dir> [--date-suffix] [--monthly] [--keep-history]

set -euo pipefail

# Function to log with timestamp
log() {
    echo "[$(date -u '+%Y-%m-%d %H:%M:%S UTC')] $*"
}

# Function to restore original branch
restore_original_branch() {
    if [[ -n "${ORIGINAL_BRANCH:-}" ]] && [[ -n "${GITHUB_WORKSPACE:-}" ]]; then
        log "Restoring original branch: $ORIGINAL_BRANCH"
        cd "$GITHUB_WORKSPACE" 2>/dev/null || true
        git checkout "$ORIGINAL_BRANCH" 2>/dev/null || true
    fi
}

# Set up cleanup trap
trap restore_original_branch EXIT

if [[ -n "${GITHUB_WORKSPACE:-}" ]]; then
    cd "$GITHUB_WORKSPACE" || {
        echo "Error: Could not change to workspace directory: $GITHUB_WORKSPACE"
        exit 1
    }
fi

# Parse arguments
if [[ $# -lt 2 ]]; then
    echo "Usage: $0 <branch> <source-dir> [--date-suffix] [--monthly] [--keep-history]"
    echo "Options:"
    echo "  --date-suffix    Append current month-day to branch name"
    echo "  --monthly        Organize files in monthly subdirectories"
    echo "  --keep-history   Keep git history (default: squash to single commit)"
    exit 1
fi

BRANCH="$1"
SOURCE_DIR="$2"
DATE_SUFFIX=""
MONTHLY=false
KEEP_HISTORY=false

if [[ -z "$BRANCH" || -z "$SOURCE_DIR" ]]; then
  echo "Usage: $0 <branch> <source-dir> [--date-suffix] [--monthly] [--keep-history]"
  exit 1
fi

shift 2
while [[ $# -gt 0 ]]; do
    case $1 in
        --date-suffix)
            DATE_SUFFIX="_$(date -u +%m-%d)"
            ;;
        --monthly)
            MONTHLY=true
            ;;
        --keep-history)
            KEEP_HISTORY=true
            ;;
        *)
            echo "Error: Unknown option '$1'"
            echo "Usage: $0 <branch> <source-dir> [--date-suffix] [--monthly] [--keep-history]"
            exit 1
            ;;
    esac
    shift
done

FILE_COUNT=$(find "$SOURCE_DIR" -type f 2>/dev/null | wc -l)
if [[ "$FILE_COUNT" -eq 0 ]]; then
    log "No files found in $SOURCE_DIR, skipping publish to $BRANCH"
    exit 0
fi

log "Found $FILE_COUNT files in $SOURCE_DIR, proceeding with publish to $BRANCH"

if [[ -n "${GITHUB_WORKSPACE:-}" ]]; then
    cd "$GITHUB_WORKSPACE" || {
        log "Error: Could not change to workspace directory: $GITHUB_WORKSPACE"
        exit 1
    }
fi

# Store current branch for restoration
ORIGINAL_BRANCH=$(git rev-parse --abbrev-ref HEAD 2>/dev/null || echo "")
log "Current branch: ${ORIGINAL_BRANCH:-unknown}"

# Check for any uncommitted changes and warn if found
if [[ -n "$(git status --porcelain)" ]]; then
    log "Warning: Uncommitted changes detected in working directory"
    git status --porcelain
    log "These changes may cause issues during branch checkout"
fi

# Fetch latest changes
log "Fetching latest changes from origin"
git fetch origin --prune

# Check if target branch exists
BRANCH_EXISTS=false
if git ls-remote --exit-code --heads origin "$BRANCH" >/dev/null 2>&1; then
    BRANCH_EXISTS=true
    log "Branch '$BRANCH' exists remotely"
else
    log "Branch '$BRANCH' does not exist remotely"
fi

# Switch to target branch
if $BRANCH_EXISTS; then
    log "Checking out existing branch '$BRANCH'"
    
    # Try to checkout directly first
    if ! git checkout "$BRANCH" 2>/dev/null; then
        log "Direct checkout failed, trying to create from remote"
        if ! git checkout -b "$BRANCH" "origin/$BRANCH" 2>/dev/null; then
            log "Error: Could not checkout branch '$BRANCH'"
            log "This may be due to uncommitted changes in the working directory"
            exit 1
        fi
    fi
    
    # Pull latest changes
    git pull origin "$BRANCH" --ff-only 2>/dev/null || {
        log "Fast-forward merge failed, trying regular merge"
        git merge "origin/$BRANCH" --no-edit
    }
else
    log "Creating new orphan branch '$BRANCH'"
    git checkout --orphan "$BRANCH"
    git rm -rf . 2>/dev/null || true
    
    # Initialize branch with appropriate README
    case "$BRANCH" in
    summaries*)
        cat > README.md << EOF
# DNS Toolkit Summary Files Archive

This branch contains JSON summary files organized by month (01–12), with 1 year history.

Generated on: $(date -u '+%Y-%m-%d %H:%M UTC')
EOF
            ;;
        output*)
            cat > README.md << EOF
# DNS Toolkit Output Files

This branch contains processed output files from the DNS toolkit.

Generated on: $(date -u '+%Y-%m-%d %H:%M UTC')
EOF
            ;;
        *)
            cat > README.md << EOF
# $BRANCH Branch

Automated branch created by DNS Toolkit workflow.

Generated on: $(date -u '+%Y-%m-%d %H:%M UTC')
EOF
            ;;
    esac
    
    git add README.md
    git commit -m "Initialize $BRANCH branch"
    git push -u origin "$BRANCH"
    log "Initialized new branch '$BRANCH' with README"
fi

if $MONTHLY; then
    MONTH_DIR=$(date -u +%m)
    TARGET_DIR="$MONTH_DIR"
    log "Using monthly organization, target directory: $TARGET_DIR"
    
    # Clean only the current month directory to ensure only current version
    log "Cleaning current month directory ($MONTH_DIR) to maintain only current version"
    rm -rf "$TARGET_DIR" 2>/dev/null || true
    mkdir -p "$TARGET_DIR"
    
    # Clean any loose files in root (preserve other month directories and README.md)
    find . -maxdepth 1 -type f ! -name ".git*" ! -name "README.md" -delete 2>/dev/null || true
else
    TARGET_DIR="."
    log "Using root directory for files"
    
    # Clean root directory but preserve existing README.md if it exists in source
    log "Cleaning root directory content"
    if [[ -f "$SOURCE_DIR/README.md" ]]; then
        find . -maxdepth 1 -type f ! -name ".git*" -delete 2>/dev/null || true
        find . -maxdepth 1 -type d ! -name ".git" ! -name "." -exec rm -rf {} + 2>/dev/null || true
    else
        find . -maxdepth 1 -type f ! -name ".git*" ! -name "README.md" -delete 2>/dev/null || true
        find . -maxdepth 1 -type d ! -name ".git" ! -name "." -exec rm -rf {} + 2>/dev/null || true
    fi
fi

log "Copying files to target location"
if [[ "$TARGET_DIR" == "." ]]; then
    find "$SOURCE_DIR" -maxdepth 1 -type f -exec cp {} . \; 2>/dev/null || true
    find "$SOURCE_DIR" -maxdepth 1 -type d ! -path "$SOURCE_DIR" -exec cp -r {} . \; 2>/dev/null || true
else
    # For monthly organization, copy files to the target directory
    cp -r "$SOURCE_DIR"/* "$TARGET_DIR/" 2>/dev/null || true
    
    # Special handling for summaries branch: move README.md to root
    if [[ "$BRANCH" == "summaries"* ]] && [[ -f "$TARGET_DIR/README.md" ]]; then
        log "Moving summaries README.md to root directory"
        mv "$TARGET_DIR/README.md" "./README.md"
    fi
fi

# Ensure we're in a git repository before checking status
if [[ ! -d ".git" ]]; then
    log "Error: Not in a git repository"
    exit 1
fi

if [[ -n "$(git status --porcelain)" ]]; then
    log "Changes detected, committing and pushing"
    git add .
    
    COMMIT_MSG="Update $BRANCH files - $(date -u '+%Y-%m-%d %H:%M UTC')"
    if $MONTHLY; then
        COMMIT_MSG="$COMMIT_MSG (month: $(date -u +%m))"
    fi
    
    # Option to squash history to maintain only single version in git history
    if ! $KEEP_HISTORY && $BRANCH_EXISTS; then
        log "Squashing git history to maintain single version"
        git reset --soft $(git rev-list --max-parents=0 HEAD | tail -1)
        git commit --amend -m "$COMMIT_MSG"
        git push origin "$BRANCH" --force-with-lease
    else
        git commit -m "$COMMIT_MSG"
        git push origin "$BRANCH"
    fi
    
    log "Successfully published $FILE_COUNT files to $BRANCH branch"
else
    log "No changes detected, nothing to commit to $BRANCH branch"
fi

log "Publish operation completed successfully"