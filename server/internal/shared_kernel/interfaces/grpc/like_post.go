package gapi

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/application/commands"
	pb "github.com/dritelabs/blog-reactive/internal/blog/infrastructure/pb/blog/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) LikePost(ctx context.Context, in *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	err := s.blogApplication.Commands.LikePost.Execute(ctx, commands.LikePostCommand{
		PostID: in.GetId(),
		UserID: in.GetUserId(),
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.LikePostResponse{}, nil
}
