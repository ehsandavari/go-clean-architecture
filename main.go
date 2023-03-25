package main

import (
	"github.com/ehsandavari/golang-clean-architecture/application"
	"github.com/ehsandavari/golang-clean-architecture/infrastructure"
	"github.com/ehsandavari/golang-clean-architecture/persistence"
	"github.com/ehsandavari/golang-clean-architecture/presentation"
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
	fx.New(Modules...).Run()
}
