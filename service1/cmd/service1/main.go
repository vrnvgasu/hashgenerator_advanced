package main

import (
	"fmt"
	"net"
	"service1/internal/handler"
	"service1/pkg/composer/hashcompose"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting server ...")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	server := &handler.HashComposeServiceServer{}
	hashcompose.RegisterHashComposeServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
