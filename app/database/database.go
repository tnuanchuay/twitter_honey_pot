package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func GetDatabase() (*sql.DB, error){
	db, err := sql.Open("mysql", os.Getenv(ConnectionStringVariableName))
	return db, err
}
