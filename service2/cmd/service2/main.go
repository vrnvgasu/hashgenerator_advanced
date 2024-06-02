package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"service2/config"
	"service2/internal/lg"
	"service2/internal/repository"
	"syscall"
	"time"

	"service2/internal/handler"
	"service2/internal/handler/operations"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
	"github.com/vrnvgasu/logwrapper"
)

const pack = "main"

func main() {
	serverIsStopped := false

	lg.Logger = logwrapper.NewLogger(logrus.DebugLevel, []logrus.Hook{})
	config.Cfg = config.NewConfig("service2", "service2", "127.0.0.1", "15432", "service2_db")
	repository.Repo = repository.NewRepository(config.Cfg)
	runMigration()

	ctx := context.Background()

	swaggerSpec, err := loads.Embedded(handler.SwaggerJSON, handler.FlatSwaggerJSON)
	if err != nil {
		lg.Fatal(ctx, "load swagger", pack, err)
	}

	api := operations.NewService2API(swaggerSpec)
	server := handler.NewServer(api)
	server.Port = 8080
	defer func() {
		if !serverIsStopped {
			server.Shutdown()
		}
	}()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Service 2: Statefull сервис, который соответствует спецификации Swagger в api/api.yml"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			lg.Fatal(ctx, "parser.AddGroup", pack, err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}

		lg.Fatal(ctx, "parser.Parse", pack, err)
		os.Exit(code)
	}

	server.ConfigureAPI()

	go func(server *handler.Server) {
		lg.Info(ctx, "server.Serve", pack, "Starting server")
		if err := server.Serve(); err != nil && err != http.ErrServerClosed {
			lg.Fatal(ctx, "server.Serve", pack, err)
		}
	}(server)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c

	lg.Info(ctx, "server.Serve", pack, "Server will shutdown gracefully")
	time.Sleep(5 * time.Second)
	if err := server.Shutdown(); err != nil {
		lg.Fatal(ctx, "server.Serve", pack, err)
	}
	serverIsStopped = true
	lg.Info(ctx, "server.Serve", pack, "Server stopped")
}

func runMigration() {
	lg.Info(context.Background(), "runMigration", pack, "run migrations")
	err := goose.Up(repository.Repo.DB, "./migrations/")
	if err != nil {
		lg.Fatal(context.Background(), "runMigration", pack, err)
		panic(err)
	}
}
