package events

import (
	"time"
)

type PostLiked struct {
	CreatedAt time.Time
	ID        string
	UserID    string
}

func (e *PostLiked) String() string {
	return "PostLiked"
}

func NewPostLiked(id, userId string) *PostLiked {
	return &PostLiked{
		ID:     id,
		UserID: userId,
	}
}
