package queries

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
)

type GetPostQuery struct {
	ID string
}

type GetPostQueryHandler struct {
	postRepository repositories.PostRepository
}

func (c *GetPostQueryHandler) Execute(ctx context.Context, q GetPostQuery) (*entities.Post, error) {
	post, err := c.postRepository.Get(q.ID)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func NewGetPostQueryHandler(postRepository repositories.PostRepository) GetPostQueryHandler {
	return GetPostQueryHandler{
		postRepository,
	}
}
