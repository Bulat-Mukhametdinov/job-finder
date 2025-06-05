package app

import (
	"job-finder/internal/storage"

	"github.com/jmoiron/sqlx"
)

type App struct {
	DB             *sqlx.DB
	SessionStorage *storage.SessionStorage
	UserStorage    *storage.UserStorage
}

func NewApp(db *sqlx.DB) *App {
	return &App{
		DB:             db,
		SessionStorage: &storage.SessionStorage{DB: db},
		UserStorage:    &storage.UserStorage{DB: db},
	}
}
