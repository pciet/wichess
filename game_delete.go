package wichess

import "database/sql"

func DeleteGame(tx *sql.Tx, id GameIdentifier) { SQLExecRow(tx, GamesDelete, id) }
