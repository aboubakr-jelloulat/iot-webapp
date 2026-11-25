#!/usr/bin/env bash
set -euo pipefail

SERVER_IP="${1:?Server IP is required}"

# Install dependencies
apt-get update
apt-get install -y --no-install-recommends curl ca-certificates

# Install K3s server
curl -sfL https://get.k3s.io | INSTALL_K3S_EXEC="server --node-ip=${SERVER_IP}" sh -

# K3s can start serving the API before the node registers.
NODE_NAME="$(hostname | tr '[:upper:]' '[:lower:]')"
k3s kubectl wait --for=create "node/${NODE_NAME}" --timeout=120s
k3s kubectl wait --for=condition=Ready "node/${NODE_NAME}" --timeout=120s

echo "K3s installed successfully"
