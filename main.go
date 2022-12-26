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
	Application.Modules = append(Application.Modules, Infrastructure.Module)
	Application.Modules = append(Application.Modules, Persistence.Module)
	Application.Modules = append(Application.Modules, Presentation.Module)
	Application.Modules = append(Application.Modules, fx.Invoke(serve))
	fx.New(
		Application.Modules...,
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
