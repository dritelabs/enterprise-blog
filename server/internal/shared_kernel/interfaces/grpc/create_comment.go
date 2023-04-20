package gapi

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/application/commands"
	pb "github.com/dritelabs/blog-reactive/internal/blog/infrastructure/pb/blog/v1"
)

func (s *Server) CreateComment(ctx context.Context, in *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	s.blogApplication.Commands.CreateComment.Execute(ctx, commands.CreateCommentCommand{
		Description: in.GetDescription(),
	})

	return &pb.CreateCommentResponse{}, nil
}
