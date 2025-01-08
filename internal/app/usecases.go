package app

import (
	"cynxhost/internal/usecase"
	"cynxhost/internal/usecase/servertemplateusecase"
	"cynxhost/internal/usecase/userusecase"
)

type UseCases struct {
	UserUseCase           usecase.UserUseCase
	ServerTemplateUseCase usecase.ServerTemplateUseCase
}

func NewUseCases(repos *Repos) *UseCases {

	return &UseCases{

		UserUseCase:           userusecase.New(repos.TblUser, repos.JWTManager),
		ServerTemplateUseCase: servertemplateusecase.New(repos.TblServerTemplate),
	}
}
