package storage

import (
	"io/ioutil"

	"github.com/jmoiron/sqlx"
)

func RunMigrations(db *sqlx.DB) error {
	sql, err := ioutil.ReadFile("migrations/init.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(sql))
	return err
}
