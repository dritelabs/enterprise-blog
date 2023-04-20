package events

import "github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"

type BlogCreated struct {
	ID string
}

func (e *BlogCreated) String() string {
	return "BlogCreated"
}

func NewBlogCreated(id string) domain.Event {
	return &BlogCreated{
		ID: id,
	}
}
