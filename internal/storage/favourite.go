package storage

import (
	"job-finder/internal/models"

	"github.com/jmoiron/sqlx"
)

type FavouriteStorage struct {
	DB *sqlx.DB
}

func (s *FavouriteStorage) Create(vacancy *models.Favourite) error {
	_, err := s.DB.NamedExec(`
		INSERT INTO favourites (id, user_id)
		VALUES (:id, :user_id)
	`, vacancy)
	return err
}

func (s *FavouriteStorage) Delete(vacancy *models.Favourite) error {
	_, err := s.DB.NamedExec(`
		DELETE FROM favourites
		WHERE id = :id
	`, vacancy)
	return err
}

func (s *FavouriteStorage) UpdateComment(id string, newСomment string) error {
	_, err := s.DB.Exec(`
		UPDATE favourites
		SET comment = ?
		WHERE id = ?
	`, newСomment, id)

	return err
}

func (s *FavouriteStorage) GetById(id string) (*models.Favourite, error) {
	var u models.Favourite
	err := s.DB.Get(&u, "SELECT * FROM favourites WHERE id = ?", id)
	return &u, err
}

func (s *FavouriteStorage) GetByUserID(userId int) ([]models.Favourite, error) {
	var favourites []models.Favourite
	err := s.DB.Select(&favourites, "SELECT * FROM favourites WHERE user_id = ?", userId)
	return favourites, err
}
