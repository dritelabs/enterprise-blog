package events

import (
	"time"

	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type PostDeleted struct {
	CreatedAt time.Time
	ID        string
}

func (e *PostDeleted) String() string {
	return "PostDeleted"
}

func NewPostDeleted(id string) domain.Event {
	return &PostDeleted{
		ID: id,
	}
}
