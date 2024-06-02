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

func main() {
	lg.Logger = logwrapper.NewLogger(logrus.DebugLevel, []logrus.Hook{})

	lg.Logger.Payload(logwrapper.NewPayload().Package("main")).Info("Starting server")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	server := &handler.HashComposeServiceServer{}
	hashcompose.RegisterHashComposeServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		lg.Logger.Payload(logwrapper.NewPayload().Op("server.Serve").Package("main")).Fatal(err)
		panic(err)
	}

}
