#!/usr/bin/env bash

IFS='
'

RC=rc

if [ -z "$1" ]; then
    echo "No folder supplied"
    exit 1
else
    testDir="tests/$1"
	testName=$1
    if test -d "$testDir"; then
        # Fix for mac and linux temp directories
        if [ -z "$TMPDIR" ]; then
            # Linux
            rm -rf /tmp/binaries/student/
        else
            # Mac
            rm -rf $TMPDIR/binaries/student/
        fi
        # End of fix

		# Checking if our solution respects the same restrictions as student's
		if ! $RC -config "$testDir/rc.yml" "$testDir/$testName.go"; then
			echo "Our solution doesn't respect restrictions"
			exit 1
		fi

		# Check if solution has tests	
		if ! test -f "./tests/$1/$1_test.go" ; then
			echo "No tests for our solution"
			exit 1
		fi	
		
		echo "Testing our solution:"
		if ! go test "./$testDir"; then
			exit 1
		fi
		
		# Test the student solution
		if ! go run "./$testDir"; then
			echo "$1 tests failed!"
			exit 1
		else
			echo "$1 test executed successfully!"
			exit 0
		fi
    else
        echo "Error: $testDir does not exist."
    fi
fi
