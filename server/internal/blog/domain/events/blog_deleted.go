package events

import (
	"time"
)

type BlogDeleted struct {
	CreatedAt time.Time
	ID        string
}

func (e *BlogDeleted) String() string {
	return "BlogDeleted"
}

func NewBlogDeleted(id string) *BlogDeleted {
	return &BlogDeleted{
		CreatedAt: time.Now(),
		ID:        id,
	}
}
