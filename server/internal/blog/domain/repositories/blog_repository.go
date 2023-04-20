package repositories

import (
	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
)

type BlogRepository domain.Repository[entities.Blog]
