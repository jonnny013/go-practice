package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	goshoppingstore "github.com/jonnny013/go-practice"
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
		ItemStore:    NewItemStore(db),
		CommentStore: NewCommentStore(db),
	}, nil
}

type Store struct {
	goshoppingstore.ItemStore
	goshoppingstore.CommentStore
}
