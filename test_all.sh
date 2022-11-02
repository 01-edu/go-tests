#!/usr/bin/env bash

#set -euo pipefail
IFS='
'
# clear the student folder go binary cache
# Fix for mac and linux temp directories
if [ -z "$TMPDIR" ]; then
    # Linux
    rm -rf /tmp/binaries/student/
else
    # Mac
    rm -rf $TMPDIR/binaries/student/
fi
# End of fix

for dir in ./tests/*; do
    echo "$dir executed"
    if ! go run "$dir"; then
        echo "$dir FAILED"
    fi
done
