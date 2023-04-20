package events

import "github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"

type CommentCreated struct {
	ID          string
	Description string
}

func (e *CommentCreated) String() string {
	return "CommentCreated"
}

func NewCommentCreated(id string) domain.Event {
	return &CommentCreated{
		ID: id,
	}
}
