package main

import (
	"net"
	"service1/internal/handler"
	"service1/internal/lg"
	"service1/pkg/composer/hashcompose"

	"github.com/sirupsen/logrus"
	"github.com/vrnvgasu/logwrapper"
	"google.golang.org/grpc"
)

const pack = "main"

func main() {
	lg.Logger = logwrapper.NewLogger(logrus.DebugLevel, []logrus.Hook{})

	lg.Info("serverStar", pack, "Starting server")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		lg.Fatal("net.Listen", pack, err)
	}
	s := grpc.NewServer()
	server := &handler.HashComposeServiceServer{}
	hashcompose.RegisterHashComposeServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		lg.Fatal("server.Serve", pack, err)
	}
}
