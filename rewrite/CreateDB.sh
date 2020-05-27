#!/usr/bin/env bash

stop_pg () {
    pg_ctl -D database/data stop
    if [ $? -ne 0 ]; then
        exit 1
    fi
}

initdb database/data
if [ $? -ne 0 ]; then
    exit 1
fi

pg_ctl -D database/data -l database/log.txt start
if [ $? -ne 0 ]; then
    exit 1
fi

createdb test
if [ $? -ne 0 ]; then
    stop_pg
    exit 1
fi

psql -d test -f postgres_tables.sql -h localhost -p 5432
if [ $? -ne 0 ]; then
    stop_pg
    exit 1
fi

pg_ctl -D database/data stop
if [ $? -ne 0 ]; then
    exit 1
fi
