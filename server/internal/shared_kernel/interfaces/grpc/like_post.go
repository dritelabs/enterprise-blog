package gapi

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/application/commands"
	pb "github.com/dritelabs/blog-reactive/internal/blog/infrastructure/pb/blog/v1"
)

func (s *Server) LikePost(ctx context.Context, in *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	s.blogApplication.Commands.LikePost.Execute(ctx, commands.LikePostCommand{
		PostID: in.GetId(),
		UserID: in.GetUserId(),
	})

	return &pb.LikePostResponse{}, nil
}
