package events

import (
	"time"

	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type CommentCreated struct {
	CreatedAt   time.Time
	Description string
	ID          string
}

func (e *CommentCreated) String() string {
	return "CommentCreated"
}

func NewCommentCreated(description, id string) domain.Event {
	return &CommentCreated{
		CreatedAt:   time.Now(),
		Description: description,
		ID:          id,
	}
}
