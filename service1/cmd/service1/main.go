package main

import (
	"net"
	"os"
	"os/signal"
	"service1/config"
	"service1/internal/handler"
	"service1/internal/lg"
	"service1/pkg/composer/hashcompose"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vrnvgasu/logwrapper"
	"google.golang.org/grpc"
)

const pack = "main"

func main() {
	lg.Logger = logwrapper.NewLogger(logrus.DebugLevel, []logrus.Hook{})
	config.Cfg = config.MustLoad()
	lg.Logger = logwrapper.NewLogger(logrus.Level(config.Cfg.DebugLevel), []logrus.Hook{})

	lg.Info("serverStar", pack, "Starting server")
	lis, err := net.Listen("tcp", ":"+config.Cfg.Port)
	if err != nil {
		lg.Fatal("net.Listen", pack, err)
	}
	s := grpc.NewServer()
	server := &handler.HashComposeServiceServer{}
	hashcompose.RegisterHashComposeServiceServer(s, server)

	go func(server *handler.HashComposeServiceServer) {
		if err := s.Serve(lis); err != nil {
			lg.Fatal("server.Serve", pack, err)
		}
	}(server)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	lg.Info("server.Serve", pack, "Server will shutdown gracefully")
	time.Sleep(5 * time.Second)
	lg.Info("server.Serve", pack, "Server stopped")
}
