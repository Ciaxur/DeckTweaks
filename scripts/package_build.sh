#!/usr/bin/env bash

# dev script to package up the built plugin to be yeeted to the SteamDeck.
ROOT_DIR=$(realpath "$(dirname "$0")"/..)
BACKEND_DIR="backend"
BIN_DIR="bin"
ASSETS_DIR="assets"
DIST_DIR="dist"
MISC_FILES=("main.py" "LICENSE" "README.md" "package.json" "plugin.json")
PACKAGE_FILE="plugin.zip"

# Build frontend and backend.
cd "$ROOT_DIR" || exit 1
pnpm run build || exit 1
"$BACKEND_DIR"/build.sh || exit 1

# Package backend.
mkdir -p "$BIN_DIR"
cp "$BACKEND_DIR"/out/* "$BIN_DIR"

# Package for distribution.
set -x
zip -r \
  "$PACKAGE_FILE" \
  "$DIST_DIR" \
  "$BIN_DIR" \
  "$ASSETS_DIR" \
  "${MISC_FILES[@]}"