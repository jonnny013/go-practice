package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	goshoppingstore "github.com/jonnny013/go-practice"
)

func NewItemStore(db *sqlx.DB) *ItemStore {
	return &ItemStore{
		DB: db,
	}
}

type ItemStore struct {
	*sqlx.DB
}

func (s *ItemStore) Item(id uuid.UUID) (goshoppingstore.Item, error) {
	var t goshoppingstore.Item
	if err := s.Get(&t, `SELECT * FROM items WHERE id = $1`, id); err != nil {
		return goshoppingstore.Item{}, fmt.Errorf("error getting item: %w", err)
	}
	return t, nil
}

func (s *ItemStore) Items() ([]goshoppingstore.Item, error) {
	var t []goshoppingstore.Item
	if err := s.Select(&t, `SELECT * FROM items`); err != nil {
		return []goshoppingstore.Item{}, fmt.Errorf("error getting item: %w", err)
	}
	return t, nil
}

func (s *ItemStore) CreateItem(t *goshoppingstore.Item) error {
	if err := s.Get(t, `INSERT INTO items VALUES ($1, $2, $3) RETURNING *`, t.Id, t.Name, t.Description); err != nil {
		return fmt.Errorf("error creating item %w", err)
	}
	return nil
}

func (s *ItemStore) UpdateItem(t *goshoppingstore.Item) error {
	if err := s.Get(t, `UPDATE items SET title = $1, description = $2 WHERE id = $3 RETURNING *`, t.Name, t.Description, t.Id); err != nil {
		return fmt.Errorf("error updating item %w", err)
	}
	return nil
}
func (s *ItemStore) DeleteItem(t *goshoppingstore.Item) error {
	if _, err := s.Exec(`DELETE FROM items WHERE id = $1 RETURNING *`, t.Id); err != nil {
		return fmt.Errorf("error updating item %w", err)
	}
	return nil
}
