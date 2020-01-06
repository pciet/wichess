#!/bin/bash

#must be called from in the same working directory

stop_pg () {
    pg_ctl -D database/data stop
    if [ $? -ne 0 ]; then
        exit 1
    fi
}

# if these are installed but which doesn't find them then check your PATH environment variable

which -s go
if [ $? -ne 0 ]; then
    echo "Go programming language tools are required, install from golang.org"
    exit 1
fi

which -s brew
if [ $? -ne 0 ]; then
    echo "Homebrew is required, install from http://brew.sh"
    exit 1
fi

go build
if [ $? -ne 0 ]; then
    exit 1
fi

brew install postgresql
if [ $? -ne 0 ]; then
    exit 1
fi

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

stop_pg

brew install povray
if [ $? -ne 0 ]; then
    exit 1
fi

cd graphics

./render.sh
if [ $? -ne 0 ]; then
    exit 1
fi

cd ..
