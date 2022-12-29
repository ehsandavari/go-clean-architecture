package main

import (
	"context"
	"go.uber.org/fx"
	"golangCodeBase/application"
	"golangCodeBase/infrastructure"
	"golangCodeBase/persistence"
	"golangCodeBase/presentation"
	"golangCodeBase/presentation/api"
)

func main() {
	run()
}

func run() {
	var Modules []fx.Option
	Modules = append(Modules, infrastructure.Modules...)
	Modules = append(Modules, persistence.Modules...)
	Modules = append(Modules, presentation.Modules...)
	Modules = append(Modules, application.Modules...)
	fx.New(
		append(Modules, fx.Invoke(serve))...,
	).Run()
}

func serve(lc fx.Lifecycle, api *api.SApplication) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go api.SetupAPI()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
