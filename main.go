package main

import (
	"context"
	"github.com/ehsandavari/golang-clean-architecture/application"
	"github.com/ehsandavari/golang-clean-architecture/application/common/interfaces"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure"
	"github.com/ehsandavari/golang-clean-architecture/persistence"
	"github.com/ehsandavari/golang-clean-architecture/presentation"
	"github.com/ehsandavari/golang-clean-architecture/presentation/http/api"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"log"
)

func main() {
	loadEnv()
	run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file", err)
	}
}

func run() {
	var Modules []fx.Option
	Modules = append(Modules, infrastructure.Modules)
	Modules = append(Modules, persistence.Modules)
	Modules = append(Modules, presentation.Modules)
	Modules = append(Modules, application.Modules...)
	fx.New(
		append(Modules, fx.Invoke(serve))...,
	).Run()
}

func serve(lc fx.Lifecycle, logger interfaces.ILogger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return api.Setup(logger)
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
