package entities

import (
	"time"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/events"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain/entities"
	"github.com/lucsky/cuid"
)

type Comment struct {
	entities.Aggregate
	Description string
	DeletedAt   *time.Time
	ID          string
}

func (e *Comment) Delete() {
	now := time.Now()
	e.DeletedAt = &now
}

func NewComment(description string) Comment {
	comment := Comment{
		ID:          cuid.New(),
		Description: description,
	}

	comment.Apply(events.NewCommentCreated(comment.ID))

	return comment
}
