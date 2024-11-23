package app

import (
	"context"
)

type App struct {
	Dependencies *Dependencies
	Repos        *Repos
	Usecases     *Usecases
}

func NewApp(ctx context.Context, configPath string) (*App, error) {
	dependencies := NewDependencies(configPath)

	logger := dependencies.Logger

	logger.Infoln("Running database migrations")
	dependencies.DatabaseClient.RunMigrations("migrations")

	logger.Infoln("Initializing Repositories")
	repos := NewRepos(dependencies)

	logger.Infoln("Initializing Usecases")
	usecases := NewUsecases(repos)

	logger.Infoln("App initialized")
	return &App{
		Dependencies: dependencies,
		Repos:        repos,
		Usecases:     usecases,
	}, nil
}
