package cmd

import (
	"fm.auth/config"
	"fmt"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

func GRPCServe(ctx *cli.Context) error {
	configs := config.GetConfig()
	listen, err := net.Listen("tcp", ":"+configs.GrpcPort)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge: 30 * time.Second,
		}),
	}

	serve := grpc.NewServer(opts...)
	reflection.Register(serve)

	fmt.Println("GRPC Running: " + configs.GrpcPort)
	return serve.Serve(listen)
}
