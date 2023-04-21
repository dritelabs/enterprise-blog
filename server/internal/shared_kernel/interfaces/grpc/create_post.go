package gapi

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/application/commands"
	pb "github.com/dritelabs/blog-reactive/internal/blog/infrastructure/pb/blog/v1"
)

func (s *Server) CreatePost(ctx context.Context, in *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	s.blogApplication.Commands.CreatePost.Execute(ctx, &commands.CreatePostCommand{
		Name:        in.GetName(),
		Description: in.GetDescription(),
	})

	return &pb.CreatePostResponse{}, nil
}
