package events

import "github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"

type PostCreated struct {
	Description string
	ID          string
	Name        string
}

func (e *PostCreated) String() string {
	return "PostCreated"
}

func NewPostCreated(description, id, name string) domain.Event {
	return &PostCreated{
		Description: description,
		ID:          id,
		Name:        name,
	}
}
