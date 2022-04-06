#!/usr/bin/env bash

set -euo pipefail
IFS='
'

for dir in ./tests/*; do
    echo $dir
    if ! go run "$dir"; then
        echo "$dir FAILED"
    fi
done
