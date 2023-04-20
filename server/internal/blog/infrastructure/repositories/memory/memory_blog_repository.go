package memory

import (
	"errors"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
)

var (
	ErrBlogNotFound = errors.New("blog not found")
)

type BlogStore map[string]entities.Blog

type MemoryBlogRepository struct {
	store BlogStore
}

func (r *MemoryBlogRepository) Get(id string) (*entities.Blog, error) {
	e, ok := r.store[id]
	if !ok {
		return nil, ErrBlogNotFound
	}

	return &e, nil
}

func (r *MemoryBlogRepository) Save(e entities.Blog) (*entities.Blog, error) {
	r.store[e.ID] = e

	return &e, nil
}

func (r *MemoryBlogRepository) Update(e entities.Blog) (*entities.Blog, error) {
	r.store[e.ID] = e

	return &e, nil
}

func NewMemoryBlogRepository(store BlogStore) repositories.BlogRepository {
	return &MemoryBlogRepository{
		store,
	}
}
