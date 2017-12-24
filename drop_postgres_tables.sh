#!/bin/sh
DBNAME="test"
USER="test"
HOST="localhost"
PORT="5432"

psql -h $HOST -p $PORT -c "DROP TABLE players; DROP TABLE sessions; DROP TABLE friends; DROP TABLE games; DROP TABLE pieces;" $DBNAME $USER
