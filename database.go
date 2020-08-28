package wichess

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"

	_ "github.com/lib/pq"
)

const DatabaseConfigFile = "dbconfig.json"

var Database *sql.DB

func DatabaseTransaction() *sql.Tx {
	tx, err := Database.Begin()
	if err != nil {
		Panic(err)
	}
	return tx
}

type DatabaseConfiguration struct {
	Database string
	User     string
	Password string
	Host     string
	Port     string
	SslMode  string
}

// https://github.com/pciet/wichess/issues/9
const (
	MaxIdleDatabaseConns = 10
	MaxOpenDatabaseConns = 50
)

func InitializeDatabaseConnection() {
	file, err := ioutil.ReadFile(DatabaseConfigFile)
	if err != nil {
		Panic(err)
	}
	var config DatabaseConfiguration
	err = json.Unmarshal(file, &config)
	if err != nil {
		Panic(err)
	}
	args := "dbname=" + config.Database + " host=" + config.Host + " port=" + config.Port +
		" sslmode=" + config.SslMode
	if config.User != "" {
		args += " user=" + config.User + " password=" + config.Password
	}
	Database, err = sql.Open("postgres", args)
	if err != nil {
		Panic(err)
	}
	err = Database.Ping()
	if err != nil {
		Panic(err)
	}
	Database.SetMaxIdleConns(MaxIdleDatabaseConns)
	Database.SetMaxOpenConns(MaxOpenDatabaseConns)
}
