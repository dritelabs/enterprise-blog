package events

import "github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"

type BlogDeleted struct {
	ID string
}

func (e *BlogDeleted) String() string {
	return "BlogDeleted"
}

func NewBlogDeleted(id string) domain.Event {
	return &BlogDeleted{
		ID: id,
	}
}
