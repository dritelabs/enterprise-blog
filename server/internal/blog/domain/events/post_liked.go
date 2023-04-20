package events

import (
	"time"

	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type PostLiked struct {
	CreatedAt time.Time
	ID        string
	UserID    string
}

func (e *PostLiked) String() string {
	return "PostLiked"
}

func NewPostLiked(id, userId string) domain.Event {
	return &PostLiked{
		ID:     id,
		UserID: userId,
	}
}
