package commands

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type CreatePostCommand struct {
	Name        string
	Description string
}

type CreatePostCommandHandler struct {
	postRepository repositories.PostRepository
	eventBus       domain.EventBus
}

func (c *CreatePostCommandHandler) Execute(ctx context.Context, cmd *CreatePostCommand) error {
	post := entities.NewPost(cmd.Name, cmd.Description)
	post.WithEventBus(c.eventBus)
	post.Commit()

	return nil
}

func NewCreatePostCommandHandler(
	postRepository repositories.PostRepository,
	eventBus domain.EventBus,
) CreatePostCommandHandler {
	return CreatePostCommandHandler{
		postRepository,
		eventBus,
	}
}
