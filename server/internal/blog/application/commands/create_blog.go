package commands

import (
	"context"
	"fmt"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
)

type CreateBlogCommand struct {
	Name        string
	Description string
}

type CreateBlogCommandHandler struct {
	blogRepository repositories.BlogRepository
}

func (c *CreateBlogCommandHandler) Execute(ctx context.Context, cmd CreateBlogCommand) error {
	blog := entities.NewBlog(cmd.Name, cmd.Description)

	fmt.Print(blog)

	return nil
}

func NewCreateBlogCommandHandler(blogRepository repositories.BlogRepository) CreateBlogCommandHandler {
	return CreateBlogCommandHandler{
		blogRepository,
	}
}
