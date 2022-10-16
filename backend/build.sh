#!/usr/bin/env bash

BACKEND_DIR=$(dirname $0)
BUILD_DIR="$BACKEND_DIR/out"

cd $BACKEND_DIR
go version
go mod tidy

mkdir -p "$BUILD_DIR"
go build -o "$BUILD_DIR/server" ./cmd/server
