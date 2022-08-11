#!/usr/bin/env bash

IFS='
'

if [ -z "$1" ]
 then
   echo "No folder supplied"
   exit 1
else
    testFile="./tests/$1_test/main.go"
    if test -f "$testFile"; then

    rm -rf /tmp/binaries/student/
    go run $testFile
    echo "$1 test executed!"
    else
        echo "Error: $testFile does not exist."
    fi
fi
