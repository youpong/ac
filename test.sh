#!/bin/bash

function runtest() {
    echo "$2" | ./ac >a.s
    gcc a.s
    ./a.out
    if [[ "$1" != $? ]]; then
	echo "Error: $3" 2>&1
	exit 1
    fi
}

runtest    1    '1 '  "$LINENO"
runtest    2   ' 2;' "$LINENO"
runtest  255    '-1'  "$LINENO"
runtest    7 '3 + 4' "$LINENO"
runtest    5 '8 - 3' "$LINENO"
#runtest    5 '-6+11' "$LINENO"

rm a.s a.out

echo Okay.
