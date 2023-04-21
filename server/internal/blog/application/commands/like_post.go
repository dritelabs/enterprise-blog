package commands

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type LikePostCommand struct {
	PostID string
	UserID string
}

type LikePostCommandHandler struct {
	postRepository repositories.PostRepository
	eventBus       domain.EventBus
}

func (c *LikePostCommandHandler) Execute(ctx context.Context, cmd *LikePostCommand) error {
	post, err := c.postRepository.Get(cmd.PostID)
	if err != nil {
		return err
	}

	post.WithEventBus(c.eventBus)
	post.Like(cmd.UserID)
	post.Commit()

	return nil
}

func NewLikePostCommandHandler(
	postRepository repositories.PostRepository,
	eventBus domain.EventBus,
) LikePostCommandHandler {
	return LikePostCommandHandler{
		postRepository,
		eventBus,
	}
}
