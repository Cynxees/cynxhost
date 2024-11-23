package app

import (
	"cynxhost/internal/usecase"
	"cynxhost/internal/usecase/checkusernameusecase"
	"cynxhost/internal/usecase/loginuserusecase"
	"cynxhost/internal/usecase/paginateuserusecase"
	"cynxhost/internal/usecase/registeruserusecase"
)

type Usecases struct {
	RegisterUserUseCase  usecase.RegisterUserUseCase
	LoginUserUseCase     usecase.LoginUserUseCase
	PaginateUserUseCase  usecase.PaginateUserUseCase
	CheckUsernameUseCase usecase.CheckUsernameUseCase
}

func NewUsecases(repos *Repos) *Usecases {

	return &Usecases{
		RegisterUserUseCase: registeruserusecase.New(repos.TblUser, repos.JWTManager),
		LoginUserUseCase:    loginuserusecase.New(repos.TblUser, repos.JWTManager),
		PaginateUserUseCase: paginateuserusecase.New(repos.TblUser),
		CheckUsernameUseCase: checkusernameusecase.New(repos.TblUser),
	}
}
