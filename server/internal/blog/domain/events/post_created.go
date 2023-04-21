package events

import (
	"time"

	valueobjects "github.com/dritelabs/blog-reactive/internal/blog/domain/value_objects"
)

type PostCreated struct {
	CreatedAt   time.Time
	Description string
	DeletedAt   *time.Time
	ID          string
	Likes       valueobjects.Like
	Name        string
}

func (e *PostCreated) String() string {
	return "PostCreated"
}

func NewPostCreated(
	description string,
	deletedAt *time.Time,
	id string,
	likes valueobjects.Like,
	name string,
) *PostCreated {
	return &PostCreated{
		CreatedAt:   time.Now(),
		Description: description,
		DeletedAt:   deletedAt,
		ID:          id,
		Likes:       likes,
		Name:        name,
	}
}
