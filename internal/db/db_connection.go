package db

import (
	"fmt"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jmoiron/sqlx"
	"os"
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

func setUpToDateDB(db *sqlx.DB) error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		return fmt.Errorf("cannot obtain driver: %s", err)
	}

	path := os.Getenv("MIGRATIONS_PATH")
	m, err := migrate.NewWithDatabaseInstance(
		path,
		"products", driver)
	if err != nil {
		return fmt.Errorf("cannot migrate: %s", err)
	}
	return m.Up()
}
