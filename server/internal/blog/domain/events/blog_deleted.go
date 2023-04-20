package events

import (
	"time"

	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type BlogDeleted struct {
	CreatedAt time.Time
	ID        string
}

func (e *BlogDeleted) String() string {
	return "BlogDeleted"
}

func NewBlogDeleted(id string) domain.Event {
	return &BlogDeleted{
		CreatedAt: time.Now(),
		ID:        id,
	}
}
