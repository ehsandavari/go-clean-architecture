package main

import (
	"context"
	"github.com/ehsandavari/golang-clean-architecture/application"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure"
	"github.com/ehsandavari/golang-clean-architecture/persistence"
	"github.com/ehsandavari/golang-clean-architecture/presentation"
	"github.com/ehsandavari/golang-clean-architecture/presentation/api"
	"go.uber.org/fx"
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
