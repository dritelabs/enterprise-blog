package commands

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/events"
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
	eventStore     domain.EventStore
}

func (c *CreatePostCommandHandler) Execute(ctx context.Context, cmd *CreatePostCommand) error {
	post := entities.NewPost(cmd.Description, cmd.Name)

	c.eventStore.Store(ctx, events.NewPostCreated(
		post.Description,
		post.DeletedAt,
		post.ID,
		post.Likes,
		post.Name,
	))

	post.WithEventBus(c.eventBus)
	post.Commit()

	return nil
}

func NewCreatePostCommandHandler(
	postRepository repositories.PostRepository,
	eventBus domain.EventBus,
	eventStore domain.EventStore,
) CreatePostCommandHandler {
	return CreatePostCommandHandler{
		postRepository,
		eventBus,
		eventStore,
	}
}
