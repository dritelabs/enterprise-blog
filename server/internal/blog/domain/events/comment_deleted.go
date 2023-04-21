package events

import (
	"time"
)

type CommentDeleted struct {
	CreatedAt time.Time
	ID        string
}

func (e *CommentDeleted) String() string {
	return "CommentDeleted"
}

func NewCommentDeleted(id string) *CommentDeleted {
	return &CommentDeleted{
		CreatedAt: time.Now(),
		ID:        id,
	}
}
