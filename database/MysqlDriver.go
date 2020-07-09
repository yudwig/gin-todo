package database

import (
	"database/sql"
	"fmt"
)

type MysqlConfig struct {
	user string
	password string
	host string
	port int
	database string
}

func NewDefaultMysqlConfig() *MysqlConfig {
	return &MysqlConfig {
		user: "root",
		password: "",
		host: "localhost",
		port: 3306,
		database: "gin_todo",
	}
}

type MysqlDriver struct {
	config *MysqlConfig
	db *sql.DB
}

func NewMysqlDriver(conf *MysqlConfig) (*MysqlDriver, error) {
	src := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.user, conf.password, conf.host, conf.port, conf.database)
	db, err := sql.Open("mysql", src)
	if err != nil {
		return nil, err
	}
	driver := &MysqlDriver {
		config: conf,
		db: db,
	}
	return driver, nil
}

func NewDefaultMysqlDriver() (*MysqlDriver, error) {
	conf := NewDefaultMysqlConfig()
	return NewMysqlDriver(conf)
}

