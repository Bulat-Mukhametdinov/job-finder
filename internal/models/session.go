package models

import "time"

type Session struct {
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
	CreatedAt time.Time `db:"created_at"`
	UserId    int       `db:"user_id"`
}
