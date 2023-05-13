package main

import (
	"github.com/ehsandavari/go-clean-architecture/application"
	"github.com/ehsandavari/go-clean-architecture/infrastructure"
	"github.com/ehsandavari/go-clean-architecture/persistence"
	"github.com/ehsandavari/go-clean-architecture/presentation"
	"github.com/ehsandavari/go-graceful-shutdown"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func main() {
	loadEnv()
	run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("error in loading .env file", err)
	}
}

func run() {
	newInfrastructure := infrastructure.NewInfrastructure()
	newPersistence := persistence.NewPersistence(newInfrastructure.ILogger, newInfrastructure.SPostgres)
	application.NewApplication(newInfrastructure.SConfig, newInfrastructure.ILogger, newInfrastructure.ITracer, newPersistence.UnitOfWork).Setup()
	newPresentation := presentation.NewPresentation(newInfrastructure.SConfig, newInfrastructure.ILogger)
	newPresentation.Setup()

	shutdownFunc := func() {
		newPresentation.Close()
	}
	cleanupFunc := func() {
		newInfrastructure.Close()
	}
	graceful.Shutdown(shutdownFunc, cleanupFunc, time.Duration(newInfrastructure.SConfig.Service.GracefulShutdownSecond)*time.Second)
}
