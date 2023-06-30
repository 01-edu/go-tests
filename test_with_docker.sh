#!/usr/bin/env bash

IFS='
'
cd -P "$(dirname "$0")"

EXERCISE_NAME=$1
EXERCISE_FILE=$2
ALLOWED_FUNCTIONS="$(echo "${@:3}")"

echo $EXERCISE_NAME $EXERCISE_FILE $ALLOWED_FUNCTIONS

docker run --read-only \
			--network none \
			--memory 500M \
			--cpus 2.0 \
			--user 1000:1000 \
            --env FILE="$EXERCISE_FILE" \
			--env EXERCISE="$EXERCISE_NAME" \
            --env ALLOWED_FUNCTIONS="$ALLOWED_FUNCTIONS" \
			--env USERNAME=msessa \
			--env HOME=/jail \
			--env TMPDIR=/jail \
			--workdir /jail \
			--tmpfs /jail:size=100M,noatime,exec,nodev,nosuid,uid=1000,gid=1000,nr_inodes=5k,mode=1700 \
			--volume "$(pwd)"/../piscine-go:/jail/student:ro \
			-t go_tests
