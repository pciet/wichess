#!/usr/bin/env bash

# Bingo.sh does the bingo hall test, where the maxplayers benchmark is done with 50 simulated
# players to show how fast a player in a metaphorical hall of players would experience a problem.
# The only argument is the host address. The test stops if maxplayers stops abnormally, which it 
# does to indicate a problem. The bingo.log file is overwritten with the debug output during each 
# hour loop.

i=1

while true
do 
    echo hour $i
    ./maxplayers -count 50 -length 3600 -host $1 -debug > bingo.log
    if [ $? -ne 0 ]; then
        echo maxplayers error
        exit 1
    fi
    rm bingo.log
    ((i=i+1))
done
