package sagas

import (
	"github.com/dritelabs/blog-reactive/internal/blog/domain/entities"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/events"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/repositories"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
	"github.com/rs/zerolog/log"
)

type CreateUserSaga struct {
	eventBus       domain.EventBus
	postRepository repositories.PostRepository
}

func (s *CreateUserSaga) Handle(e events.PostCreated) {
	log.Info().Msg("Hello from a saga")

	s.postRepository.Save(entities.Post{
		Description: e.Description,
		ID:          e.ID,
		Name:        e.Name,
	})

	s.eventBus.Publish("PostLiked", events.NewPostLiked(e.ID, "currentUser"))
}

func NewCreateUserSaga(
	eventBus domain.EventBus,
	postRepository repositories.PostRepository,
) CreateUserSaga {
	saga := CreateUserSaga{
		eventBus,
		postRepository,
	}

	eventBus.Subscribe("PostCreated", func(e domain.Event) {
		saga.Handle(*e.(*events.PostCreated))
	})

	return saga
}
