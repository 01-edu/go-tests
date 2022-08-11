#!/usr/bin/env bash

set -euo pipefail
IFS='
'
# clear the student folder go binary cache
rm -rf /tmp/binaries/student/

for dir in ./tests/*; do
    echo $dir
    if ! go run "$dir"; then
        echo "$dir FAILED"
    fi
done
