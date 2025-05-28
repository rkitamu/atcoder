#!/bin/bash
set -euo pipefail

# Make sure ~/.local/bin is in PATH
export PATH="$HOME/.local/bin:$PATH"
echo 'export PATH="$HOME/.local/bin:$PATH"' >> $HOME/.bashrc

echo "[setup.sh] Installing gopls..."
go install golang.org/x/tools/gopls@v0.15.3

echo "[setup.sh] Installing dlv (Delve debugger)..."
go install github.com/go-delve/delve/cmd/dlv@v1.22.1

echo "[setup.sh] Installing python3-venv (if missing)..."
apt-get update
apt-get install -y python3-venv

echo "[setup.sh] Setting up Python venv..."
python3 -m venv ~/.venv-pipx

echo "[setup.sh] Installing pipx inside venv..."
~/.venv-pipx/bin/pip install --no-cache-dir pipx

echo "[setup.sh] Ensuring pipx path and installing oj..."
~/.venv-pipx/bin/pipx install online-judge-tools

echo "[setup.sh] Installing nvm..."
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.3/install.sh | bash
export NVM_DIR="$HOME/.nvm"
[ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
[ -s "$NVM_DIR/bash_completion" ] && \. "$NVM_DIR/bash_completion"

echo "[setup.sh] Installing latest Node.js version using nvm..."
nvm install node
nvm use node

echo "[setup.sh] ✅ Done!"
