package goshoppingstore

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jonnny013/go-practice/api"
)

type Item struct {
	Id          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Comment struct {
	Id        uuid.UUID `json:"id" db:"id"`
	Item_Id   uuid.UUID `json:"item_id" db:"item_id"`
	Title     string    `json:"title" db:"title"`
	Body      string    `json:"body" db:"body"`
	Likes     int       `json:"likes" db:"likes"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ItemStore interface {
	Item(id uuid.UUID) (Item, error)
	Items() ([]Item, error)
	CreateItem(i *Item) error
	UpdateItem(i *Item) error
	DeleteItem(id uuid.UUID) error
}

type CommentStore interface {
	Comment(id uuid.UUID) (Comment, error)
	CommentsByItem(postId uuid.UUID) ([]Comment, error)
	CreateComment(i *Comment) error
	UpdateComment(i *Comment) error
	DeleteComment(id uuid.UUID) error
}

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
