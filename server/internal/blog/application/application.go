package application

import (
	"github.com/dritelabs/blog-reactive/internal/blog/application/commands"
	"github.com/dritelabs/blog-reactive/internal/blog/application/queries"
	"github.com/dritelabs/blog-reactive/internal/blog/application/sagas"
	"github.com/dritelabs/blog-reactive/internal/blog/domain/events"
	"github.com/dritelabs/blog-reactive/internal/blog/infrastructure/repositories/memory"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/domain"
	eventbus "github.com/dritelabs/blog-reactive/internal/shared_kernel/infrastructure/event_bus"
	"github.com/rs/zerolog/log"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateBlog    commands.CreateBlogCommandHandler
	CreateComment commands.CreateCommentCommandHandler
	CreatePost    commands.CreatePostCommandHandler
	LikePost      commands.LikePostCommandHandler
}

type Queries struct {
	GetPost queries.GetPostQueryHandler
}

func NewApplication(eventStore domain.EventStore) Application {
	blogRepository := memory.NewMemoryBlogRepository(memory.BlogStore{})
	commentRepository := memory.NewMemoryCommentRepository(memory.CommentStore{})
	postRepository := memory.NewMemoryPostRepository(memory.PostStore{}, eventStore)
	eventBus := eventbus.NewLocalEventBus()
	createPostSaga := sagas.NewCreatePostSaga(postRepository)

	eventBus.Subscribe("PostCreated", func(e domain.Event) {
		createPostSaga.Handle(e.(*events.PostCreated))
	})

	eventBus.Subscribe("PostLiked", func(e domain.Event) {
		ev := *e.(*events.PostLiked)
		log.Info().Msgf("post with id %s liked by user with id %s", ev.ID, ev.UserID)
	})

	return Application{
		Commands: Commands{
			CreateBlog:    commands.NewCreateBlogCommandHandler(blogRepository, eventBus),
			CreateComment: commands.NewCreateCommentCommandHandler(commentRepository, eventBus),
			CreatePost:    commands.NewCreatePostCommandHandler(postRepository, eventBus, eventStore),
			LikePost:      commands.NewLikePostCommandHandler(postRepository, eventBus, eventStore),
		},
		Queries: Queries{
			GetPost: queries.NewGetPostQueryHandler(postRepository),
		},
	}
}
