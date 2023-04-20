package entities

import (
	"time"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/events"
	valueobjects "github.com/dritelabs/blog-reactive/internal/blog/domain/value_objects"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain/entities"
	"github.com/lucsky/cuid"
)

type Post struct {
	entities.Aggregate
	Description string
	DeletedAt   *time.Time
	ID          string
	Name        string
	Likes       valueobjects.Like
}

func (e *Post) Delete() {
	now := time.Now()
	e.DeletedAt = &now
}

func (e *Post) Like(userId string) {
	if e.Likes[userId] {
		delete(e.Likes, userId)
	} else {
		e.Likes[userId] = true
	}

	e.Apply(events.NewPostLiked(e.ID, userId))

}

func NewPost(description, name string) Post {
	post := Post{
		Description: description,
		ID:          cuid.New(),
		Name:        name,
	}

	post.Apply(events.NewPostCreated(post.Description, post.ID, post.Name))

	return post
}
