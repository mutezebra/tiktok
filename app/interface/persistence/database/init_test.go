package database

import (
	"database/sql"
	"testing"
)

func mysqlInit(t *testing.T) error {
	t.Helper()
	dsn := "your dsn"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	_db = db
	return nil
}
