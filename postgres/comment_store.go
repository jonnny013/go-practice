package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	goshoppingstore "github.com/jonnny013/go-practice"
)

type CommentStore struct {
	*sqlx.DB
}

func (s *CommentStore) Comment(id uuid.UUID) (goshoppingstore.Comment, error) {
	var t goshoppingstore.Comment
	if err := s.Get(&t, `SELECT * FROM comments WHERE id = $1`, id); err != nil {
		return goshoppingstore.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}
	return t, nil
}

func (s *CommentStore) Comments() ([]goshoppingstore.Comment, error) {
	var t []goshoppingstore.Comment
	if err := s.Select(&t, `SELECT * FROM comments`); err != nil {
		return []goshoppingstore.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}
	return t, nil
}

func (s *CommentStore) CommentsByItem(postId uuid.UUID) ([]goshoppingstore.Comment, error) {
	var cc []goshoppingstore.Comment
	if err := s.Select(&cc, `SELECT * FROM comments where post_id = $1`, postId); err != nil {
		return []goshoppingstore.Comment{}, fmt.Errorf("error getting comments: %w", err)
	}
	return cc, nil
}

func (s *CommentStore) CreateComment(t *goshoppingstore.Comment) error {
	if err := s.Get(t, `INSERT INTO comments (id, item_id, title, body, likes) VALUES ($1, $2, $3, $4, $5) RETURNING *`, t.Id, t.Item_Id, t.Title, t.Body, t.Likes); err != nil {
		return fmt.Errorf("error creating comment %w", err)
	}
	return nil
}

func (s *CommentStore) UpdateComment(t *goshoppingstore.Comment) error {
	if err := s.Get(t, `UPDATE comments SET title = $1, body = $2, likes = $3 WHERE id = $4 RETURNING *`, t.Title, t.Body, t.Likes, t.Id); err != nil {
		return fmt.Errorf("error updating comment %w", err)
	}
	return nil
}
func (s *CommentStore) DeleteComment(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM comments WHERE id = $1 RETURNING *`, id); err != nil {
		return fmt.Errorf("error updating comment %w", err)
	}
	return nil
}
