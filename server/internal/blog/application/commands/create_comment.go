package commands

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type CreateCommentCommand struct {
	Description string
}

type CreateCommentCommandHandler struct {
	commentRepository repositories.CommentRepository
	eventBus          domain.EventBus
}

func (c *CreateCommentCommandHandler) Execute(ctx context.Context, cmd *CreateCommentCommand) error {
	comment := entities.NewComment(cmd.Description)
	comment.WithEventBus(c.eventBus)
	comment.Commit()

	return nil
}

func NewCreateCommentCommandHandler(
	commentRepository repositories.CommentRepository,
	eventBus domain.EventBus,
) CreateCommentCommandHandler {
	return CreateCommentCommandHandler{
		commentRepository,
		eventBus,
	}
}
