#!/usr/bin/env bash

set -euo pipefail
IFS='
'

for dir in ./tests/*; do
    if ! go run "$dir"; then
        echo "$dir FAILED"
    fi
done
