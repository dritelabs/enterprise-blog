package commands

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
)

type LikePostCommand struct {
	PostID string
	UserID string
}

type LikePostCommandHandler struct {
	postRepository repositories.PostRepository
}

func (c *LikePostCommandHandler) Execute(ctx context.Context, cmd LikePostCommand) error {
	post, err := c.postRepository.Get(cmd.PostID)
	if err != nil {
		return err
	}

	post.Like(cmd.UserID)

	return nil
}

func NewLikePostCommandHandler(postRepository repositories.PostRepository) LikePostCommandHandler {
	return LikePostCommandHandler{
		postRepository,
	}
}
