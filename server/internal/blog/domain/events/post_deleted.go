package events

import "github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"

type PostDeleted struct {
	ID string
}

func (e *PostDeleted) String() string {
	return "PostDeleted"
}

func NewPostDeleted(id string) domain.Event {
	return &PostDeleted{
		ID: id,
	}
}
