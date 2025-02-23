package testusecase

import (
	"cynxhost/internal/app"
)

type TestUseCaseImpl struct {
	repos        *app.Repos
	dependencies *app.Dependencies
}

func New(repos *app.Repos, dependencies *app.Dependencies) *TestUseCaseImpl {
	return &TestUseCaseImpl{
		repos:        repos,
		dependencies: dependencies,
	}
}
