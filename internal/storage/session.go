package storage

import (
	"job-finder/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type SessionStorage struct {
	DB *sqlx.DB
}

func (s *SessionStorage) Create(session *models.Session) error {
	_, err := s.DB.NamedExec(`
        INSERT INTO sessions (token, expires_at, created_at, user_id) 
        VALUES (:token, :expires_at, :created_at, :user_id)
    `, session)
	return err
}

func (s *SessionStorage) GetByToken(token string) (*models.Session, error) {
	var session models.Session
	err := s.DB.Get(&session, `
        SELECT * FROM sessions 
        WHERE token = ? AND expires_at > ?
    `, token, time.Now())
	return &session, err
}

func (s *SessionStorage) DeleteByToken(token string) error {
	_, err := s.DB.Exec("DELETE FROM sessions WHERE token = ?", token)
	return err
}

func (s *SessionStorage) DeleteByUserID(userID int) error {
	_, err := s.DB.Exec("DELETE FROM sessions WHERE user_id = ?", userID)
	return err
}

func (s *SessionStorage) CleanExpired() error {
	_, err := s.DB.Exec("DELETE FROM sessions WHERE expires_at < ?", time.Now())
	return err
}
