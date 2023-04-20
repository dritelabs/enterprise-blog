package gapi

import (
	"context"

	"github.com/dritelabs/blog-reactive/internal/blog/application/queries"
	pb "github.com/dritelabs/blog-reactive/internal/blog/infrastructure/pb/blog/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetPost(ctx context.Context, in *pb.GetPostRequest) (*pb.GetPostResponse, error) {
	post, err := s.blogApplication.Queries.GetPost.Execute(ctx, queries.GetPostQuery{
		ID: in.GetId(),
	})

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "post not found")
	}

	return &pb.GetPostResponse{
		Post: &pb.Post{
			Id:          post.ID,
			Name:        post.Name,
			Description: post.Description,
		},
	}, nil
}
