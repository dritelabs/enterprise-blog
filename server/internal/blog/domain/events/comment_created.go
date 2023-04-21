package events

import (
	"time"
)

type CommentCreated struct {
	CreatedAt   time.Time
	Description string
	ID          string
}

func (e *CommentCreated) String() string {
	return "CommentCreated"
}

func NewCommentCreated(description, id string) *CommentCreated {
	return &CommentCreated{
		CreatedAt:   time.Now(),
		Description: description,
		ID:          id,
	}
}
