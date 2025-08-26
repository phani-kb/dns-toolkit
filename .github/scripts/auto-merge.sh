#!/usr/bin/env bash
set -euo pipefail

# Usage: ./safe-auto-merge.sh <files-pattern> <commit-message> [branch-name] [description]

if [[ $# -lt 3 ]]; then
  echo "Usage: $0 <files-pattern> <commit-message> [branch-name] [description]" >&2
  exit 1
fi

FILES_PATTERN="$1"
COMMIT_MESSAGE="$2"
BRANCH_NAME="${3:-}"
PR_BODY="${4:-}"

echo "Files: $FILES_PATTERN"
echo "Message: $COMMIT_MESSAGE"
echo "Branch: $BRANCH_NAME"

if [[ -z "${PUSH_TOKEN:-}" ]]; then
    echo "No PUSH_TOKEN available, changes only committed locally"
    exit 0
fi
export GH_TOKEN="${PUSH_TOKEN}"


if [[ -z "${GITHUB_REPOSITORY:-}" ]]; then
    echo "GITHUB_REPOSITORY not set, skipping auto-merge" >&2 && exit 1
fi

timestamp=$(date -u '+%Y-%m-%d %H:%M:%S UTC')

git config user.name "GitHub Actions Bot"
git config user.email "actions@github.com"

ORIG_BRANCH="$(git rev-parse --abbrev-ref HEAD)"
STASHED_ORIG=0
STASHED_TARGET=0
ORIG_STASH_REF=""
TARGET_STASH_REF=""

cleanup() {
  set +e
  if [[ $STASHED_TARGET -eq 1 ]]; then
    echo "Restoring stashed changes on target branch '$BRANCH_NAME'..."
    git checkout "$BRANCH_NAME" >/dev/null 2>&1 || true
    if git stash list | grep -q "$TARGET_STASH_REF"; then
      git stash pop "$TARGET_STASH_REF" || {
        echo "Conflicts while popping target branch stash. Stash kept." >&2
      }
    fi
  fi

  if [[ "$ORIG_BRANCH" != "$BRANCH_NAME" ]]; then
    git checkout "$ORIG_BRANCH" >/dev/null 2>&1 || true
  fi
  if [[ $STASHED_ORIG -eq 1 ]]; then
    echo "Restoring stashed changes on original branch '$ORIG_BRANCH'..."
    if git stash list | grep -q "$ORIG_STASH_REF"; then
      git stash pop "$ORIG_STASH_REF" || {
        echo "Conflicts while popping original branch stash. Stash kept." >&2
      }
    fi
  fi
}
trap cleanup EXIT

if [[ "$ORIG_BRANCH" != "$BRANCH_NAME" ]]; then
  if ! git diff --quiet || ! git diff --cached --quiet || [[ -n "$(git ls-files --others --exclude-standard)" ]]; then
    TEMP_COPY_DIR="/tmp/ci-orig-files-$$"
    mkdir -p "$TEMP_COPY_DIR"
    shopt -s nullglob
    temp_files=($FILES_PATTERN)
    shopt -u nullglob
    if [[ ${#temp_files[@]} -gt 0 ]]; then
      echo "Saving ${#temp_files[@]} files to temporary dir before branch switch: ${temp_files[*]}"
      for f in "${temp_files[@]}"; do
        if [[ -f "$f" ]]; then
          mkdir -p "$TEMP_COPY_DIR/$(dirname "$f")"
          cp -- "$f" "$TEMP_COPY_DIR/$f" || true
        fi
      done
    else
      echo "No matching files to save before stash for pattern: $FILES_PATTERN"
    fi

    ORIG_STASH_REF="stash@{0}"
    echo "Stashing changes on '$ORIG_BRANCH' before switching..."
    git stash push -u -m "auto-stash: $ORIG_BRANCH -> $BRANCH_NAME @ $timestamp"
    STASHED_ORIG=1
    ORIG_STASH_REF="$(git stash list | head -n1 | cut -d: -f1)"
  fi

  git fetch origin "+refs/heads/*:refs/remotes/origin/*" >/dev/null 2>&1 || true
  if git show-ref --verify --quiet "refs/heads/$BRANCH_NAME"; then
    git checkout "$BRANCH_NAME"
  elif git ls-remote --exit-code --heads origin "$BRANCH_NAME" >/dev/null 2>&1; then
    git checkout -t "origin/$BRANCH_NAME"
  else
    DEFAULT_BASE="$(gh repo view "$GITHUB_REPOSITORY" --json defaultBranchRef -q '.defaultBranchRef.name' 2>/dev/null || echo main)"
    git fetch origin "$DEFAULT_BASE:$DEFAULT_BASE" >/dev/null 2>&1 || true
    git checkout -b "$BRANCH_NAME" "$DEFAULT_BASE"
  fi
  if [[ -n "${TEMP_COPY_DIR:-}" ]] && [[ -d "$TEMP_COPY_DIR" ]]; then
    echo "Restoring files from temporary copy directory onto '$BRANCH_NAME'..."
    find "$TEMP_COPY_DIR" -type f | while read -r src; do
      dest="${src#$TEMP_COPY_DIR/}"
      mkdir -p "$(dirname "$dest")"
      cp -- "$src" "$dest" || true
      echo "  restored $dest"
    done
    rm -rf "$TEMP_COPY_DIR" || true
  fi
else
  if ! git diff --quiet || ! git diff --cached --quiet || [[ -n "$(git ls-files --others --exclude-standard)" ]]; then
    echo "Stashing uncommitted changes on target branch '$BRANCH_NAME'..."
    git stash push -u -m "auto-stash: $BRANCH_NAME (pre-commit) @ $timestamp"
    STASHED_TARGET=1
    TARGET_STASH_REF="$(git stash list | head -n1 | cut -d: -f1)"
  fi
fi

if [[ "$FILES_PATTERN" == "." || "$FILES_PATTERN" == "./" ]]; then
  git add -A
else
  shopt -s nullglob  # to handle no matches gracefully
  files_array=($FILES_PATTERN)
  shopt -u nullglob
  
  if [[ ${#files_array[@]} -gt 0 ]]; then
    echo "Found ${#files_array[@]} files matching pattern: $FILES_PATTERN"
    echo "Files to add: ${files_array[*]}"
    echo "Git status before adding:"
    git status --porcelain -- "${files_array[@]}"
    echo "Working directory changes for these files:"
    git diff -- "${files_array[@]}" | head -20
    git add "${files_array[@]}" || {
      echo "Failed to add matched files"
      exit 1
    }
    echo "Git status after adding:"
    git status --porcelain -- "${files_array[@]}"
    echo "Staged changes for these files:"
    git diff --cached -- "${files_array[@]}" | head -20
  else
    echo "No files found matching pattern: $FILES_PATTERN"
    if [[ -f "$FILES_PATTERN" ]]; then
      git add "$FILES_PATTERN"
      echo "Added literal file: $FILES_PATTERN"
    else
      echo "No files to add for pattern: $FILES_PATTERN"
      exit 0
    fi
  fi
fi

echo "Checking for staged changes..."
echo "Git diff --cached status:"
git diff --cached --name-status
echo "Git status --porcelain:"
git status --porcelain

if git diff --cached --quiet; then
  echo "No staged changes for pattern '$FILES_PATTERN'. Nothing to do."
  exit 0
fi

git commit -m "$COMMIT_MESSAGE"

remote_url="$(git remote get-url origin)"
if [[ "$remote_url" =~ ^http ]]; then
  authed_url="$remote_url"
  if [[ "$remote_url" =~ ^https://([^@]+)@github\.com/(.*)$ ]]; then
    authed_url="$remote_url"
  else
    authed_url="${remote_url/https:\/\//https:\/\/x-access-token:${PUSH_TOKEN}@}"
  fi
else
  if [[ "$remote_url" =~ git@github\.com:(.*)\.git$ ]]; then
    authed_url="https://x-access-token:${PUSH_TOKEN}@github.com/${BASH_REMATCH[1]}.git"
  else
    echo "Unrecognized origin URL format: $remote_url" >&2
    exit 1
  fi
fi

echo "Pushing branch '$BRANCH_NAME'..."
git push "$authed_url" "HEAD:refs/heads/$BRANCH_NAME" --force-with-lease >/dev/null

BASE_BRANCH="$(gh repo view "$GITHUB_REPOSITORY" --json defaultBranchRef -q '.defaultBranchRef.name' 2>/dev/null || echo main)"

PR_NUMBER="$(gh pr list --base "$BASE_BRANCH" --head "$BRANCH_NAME" --json number -q '.[0].number' 2>/dev/null || true)"
if [[ -z "${PR_NUMBER:-}" ]]; then
  echo "Creating PR from '$BRANCH_NAME' -> '$BASE_BRANCH'..."
  PR_TITLE="$(printf "%s" "$COMMIT_MESSAGE" | head -n1)"
  if [[ -n "$PR_BODY" ]]; then
    PR_NUMBER="$(gh pr create --base "$BASE_BRANCH" --head "$BRANCH_NAME" --title "$PR_TITLE" --body "$PR_BODY" --json number -q '.number')"
  else
    PR_NUMBER="$(gh pr create --base "$BASE_BRANCH" --head "$BRANCH_NAME" --title "$PR_TITLE" --fill --json number -q '.number')"
  fi
else
  echo "PR already exists: #$PR_NUMBER"
  if [[ -n "$PR_BODY" ]]; then
    gh pr edit "$PR_NUMBER" --body "$PR_BODY" >/dev/null || true
  fi
fi

echo "Enabling auto-merge (squash) for PR #$PR_NUMBER..."
if ! gh pr merge "$PR_NUMBER" --auto --squash --delete-branch >/dev/null 2>&1; then
  echo "Auto-merge queued (will merge when checks pass) or not available."
fi

echo "Done. PR #$PR_NUMBER is open against '$BASE_BRANCH' and set to auto-merge."
