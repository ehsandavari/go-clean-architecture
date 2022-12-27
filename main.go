package main

import (
	"GolangCodeBase/Application"
	"GolangCodeBase/Infrastructure"
	"GolangCodeBase/Persistence"
	"GolangCodeBase/Presentation"
	"GolangCodeBase/Presentation/Api"
	"context"
	"go.uber.org/fx"
)

func main() {
	run()
}

func run() {
	var Modules []fx.Option
	Modules = append(Modules, Infrastructure.Modules...)
	Modules = append(Modules, Persistence.Modules...)
	Modules = append(Modules, Presentation.Modules...)
	Modules = append(Modules, Application.Modules...)
	fx.New(
		append(Modules, fx.Invoke(serve))...,
	).Run()
}

func serve(lc fx.Lifecycle, api *Api.SApplication) {
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
