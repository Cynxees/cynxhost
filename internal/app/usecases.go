package app

import (
	"cynxhost/internal/usecase"
	"cynxhost/internal/usecase/registeruserusecase"
)

type Usecases struct {
	RegisterUserUseCase usecase.RegisterUserUseCase
}

func NewUsecases(repos *Repos) *Usecases {

	registerUserUseCase := registeruserusecase.New(repos.TblUser)

	return &Usecases{
		RegisterUserUseCase: registerUserUseCase,
	}
}
