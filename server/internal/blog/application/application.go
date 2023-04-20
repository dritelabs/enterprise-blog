package application

import (
	"github.com/dritelabs/blog-reactive/internal/blog/application/commands"
	"github.com/dritelabs/blog-reactive/internal/blog/application/queries"
	"github.com/dritelabs/blog-reactive/internal/blog/application/sagas"
	"github.com/dritelabs/blog-reactive/internal/blog/infrastructure/repositories/memory"
	eventbus "github.com/dritelabs/blog-reactive/internal/shared_kernel/infrastructure/event_bus"
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

func NewApplication() Application {
	blogRepository := memory.NewMemoryBlogRepository(memory.BlogStore{})
	commentRepository := memory.NewMemoryCommentRepository(memory.CommentStore{})
	postRepository := memory.NewMemoryPostRepository(memory.PostStore{})
	eventBus := eventbus.NewLocalEventBus()

	sagas.NewCreateUserSaga(eventBus, postRepository)

	return Application{
		Commands: Commands{
			CreateBlog:    commands.NewCreateBlogCommandHandler(blogRepository),
			CreateComment: commands.NewCreateCommentCommandHandler(commentRepository),
			CreatePost:    commands.NewCreatePostCommandHandler(postRepository, eventBus),
		},
		Queries: Queries{
			GetPost: queries.NewGetPostQueryHandler(postRepository),
		},
	}
}
