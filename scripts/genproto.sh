#!/usr/bin/env bash

set -euo pipefail

# enables globstar, using `**`.
shopt -s globstar

if ! [[ "$0" =~ scripts/genproto.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

source ./scripts/lib.sh

API_ROOT="./api"

# directories containing protos to be built
function dirs {
  dirs=()
  while IFS= read -r dir; do
      dirs+=("$dir")
  done < <(find . -type f -name "*.proto" -exec dirname {} \; | xargs -n1 basename | sort -u)
  echo "${dirs[@]}"
}

dirs