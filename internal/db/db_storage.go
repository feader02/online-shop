package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strings"
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

	fmt.Println(dbConf.FormatDSN())

	return s
}

func (ms *MySQLStorage) Get(dest interface{}, tableName string, filter string, args ...interface{}) error {
	db := ms.db
	if db == nil {
		return fmt.Errorf("DB is empty")
	}

	query := "SELECT * FROM " + tableName
	if filter != "" {
		query += " " + filter
	}
	err := db.Select(dest, query, args...)
	return err
}

func (ms *MySQLStorage) Add(dest map[string]interface{}, tableName string, order []string) error {
	db := ms.db
	queryKeys := make([]interface{}, len(order))
	queryArgs := make([]interface{}, len(order))

	login, okay := dest["login"]
	if !okay {
		return fmt.Errorf("login not provided")
	}

	for i, key := range order {
		queryKeys[i] = key
		if val, ok := dest[key]; ok {
			queryArgs[i] = val
		} else {
			return fmt.Errorf("key %s not found in dest map", key)
		}
	}

	keys := strings.Join(order, ", ")
	placeholders := strings.Repeat("?, ", len(queryArgs)-1) + "?"

	ok, err := ms.DataExists(tableName, "login = ?", login)
	if err != nil {
		return fmt.Errorf("error checking if user exists: %v", err)
	}
	if ok {
		return fmt.Errorf("user with login %s already exists", login)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, keys, placeholders)
	_, err = db.Exec(query, queryArgs...)
	return err
}

func (ms *MySQLStorage) DataExists(tableName string, whereClause string, args ...interface{}) (bool, error) {
	db := ms.db
	var exists bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s)", tableName, whereClause)

	err := db.QueryRow(query, args...).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("DataExists: %v", err)
	}

	return exists, nil
}

func (ms *MySQLStorage) ValidateLogin(login string, password string, tableName string) (bool, error) {
	db := ms.db
	var exists bool

	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE login = ? AND password = ?)", tableName)

	err := db.QueryRow(query, login, password).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ValidateLogin: %v", err)
	}

	return exists, nil
}
