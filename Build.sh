#!/usr/bin/env bash

# Build.sh makes the wichess host program and mem folder. This script and all others for wichess
# must be run in the working directory they reside in. To completely build the wichess host program
# the piece images must be separately created by Render.sh in the img folder.

if ! command -v go &> /dev/null
then
    echo "go not found"
    exit 1
fi

if [ ! -d mem ]
then
    mkdir mem
fi

go build -o wichess github.com/pciet/wichess/cmd
