#!/bin/sh
set -e

# settings
owner="carbonetes"
repo="ci"
binary_name="carbonetes-ci"
version=""
github_url="https://github.com"
executable_folder="./bin"
format="tar.gz"

usage() (
  this=$1
  cat <<EOF
$this: download go binaries for ${owner}/${repo}

Usage: $this [-d dir] [-v tag]
  -d  Installation directory (default: ./bin)
  -v  Specific release version (default: latest)
EOF
  exit 2
)

get_arch() {
  a=$(uname -m)
  case ${a} in
    x86_64 | amd64) echo "amd64" ;;
    aarch64 | arm64 | arm) echo "arm64" ;;
    ppc64le) echo "ppc64le" ;;
    *) echo "unsupported-arch" ;;
  esac
}

get_os() {
  os=$(uname -s | tr '[:upper:]' '[:lower:]')
  case "$os" in
    darwin) echo "darwin" ;;
    linux) echo "linux" ;;
    cygwin* | mingw* | msys*) echo "windows" ;;
    *) echo "unsupported-os" ;;
  esac
}

check_shasum() {
  command -v shasum >/dev/null || {
    echo "Error: shasum is not installed. Please install shasum." >&2
    exit 1
  }
}

get_latest_release() {
  curl -s "https://api.github.com/repos/${owner}/${repo}/releases/latest" |
    grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
}

verify_sha256() {
  asset_file="$1"
  checksum_file="$2"
  base=$(basename "$asset_file")
  expected=$(grep "$base" "$checksum_file" | cut -d ' ' -f 1)
  actual=$(shasum -a 256 "$asset_file" | cut -d ' ' -f 1)

  if [ "$expected" != "$actual" ]; then
    echo "Checksum mismatch for $base!"
    echo "Expected: $expected"
    echo "Actual:   $actual"
    exit 1
  fi
}

extract() {
  file="$1"
  case "$file" in
    *.tar.gz) tar -xzf "$file" ;;
    *.zip) unzip -q "$file" ;;
    *) echo "Unknown archive format: $file" && return 1 ;;
  esac
}

install_binary() {
  archive="$1"
  dest="$2"
  bin_name="$3"

  dir=$(dirname "$archive")
  (cd "$dir" && extract "$archive")

  mkdir -p "$dest"
  install "$dir/$bin_name" "$dest"
  chmod +x "$dest/$bin_name"
}

install_carbonetes_ci() {
  check_shasum

  while getopts "v:d:" arg; do
    case "$arg" in
      v) version="$OPTARG" ;;
      d) executable_folder="$OPTARG" ;;
    esac
  done
  shift $((OPTIND - 1))

  tmp_dir=$(mktemp -d -t carbonetes-ci-XXXXXXXXXX)
  trap 'rm -rf "$tmp_dir"' EXIT

  os=$(get_os)
  arch=$(get_arch)
  if [ "$os" = "unsupported-os" ] || [ "$arch" = "unsupported-arch" ]; then
    echo "Unsupported OS or architecture: os=$os arch=$arch"
    exit 1
  fi

  if [ -z "$version" ]; then
    version=$(get_latest_release)
  fi

  # trim "v" from version
  version_tag="${version#v}"

  ext="tar.gz"
  [ "$os" = "windows" ] && ext="zip"

  archive_name="${binary_name}_${version_tag}_${os}_${arch}.${ext}"
  checksum_name="${binary_name}_${version_tag}_checksums.txt"
  download_url="${github_url}/${owner}/${repo}/releases/download/${version}/${archive_name}"
  checksum_url="${github_url}/${owner}/${repo}/releases/download/${version}/${checksum_name}"

  archive_path="${tmp_dir}/${archive_name}"
  checksum_path="${tmp_dir}/${checksum_name}"

  echo "[1/4] Downloading binary from $download_url"
  curl -sSfL -o "$archive_path" "$download_url"

  echo "[2/4] Downloading checksum from $checksum_url"
  curl -sSfL -o "$checksum_path" "$checksum_url"

  echo "[3/4] Verifying checksum..."
  verify_sha256 "$archive_path" "$checksum_path"

  echo "[4/4] Installing to ${executable_folder}"
  install_binary "$archive_path" "$executable_folder" "$binary_name"

  echo "[>] ${binary_name} installed successfully to ${executable_folder}/${binary_name}"
  echo "Add to PATH if needed:"
  echo "  export PATH=${executable_folder}:\$PATH"
  echo "Run '${binary_name} --help' to get started."
}

install_carbonetes_ci "$@"
