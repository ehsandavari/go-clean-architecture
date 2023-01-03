package main

import (
	"context"
	"github.com/ehsandavari/golang-clean-architecture/application"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure"
	"github.com/ehsandavari/golang-clean-architecture/persistence"
	"github.com/ehsandavari/golang-clean-architecture/presentation"
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

func serve(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go presentation.Runner()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
