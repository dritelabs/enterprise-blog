package commands

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type CreateBlogCommand struct {
	Name        string
	Description string
}

type CreateBlogCommandHandler struct {
	blogRepository repositories.BlogRepository
	eventBus       domain.EventBus
}

func (c *CreateBlogCommandHandler) Execute(ctx context.Context, cmd *CreateBlogCommand) error {
	blog := entities.NewBlog(cmd.Name, cmd.Description)
	blog.WithEventBus(c.eventBus)
	blog.Commit()

	return nil
}

func NewCreateBlogCommandHandler(
	blogRepository repositories.BlogRepository,
	eventBus domain.EventBus,
) CreateBlogCommandHandler {
	return CreateBlogCommandHandler{
		blogRepository,
		eventBus,
	}
}
