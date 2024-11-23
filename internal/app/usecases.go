package app

import (
	"cynxhost/internal/usecase"
	"cynxhost/internal/usecase/loginuserusecase"
	"cynxhost/internal/usecase/paginateuserusecase"
	"cynxhost/internal/usecase/registeruserusecase"
)

type Usecases struct {
	RegisterUserUseCase usecase.RegisterUserUseCase
	LoginUserUseCase    usecase.LoginUserUseCase
	PaginateUserUseCase usecase.PaginateUserUseCase
}

func NewUsecases(repos *Repos) *Usecases {

	return &Usecases{
		RegisterUserUseCase: registeruserusecase.New(repos.TblUser),
		LoginUserUseCase:    loginuserusecase.New(repos.TblUser, repos.JWTManager),
		PaginateUserUseCase: paginateuserusecase.New(repos.TblUser),
	}
}
