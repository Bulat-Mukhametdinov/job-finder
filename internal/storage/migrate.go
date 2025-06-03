package storage

import (
	"os"

	"github.com/jmoiron/sqlx"
)

func RunMigrations(db *sqlx.DB) error {
	sql, err := os.ReadFile("migrations/init.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sql))
	return err
}
