package app

import "context"

type App struct {
	Dependencies *Dependencies
	Repos        *Repos
	Usecases     *Usecases
}

func NewApp(ctx context.Context, configPath string) (*App, error) {
	dependencies := NewDependencies(configPath)
	repos := NewRepos(dependencies)
	usecases := NewUsecases(repos)

	return &App{
		Dependencies: dependencies,
		Repos:        repos,
		Usecases:     usecases,
	}, nil
}