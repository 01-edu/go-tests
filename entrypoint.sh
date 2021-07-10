#!/bin/sh

set -e

cp -r student piscine-go
cd piscine-go

if test "$EXAM_MODE"; then
	go mod init main 2>/dev/null
	GOSUMDB=off go get github.com/01-edu/z01@v0.1.0 2>/dev/null
fi

if test "$EXAM_RUN_ONLY" = true; then
	if command -v "${EXERCISE}_test" >/dev/null 2>&1; then
		# The exercise is a program
		go run "./$EXERCISE" "$@"
	else
		# The exercise is a function
		go run . "$@"
	fi
	exit
fi

if ! test "$EXAM_MODE"; then
	s=$(gofumpt -d .)
	if test "$s"; then
		echo 'Your Go files are not correctly formatted :'
		echo
		echo '$ gofumpt -d .'
		echo "$s"
		exit 1
	fi
fi

if ! find . -type f -name '*.go' | grep -q .; then
	echo "Missing Go file: $FILE"
	exit 1
fi

if find . -type f -name '*.go' -exec grep -qE 'print(ln)?\(' {} +; then
	echo "Your Go files cannot use print & println builtins"
	exit 1
fi

# Check restrictions
if test "$ALLOWED_FUNCTIONS" && test "$FILE"; then
	# shellcheck disable=SC2086
	rc "$FILE" $ALLOWED_FUNCTIONS
fi

if ! test -e go.mod ; then
	echo "Cannot find go.mod file, create your module with:"
	echo "    go mod init piscine"
	exit 1
fi

cp -r /go-tests ~
cd ~/go-tests

# Compile and run test
if command -v "${EXERCISE}_test" >/dev/null 2>&1; then
	# The exercise is a program
	"${EXERCISE}_test"
else
	# The exercise is a function
	go run "./tests/${EXERCISE}_test"
fi
