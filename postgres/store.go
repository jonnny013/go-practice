package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening the database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to the database: %w", err)
	}

	return &Store{
		ItemStore: &ItemStore{
			DB: db,
		},
		CommentStore: &CommentStore{
			DB: db,
		},
	}, nil
}

type Store struct {
	*ItemStore
	*CommentStore
}
