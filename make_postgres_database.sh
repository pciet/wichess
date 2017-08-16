#!/bin/sh
DBNAME="test"
USER="test"
QUERY="make_postgres_tables.sql"
HOST="localhost"
PORT="5432"

set -x
psql -d $DBNAME -f $QUERY -h $HOST -p $PORT -U $USER
