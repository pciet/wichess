// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	database_user = "test"
	database_name = "test"

	database_player_table = "players"

	database_player_table_name_key  = "name"
	database_player_table_crypt_key = "crypt"
)

var database *sql.DB

func init() {
	var err error
	database, err = sql.Open("postgres", fmt.Sprintf("user=%v dbname=%v sslmode=disable", database_user, database_name))
	if err != nil {
		panic(err.Error())
	}
}

func playerFromDatabase(name string) (bool, string) {
	rows, err := database.Query(fmt.Sprintf("SELECT * FROM %v WHERE %v=$1", database_player_table, database_player_table_name_key), name)
	if err != nil {
		panicExit(err.Error())
		return false, ""
	}
	defer rows.Close()
	exists := rows.Next()
	if exists == false {
		return false, ""
	}
	var n, c string
	err = rows.Scan(&n, &c)
	if err != nil {
		panicExit(err.Error())
	}
	if rows.Next() {
		panicExit(fmt.Sprintf("duplicate database entries for %v", name))
		return false, ""
	}
	return true, c

}

func newPlayerInDatabase(name, crypt string) {
	_, err := database.Exec(fmt.Sprintf("INSERT INTO %v (%v, %v) VALUES ($1, $2)", database_player_table, database_player_table_name_key, database_player_table_crypt_key), name, crypt)
	if err != nil {
		panicExit(err.Error())
	}
	return
}
