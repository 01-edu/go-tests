#!/bin/sh

set -e

# ! support both variables CODE_EDITOR_RUN_ONLY and EXAM_RUN_ONLY
CODE_EDITOR_RUN_ONLY="${CODE_EDITOR_RUN_ONLY:-$EXAM_RUN_ONLY}"
# ! support both variables CODE_EDITOR_MODE and EXAM_MODE
CODE_EDITOR_MODE="${CODE_EDITOR_MODE:-$EXAM_MODE}"

cp -r student piscine-go
cd piscine-go

if test "$CODE_EDITOR_MODE"; then
	go mod init piscine 2>/dev/null
	GOSUMDB=off go get github.com/01-edu/z01@v0.1.0 2>/dev/null
fi

if test "$CODE_EDITOR_RUN_ONLY" = true; then
	# ! to support both the old and the new version of the runner we
	# ! need to check the files in the code editor
	# if the files in the editor contain the "main.go" we are running a program
	if echo "$EDITOR_FILES" | tr ',' '\n' | grep -q '/main.go' || command -v "${EXERCISE}_test" >/dev/null 2>&1; then
		go run "./$EXERCISE" "$@"
	else
		# The exercise is a function
		go run . "$@"
	fi
	exit
fi

if ! test "$CODE_EDITOR_MODE"; then
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

if find . -type f -name '*.go' -exec grep -qE '\bprint(ln)?\s*\(' {} +; then
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
