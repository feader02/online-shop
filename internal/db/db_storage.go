package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

type MySQLStorage struct {
	db *sqlx.DB
}

func NewStorage() *MySQLStorage {
	s := &MySQLStorage{}

	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		fmt.Println("env var DB_PORT is not found")
		dbPort = "3306"
	}

	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		fmt.Println("env var DB_HOST is not found")
		dbHost = "localhost"
	}

	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		fmt.Println("env var DB_USER is not found")
		dbUser = "root"
	}
	dbPass := os.Getenv("DB_PASSWORD")

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		fmt.Printf("env var DB_NAME is not found")
		dbName = "avion"
	}

	dbConf := mysqlDriver.Config{
		User:                 dbUser,
		Passwd:               dbPass,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", dbHost, dbPort),
		DBName:               dbName,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	s.db = DBConnect(dbConf)

	return s
}

func (ms *MySQLStorage) GetProducts(dest interface{}, tableName string, filter string, args ...interface{}) error {
	db := ms.db
	if db == nil {
		return fmt.Errorf("DB is empty")
	}

	query := "SELECT * FROM " + tableName
	if filter != "" {
		query += " WHERE " + filter
	}
	err := db.Select(dest, query, args...)
	return err
}
