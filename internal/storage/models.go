package storage

import "time"

type User struct {
	ID           int       `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

type Favourite struct {
	ID          int       `db:"id"`
	VacancyLink string    `db:"vacancy_link"`
	Comments    string    `db:"comments"`
	CreatedAt   time.Time `db:"created_at"`
	UserID      int       `db:"user_id"`
}
