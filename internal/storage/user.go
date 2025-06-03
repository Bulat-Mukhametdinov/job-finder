package storage

import (
	"job-finder/internal/models"

	"github.com/jmoiron/sqlx"
)

type UserStorage struct {
	DB *sqlx.DB
}

func (s *UserStorage) Create(user *models.User) error {
	_, err := s.DB.NamedExec(`
		INSERT INTO user (username, password_hash)
		VALUES (:username, :password_hash)
	`, user)
	return err
}

func (s *UserStorage) GetByUsername(username string) (*models.User, error) {
	var u models.User
	err := s.DB.Get(&u, "SELECT * FROM user WHERE username = ?", username)
	return &u, err
}
