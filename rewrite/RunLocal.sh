#!/usr/bin/env bash

pg_exit () {
    pg_ctl -D database/data stop
    exit 
}

trap pg_exit SIGHUP SIGINT SIGTERM

pg_ctl -D database/data -l database/log.txt start
if [ $? -ne 0 ]; then
    exit 1
fi

./wichess

pg_exit
