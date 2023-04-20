package events

import (
	"time"

	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type CommentDeleted struct {
	CreatedAt time.Time
	ID        string
}

func (e *CommentDeleted) String() string {
	return "CommentDeleted"
}

func NewCommentDeleted(id string) domain.Event {
	return &CommentDeleted{
		CreatedAt: time.Now(),
		ID:        id,
	}
}
