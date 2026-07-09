#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
GOCTL_BIN="${GOCTL_BIN:-$(go env GOPATH)/bin/goctl}"

"${GOCTL_BIN}" api go --api "${ROOT_DIR}/apps/gateway-api/agent.api" --dir "${ROOT_DIR}/apps/gateway-api" --style gozero
"${GOCTL_BIN}" api go --api "${ROOT_DIR}/apps/chat-api/chat.api" --dir "${ROOT_DIR}/apps/chat-api" --style gozero
