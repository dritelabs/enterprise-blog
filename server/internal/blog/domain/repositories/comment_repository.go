package repositories

import (
	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type CommentRepository domain.Repository[entities.Comment]
