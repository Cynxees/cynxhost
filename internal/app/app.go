package app

import (
	"context"
	"cynxhost/internal/dependencies/param"
	"log"
)

type App struct {
	Dependencies *Dependencies
	Repos        *Repos
	Usecases     *UseCases
}

func NewApp(ctx context.Context, configPath string) (*App, error) {

	log.Println("Initializing Dependencies")
	dependencies := NewDependencies(configPath)

	logger := dependencies.Logger

	logger.Infoln("Running database migrations")
	err := dependencies.DatabaseClient.RunMigrations("migrations")
	if err != nil {
		logger.Fatalln("Failed to run migrations: ", err)
		panic(err)
	}

	logger.Infoln("Initializing Repositories")
	repos := NewRepos(dependencies)

	logger.Infoln("Initializing Param")
	go param.SetupStaticParam(repos.TblParameter, logger)

	logger.Infoln("Initializing Usecases")
	usecases := NewUseCases(repos, dependencies)

	logger.Infoln("App initialized")
	return &App{
		Dependencies: dependencies,
		Repos:        repos,
		Usecases:     usecases,
	}, nil
}
