// Copyright 2017 Matthew Juran
// All Rights Reserved

package main

import (
	"database/sql"
	"encoding/json"
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

type TX struct {
	*sql.Tx
}

func (db DB) Begin() TX {
	tx, err := db.DB.Begin()
	if err != nil {
		panicExit(err.Error())
	}
	return TX{
		Tx: tx,
	}
}

func (tx TX) Commit() {
	err := tx.Tx.Commit()
	if err != nil {
		panicExit(err.Error())
	}
}

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
	args := "dbname=" + config.Database + " host=" + config.Host + " port=" + config.Port + " sslmode=" + config.SslMode
	if config.User != "" {
		args += " user=" + config.User + " password=" + config.Password
	}
	database.DB, err = sql.Open("postgres", args)
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
