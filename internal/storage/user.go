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
		INSERT INTO users (username, password_hash)
		VALUES (:username, :password_hash)
	`, user)
	return err
}

func (s *UserStorage) GetByUsername(username string) (*models.User, error) {
	var u models.User
	err := s.DB.Get(&u, "SELECT * FROM users WHERE username = ?", username)
	return &u, err
}

func (s *UserStorage) GetByUserID(userId int) (*models.User, error) {
	var u models.User
	err := s.DB.Get(&u, "SELECT * FROM users WHERE id = ?", userId)
	return &u, err
}