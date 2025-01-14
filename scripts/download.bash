#!/usr/bin/env bash

# Determine OS and architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Check if jq is installed
if ! command -v jq &> /dev/null; then
  echo "jq is required but not installed.  Please install it."
  echo ""
  case "$OS" in
    darwin)
      echo "On macOS, use Homebrew:  brew install jq"
      ;;
    linux)
      echo "On most Linux distributions, use your package manager (e.g., apt, yum, dnf):"
      echo "   - Debian/Ubuntu: sudo apt-get install jq"
      echo "   - Fedora/CentOS/RHEL: sudo dnf install jq or sudo yum install jq"
      ;;
    windows)
      echo "On Windows, download jq from https://stedolan.github.io/jq/download/"
      ;;
    *)
      echo "Installation instructions are not available for your OS."
      ;;
  esac
  read -p "Press Enter after installing jq to continue..."
  if ! command -v jq &> /dev/null; then
    echo "jq still not found. Exiting."
    exit 1
  fi
fi

# Construct the artifact filename pattern
case "$OS" in
  darwin)
    OS="darwin"
    ;;
  linux)
    OS="linux"
    ;;
  windows)
    OS="windows"
    ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac

case "$ARCH" in
  arm64)
    ARCH="arm64"
    ;;
  amd64)
    ARCH="amd64"
    ;;
  x86_64)
    ARCH="amd64"  # Treat x86_64 as amd64
    ;;
  arm)
    ARCH="arm"
    ;;
  386)
    ARCH="386"
    ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

ARTIFACT_PATTERN="parrot-${OS}_${ARCH}.tar.gz"


# Get the latest release tag
LATEST_RELEASE=$(curl -sL "https://api.github.com/repos/ARTM2000/parrot/releases/latest" | jq -r '.tag_name')

# Construct the download URL
DOWNLOAD_URL="https://github.com/ARTM2000/parrot/releases/download/${LATEST_RELEASE}/${ARTIFACT_PATTERN}"

# Download the artifact
# Download the artifact
echo "Downloading: $DOWNLOAD_URL"
curl -L "$DOWNLOAD_URL" -o parrot.tar.gz

# Check if download was successful
if [ $? -ne 0 ]; then
  echo "Download failed!"
  exit 1
fi

# Extract the artifact (assuming it's a tar.gz)
echo "Extracting parrot.tar.gz"
tar -xzvf parrot.tar.gz

# Remove the downloaded archive
rm parrot.tar.gz

# shellcheck disable=SC2002
EXPECTED_CHECKSUM=$(cat "parrot-${OS}_${ARCH}.sha256" | awk '{print $1}')
ACTUAL_CHECKSUM=$(shasum -a 256 "parrot-${OS}_${ARCH}" | awk '{print $1}')
if [ "$EXPECTED_CHECKSUM" != "$ACTUAL_CHECKSUM" ]; then
  echo "Checksum mismatch! Download corrupted."
  exit 1
fi

# Rename the binary file
mv parrot-"${OS}_${ARCH}" parrot
rm parrot-"${OS}_${ARCH}".sha256

echo "Done!"