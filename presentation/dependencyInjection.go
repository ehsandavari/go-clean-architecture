package presentation

import (
	"github.com/ehsandavari/golang-clean-architecture/presentation/api"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
	"log"
)

var Modules []fx.Option

var (
	g errgroup.Group
)

func Runner() {
	g.Go(func() error {
		return api.HttpServers["echo"].ListenAndServe()
	})
	g.Go(func() error {
		return api.HttpServers["gin"].ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
