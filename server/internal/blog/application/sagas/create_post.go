package sagas

import (
	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/events"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
	"github.com/rs/zerolog/log"
)

type CreatePostSaga struct {
	postRepository repositories.PostRepository
}

func (s *CreatePostSaga) Handle(e *events.PostCreated) {
	log.Info().Msg("Hello from a saga")

	s.postRepository.Save(entities.Post{
		Description: e.Description,
		DeletedAt:   e.DeletedAt,
		ID:          e.ID,
		Likes:       e.Likes,
		Name:        e.Name,
	})
}

func NewCreatePostSaga(postRepository repositories.PostRepository) CreatePostSaga {
	return CreatePostSaga{
		postRepository,
	}
}
