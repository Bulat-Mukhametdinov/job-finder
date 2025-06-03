package models

import "time"

type Favourite struct {
	ID          int       `db:"id"`
	VacancyLink string    `db:"vacancy_link"`
	Comments    string    `db:"comments"`
	CreatedAt   time.Time `db:"created_at"`
	UserID      int       `db:"user_id"`
}
