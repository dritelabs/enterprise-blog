package events

import (
	"time"

	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type BlogCreated struct {
	CreatedAt   time.Time
	Description string
	ID          string
	Name        string
}

func (e *BlogCreated) String() string {
	return "BlogCreated"
}

func NewBlogCreated(description, id, name string) domain.Event {
	return &BlogCreated{
		CreatedAt:   time.Now(),
		Description: description,
		ID:          id,
		Name:        name,
	}
}
