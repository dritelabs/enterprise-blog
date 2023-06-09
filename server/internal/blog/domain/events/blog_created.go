package events

import (
	"time"
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

func NewBlogCreated(description, id, name string) *BlogCreated {
	return &BlogCreated{
		CreatedAt:   time.Now(),
		Description: description,
		ID:          id,
		Name:        name,
	}
}
