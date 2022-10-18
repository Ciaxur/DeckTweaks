#!/usr/bin/env bash

BACKEND_DIR=$(dirname $0)
cd $BACKEND_DIR

go version
go mod tidy

mkdir -p out
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "out/server" ./cmd/server
