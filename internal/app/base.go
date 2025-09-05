package app

import (
	"job-finder/internal/client/rapid"
	"job-finder/internal/storage"

	"github.com/jmoiron/sqlx"
)

type App struct {
	DB               *sqlx.DB
	FavouriteStorage *storage.FavouriteStorage
	SessionStorage   *storage.SessionStorage
	UserStorage      *storage.UserStorage
	Rapid            *rapid.RapidAPI
}

func NewApp(db *sqlx.DB) *App {
	return &App{
		DB:               db,
		FavouriteStorage: &storage.FavouriteStorage{DB: db},
		SessionStorage:   &storage.SessionStorage{DB: db},
		UserStorage:      &storage.UserStorage{DB: db},
		Rapid:            rapid.NewRapidAPI(),
	}
}
