package database

import (
	"database/sql"
	"testing"
)

func mysqlInit(t *testing.T) error {
	t.Helper()
	dsn := "root:asodoln1ias@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	_db = db
	return nil
}
