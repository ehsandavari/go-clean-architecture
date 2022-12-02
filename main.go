package main

import (
	"GolangCodeBase/Application"
	"GolangCodeBase/Application/Common/Interfaces"
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
	fx.New(
		Infrastructure.Module,
		Persistence.Module,
		Presentation.Module,
		Application.Module,
		fx.Invoke(serve),
	).Run()
}

func serve(lc fx.Lifecycle, api *Api.SApplication, commands Interfaces.IOrderHandlerCommands) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go api.SetupAPI()
			return commands.SubscribeOrderCommand(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
