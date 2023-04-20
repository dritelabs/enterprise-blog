package entities

import (
	"time"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/events"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain/entities"
	"github.com/lucsky/cuid"
)

type Blog struct {
	entities.Aggregate
	Description string
	DeletedAt   *time.Time
	ID          string
	Name        string
}

func (e *Blog) Blog() {
	now := time.Now()
	e.DeletedAt = &now
}

func NewBlog(name, description string) Blog {
	blog := Blog{
		ID:          cuid.New(),
		Name:        name,
		Description: description,
	}

	blog.Apply(events.NewBlogCreated(blog.ID))

	return blog
}
