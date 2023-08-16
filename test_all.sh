#!/usr/bin/env bash

#set -euo pipefail
IFS='
'
RED="\e[31m"
GREEN="\e[32m"
WHT="\e[0m"
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

RC=0
for dir in ./tests/*; do
	exercise_name=$(echo $dir | sed 's/.\/tests\///g')
	if ! ./test_one.sh $exercise_name >/dev/null 2>&1; then
			echo -en "$RED"
			echo -e "$dir failed $WHT"
			RC=1
	else
		echo -en "$GREEN"
		echo -e "$dir success $WHT"
	fi
done
exit $RC
