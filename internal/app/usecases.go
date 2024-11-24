package app

import (
	"cynxhost/internal/usecase/userusecase"
	"cynxhost/internal/usecase/userusecase/checkusernameusecase"
	"cynxhost/internal/usecase/userusecase/loginuserusecase"
	"cynxhost/internal/usecase/userusecase/paginateuserusecase"
	"cynxhost/internal/usecase/userusecase/registeruserusecase"
)

type Usecases struct {
	RegisterUserUseCase  userusecase.RegisterUserUseCase
	LoginUserUseCase     userusecase.LoginUserUseCase
	PaginateUserUseCase  userusecase.PaginateUserUseCase
	CheckUsernameUseCase userusecase.CheckUsernameUseCase
}

func NewUsecases(repos *Repos) *Usecases {

	return &Usecases{
		RegisterUserUseCase: registeruserusecase.New(repos.TblUser, repos.JWTManager),
		LoginUserUseCase:    loginuserusecase.New(repos.TblUser, repos.JWTManager),
		PaginateUserUseCase: paginateuserusecase.New(repos.TblUser),
		CheckUsernameUseCase: checkusernameusecase.New(repos.TblUser),
	}
}
