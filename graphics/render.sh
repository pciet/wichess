#!/bin/bash
# must be called from its working directory
# TODO: indicate progress
# TODO: parallelize for number of cores, or confirm POV-Ray already does

go build generate.go
if [ $? -ne 0 ]; then
    exit 1
fi

files=""
for name in ./*.pov
do
    files="$files ${name:2}"
done

./generate $files
if [ $? -ne 0 ]; then
    exit 1
fi

mkdir ../web/img
if [ $? -ne 0 ]; then
    exit 1
fi

cp img/*.png ../web/img/
if [ $? -ne 0 ] ; then
    exit 1
fi

rm -r img
