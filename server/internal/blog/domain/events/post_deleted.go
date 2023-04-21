package events

import (
	"time"
)

type PostDeleted struct {
	CreatedAt time.Time
	ID        string
}

func (e *PostDeleted) String() string {
	return "PostDeleted"
}

func NewPostDeleted(id string) *PostDeleted {
	return &PostDeleted{
		ID: id,
	}
}
