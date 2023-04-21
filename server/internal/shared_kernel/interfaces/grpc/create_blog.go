package gapi

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/application/commands"
	pb "github.com/dritelabs/blog-reactive/internal/blog/infrastructure/pb/blog/v1"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.CreateBlogRequest) (*pb.CreateBlogResponse, error) {
	s.blogApplication.Commands.CreateBlog.Execute(ctx, &commands.CreateBlogCommand{
		Name:        in.GetName(),
		Description: in.GetDescription(),
	})

	return &pb.CreateBlogResponse{}, nil
}
