package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectSQLite(dbPath string) (*sqlx.DB, error) {
	return sqlx.Connect("sqlite3", dbPath)
}
