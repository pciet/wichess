// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

const (
	database_config_file = "dbconfig.json"

	database_player_table = "players"

	database_player_table_name_key  = "name"
	database_player_table_crypt_key = "crypt"
)

var database *sql.DB

type databaseConfiguration struct {
	Database string
	User     string
	Password string
	Host     string
	Port     string
	SslMode  string
}

func initializeDatabaseConnection() {
	file, err := ioutil.ReadFile(database_config_file)
	if err != nil {
		panicExit(err.Error())
		return
	}
	var config databaseConfiguration
	err = json.Unmarshal(file, &config)
	if err != nil {
		panicExit(err.Error())
		return
	}
	database, err = sql.Open("postgres", fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", config.Database, config.User, config.Password, config.Host, config.Port, config.SslMode))
	if err != nil {
		panicExit(err.Error())
		return
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
