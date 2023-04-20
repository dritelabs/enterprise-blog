package events

import (
	"time"

	valueobjects "github.com/dritelabs/blog-reactive/internal/blog/domain/value_objects"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type PostCreated struct {
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
) domain.Event {
	return &PostCreated{
		Description: description,
		DeletedAt:   deletedAt,
		ID:          id,
		Likes:       likes,
		Name:        name,
	}
}
