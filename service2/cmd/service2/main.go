package main

import (
	"log"
	"os"
	"service2/config"
	"service2/internal/logwrapper"
	"service2/internal/repository"

	"service2/internal/handler"
	"service2/internal/handler/operations"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

func main() {
	config.Cfg = config.NewConfig("service2", "service2", "127.0.0.1", "15432", "service2_db")
	repository.Repo = repository.NewRepository(config.Cfg)
	logwrapper.Logger = logwrapper.NewLogger(logrus.DebugLevel, []logrus.Hook{})
	runMigration()

	swaggerSpec, err := loads.Embedded(handler.SwaggerJSON, handler.FlatSwaggerJSON)
	if err != nil {
		logwrapper.Logger.Payload(logwrapper.NewPayload().Op("load swagger").Package("main")).Fatal(err)
		log.Fatalln(err)
	}

	api := operations.NewService2API(swaggerSpec)
	server := handler.NewServer(api)
	server.Port = 8080
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Service 2: Statefull сервис, который соответствует спецификации Swagger в api/api.yml"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			logwrapper.Logger.Payload(logwrapper.NewPayload().Op("parser.AddGroup").Package("main")).Fatal(err)
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}

		logwrapper.Logger.Payload(logwrapper.NewPayload().Op("parser.Parse").Package("main")).Fatal(err)
		os.Exit(code)
	}

	server.ConfigureAPI()

	logwrapper.Logger.Payload(logwrapper.NewPayload().Package("main")).Info("Starting server")
	if err := server.Serve(); err != nil {
		logwrapper.Logger.Payload(logwrapper.NewPayload().Op("server.Serve").Package("main")).Fatal(err)
		log.Fatalln(err)
	}
}

func runMigration() {
	logwrapper.Logger.Payload(logwrapper.NewPayload().Package("main")).Info("run migrations")
	err := goose.Up(repository.Repo.DB, "./migrations/")
	if err != nil {
		logwrapper.Logger.Payload(logwrapper.NewPayload().Op("runMigration").Package("main")).Fatal(err)
		panic(err)
	}
}
