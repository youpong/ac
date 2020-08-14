#!/bin/bash

AS_SRC=a.s

cat >${AS_SRC}
echo "---- ${AS_SRC} ----"
cat ${AS_SRC}
echo "---- result ----"
gcc a.s
./a.out
echo $?
