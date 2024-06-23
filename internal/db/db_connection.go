package db

import (
	"fmt"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func DBConnect(dbConf mysqlDriver.Config) *sqlx.DB {
	db, err := sqlx.Open("mysql", dbConf.FormatDSN())
	if err != nil {
		panic(err.Error())
	}

	return db
}

func (ms *MySQLStorage) Close() error {
	if ms.db == nil {
		return fmt.Errorf("DB is empty")
	}
	return ms.db.Close()
}
