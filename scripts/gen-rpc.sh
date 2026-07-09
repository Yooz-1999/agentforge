#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GOCTL_BIN="${GOCTL_BIN:-$(go env GOPATH)/bin/goctl}"

cd "${ROOT_DIR}/apps/core-rpc"
"${GOCTL_BIN}" rpc protoc core.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style gozero --module github.com/Yooz-1999/agentforge
