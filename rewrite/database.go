package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

const (
	database_config_file = "dbconfig.json"

	// https://github.com/pciet/wichess/issues/9
	database_max_idle_conns = 10
	database_max_open_conns = 50
)

var Database *sql.DB

// TODO: can deferring a transaction commit to the return of an HTTP handler cause a DB lock to be held until the connection times out if there's a problem?

func DatabaseTransaction() *sql.Tx {
	tx, err := Database.Begin()
	if err != nil {
		log.Panic(err)
	}
	return tx
}

func CommitTransaction(tx *sql.Tx) {
	err := tx.Commit()
	if err != nil {
		log.Panic(err)
	}
}

type DatabaseConfiguration struct {
	Database string
	User     string
	Password string
	Host     string
	Port     string
	SslMode  string
}

func InitializeDatabaseConnection() {
	file, err := ioutil.ReadFile(database_config_file)
	if err != nil {
		log.Panic(err)
	}
	var config DatabaseConfiguration
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Panic(err)
	}
	args := "dbname=" + config.Database + " host=" + config.Host + " port=" + config.Port + " sslmode=" + config.SslMode
	if config.User != "" {
		args += " user=" + config.User + " password=" + config.Password
	}
	Database, err = sql.Open("postgres", args)
	if err != nil {
		log.Panic(err)
	}
	err = Database.Ping()
	if err != nil {
		log.Panic(err)
	}
	Database.SetMaxIdleConns(database_max_idle_conns)
	Database.SetMaxOpenConns(database_max_open_conns)
}
