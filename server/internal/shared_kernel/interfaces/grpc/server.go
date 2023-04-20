package gapi

import (
	"github.com/dritelabs/blog-reactive/internal/blog/application"
	pb "github.com/dritelabs/blog-reactive/internal/blog/infrastructure/pb/blog/v1"
)

type Server struct {
	pb.UnimplementedBlogServiceServer
	blogApplication application.Application
}

func NewServer(blogApplication application.Application) Server {
	return Server{
		blogApplication: blogApplication,
	}
}
