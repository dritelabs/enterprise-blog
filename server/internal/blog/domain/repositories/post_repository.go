package repositories

import (
	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type PostRepository domain.Repository[entities.Post]
