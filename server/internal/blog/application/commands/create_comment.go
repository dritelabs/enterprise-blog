package commands

import (
	"context"
	"fmt"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
)

type CreateCommentCommand struct {
	Name        string
	Description string
}

type CreateCommentCommandHandler struct {
	commentRepository repositories.CommentRepository
}

func (c *CreateCommentCommandHandler) Execute(ctx context.Context, cmd CreateCommentCommand) error {
	comment := entities.NewComment(cmd.Description)

	fmt.Print(comment)

	return nil
}

func NewCreateCommentCommandHandler(commentRepository repositories.CommentRepository) CreateCommentCommandHandler {
	return CreateCommentCommandHandler{
		commentRepository,
	}
}
