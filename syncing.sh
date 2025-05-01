#!/bin/bash

# Name of upstream remote
UPSTREAM="upstream"

# Branches to sync
BRANCHES=("main" "work")

echo "Fetching updates from $UPSTREAM..."
git fetch $UPSTREAM

for BRANCH in "${BRANCHES[@]}"; do
  echo ""
  echo "🔄 Syncing branch: $BRANCH"

  # If local branch exists, check it out; otherwise create from upstream
  if git show-ref --quiet refs/heads/$BRANCH; then
    git checkout $BRANCH
  else
    echo "Creating local branch '$BRANCH' from upstream/$BRANCH..."
    git checkout -b $BRANCH $UPSTREAM/$BRANCH
  fi

  # Merge changes from upstream into local branch
  echo "Merging upstream/$BRANCH into local $BRANCH..."
  git merge $UPSTREAM/$BRANCH --no-edit

  # Push to origin (your fork)
  echo "Pushing $BRANCH to your fork (origin/$BRANCH)..."
  git push origin $BRANCH
done

echo ""
echo "✅ Done syncing all branches!"
