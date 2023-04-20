package events

import "github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"

type CommentDeleted struct {
	ID string
}

func (e *CommentDeleted) String() string {
	return "CommentDeleted"
}

func NewCommentDeleted(id string) domain.Event {
	return &CommentDeleted{
		ID: id,
	}
}
