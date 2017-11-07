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
	// https://github.com/pciet/wichess/issues/9
	database_max_idle_conns = 10
	// on macOS the fd limit is 256
	database_max_open_conns = 50
)

type DB struct {
	*sql.DB
}

var database DB

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
	}
	var config databaseConfiguration
	err = json.Unmarshal(file, &config)
	if err != nil {
		panicExit(err.Error())
	}
	database.DB, err = sql.Open("postgres", fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", config.Database, config.User, config.Password, config.Host, config.Port, config.SslMode))
	if err != nil {
		panicExit(err.Error())
	}
	err = database.Ping()
	if err != nil {
		panicExit(err.Error())
	}
	database.SetMaxIdleConns(database_max_idle_conns)
	database.SetMaxOpenConns(database_max_open_conns)
}
