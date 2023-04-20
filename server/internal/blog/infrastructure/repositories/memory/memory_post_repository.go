package memory

import (
	"errors"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
)

var (
	ErrPostNotFound = errors.New("post not found")
)

type PostStore map[string]entities.Post

type MemoryPostRepository struct {
	store PostStore
}

func (r *MemoryPostRepository) Get(id string) (*entities.Post, error) {
	e, ok := r.store[id]
	if !ok {
		return nil, ErrPostNotFound
	}

	return &e, nil
}

func (r *MemoryPostRepository) Save(e entities.Post) (*entities.Post, error) {
	r.store[e.ID] = e

	return &e, nil
}

func (r *MemoryPostRepository) Update(e entities.Post) (*entities.Post, error) {
	r.store[e.ID] = e

	return &e, nil
}

func NewMemoryPostRepository(store PostStore) repositories.PostRepository {
	return &MemoryPostRepository{
		store,
	}
}
