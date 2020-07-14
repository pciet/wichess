package main

import "database/sql"

func SQLExecRow(tx *sql.Tx, query string, args ...interface{}) {
	r, err := tx.Exec(query, args...)
	if err != nil {
		Panic(err)
	}

	c, err := r.RowsAffected()
	if err != nil {
		Panic(err)
	}

	if c != 1 {
		Panic(c, "rows affected by", query, args)
	}
}
