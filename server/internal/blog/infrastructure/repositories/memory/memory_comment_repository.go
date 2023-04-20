package memory

import (
	"errors"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
)

var (
	ErrCommentNotFound = errors.New("comment not found")
)

type CommentStore map[string]entities.Comment

type MemoryCommentRepository struct {
	store CommentStore
}

func (r *MemoryCommentRepository) Get(id string) (*entities.Comment, error) {
	e, ok := r.store[id]
	if !ok {
		return nil, ErrCommentNotFound
	}

	return &e, nil
}

func (r *MemoryCommentRepository) Save(e entities.Comment) (*entities.Comment, error) {
	r.store[e.ID] = e

	return &e, nil
}

func (r *MemoryCommentRepository) Update(e entities.Comment) (*entities.Comment, error) {
	r.store[e.ID] = e

	return &e, nil
}

func NewMemoryCommentRepository(store CommentStore) repositories.CommentRepository {
	return &MemoryCommentRepository{
		store,
	}
}
