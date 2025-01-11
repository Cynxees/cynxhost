package app

import (
	"cynxhost/internal/usecase"
	"cynxhost/internal/usecase/persistentnodeusecase"
	"cynxhost/internal/usecase/servertemplateusecase"
	"cynxhost/internal/usecase/userusecase"
)

type UseCases struct {
	UserUseCase           usecase.UserUseCase
	ServerTemplateUseCase usecase.ServerTemplateUseCase
	PersistentNodeUseCase usecase.PersistentNodeUseCase
}

func NewUseCases(repos *Repos, dependencies *Dependencies) *UseCases {

	return &UseCases{
		UserUseCase:           userusecase.New(repos.TblUser, repos.JWTManager),
		ServerTemplateUseCase: servertemplateusecase.New(repos.TblServerTemplate),
		PersistentNodeUseCase: persistentnodeusecase.New(repos.TblPersistentNode, repos.TblInstance, repos.TblInstanceType, repos.TblStorage, repos.TblServerTemplate, dependencies.AWSClient, dependencies.Logger, dependencies.Config),
	}
}
