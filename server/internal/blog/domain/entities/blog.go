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

func NewBlog(description, name string) Blog {
	blog := Blog{
		Description: description,
		ID:          cuid.New(),
		Name:        name,
	}

	blog.Apply(events.NewBlogCreated(blog.Description, blog.ID, blog.Name))

	return blog
}
