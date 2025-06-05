package models

import (
	"database/sql"
	"time"
)

type Favourite struct {
	ID        string         `db:"id"`
	Comments  sql.NullString `db:"comments"`
	CreatedAt time.Time      `db:"created_at"`
	UserID    int            `db:"user_id"`
}
