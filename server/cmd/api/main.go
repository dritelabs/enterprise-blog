package main

import (
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	blog "github.com/dritelabs/blog-reactive/internal/blog/application"
	pb "github.com/dritelabs/blog-reactive/internal/blog/infrastructure/pb/blog/v1"
	"github.com/dritelabs/blog-reactive/internal/shared_kernel/infrastructure/config"
	gapi "github.com/dritelabs/blog-reactive/internal/shared_kernel/interfaces/grpc"
)

func main() {
	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	if conf.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	app := blog.NewApplication()

	grpcServer := grpc.NewServer()
	blogServer := gapi.NewServer(app)

	pb.RegisterBlogServiceServer(grpcServer, &blogServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", conf.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create listener")
	}

	log.Info().Msgf("Server listening on %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
