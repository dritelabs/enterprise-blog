package commands

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
	"github.com/rs/zerolog/log"
)

type CreatePostCommand struct {
	Name        string
	Description string
}

type CreatePostCommandHandler struct {
	postRepository repositories.PostRepository
	eventBus       domain.EventBus
}

func (c *CreatePostCommandHandler) Execute(ctx context.Context, cmd CreatePostCommand) error {
	log.Info().Msgf("executing create post command handler")

	post := entities.NewPost(cmd.Name, cmd.Description)

	log.Info().Msgf("creating new post with id %s", post.ID)

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
