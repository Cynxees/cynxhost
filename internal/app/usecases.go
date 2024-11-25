package app

import (
	"cynxhost/internal/usecase"
	"cynxhost/internal/usecase/amiusecase"
	"cynxhost/internal/usecase/userusecase"
	"cynxhost/internal/usecase/userusecase/checkusernameusecase"
	"cynxhost/internal/usecase/userusecase/loginuserusecase"
	"cynxhost/internal/usecase/userusecase/paginateuserusecase"
	"cynxhost/internal/usecase/userusecase/registeruserusecase"
)

type UseCases struct {

	// User
	RegisterUserUseCase  userusecase.RegisterUserUseCase
	LoginUserUseCase     userusecase.LoginUserUseCase
	PaginateUserUseCase  userusecase.PaginateUserUseCase
	CheckUsernameUseCase userusecase.CheckUsernameUseCase

	AmiUseCase usecase.AmiUseCase
}

func NewUseCases(repos *Repos) *UseCases {

	return &UseCases{

		// User
		RegisterUserUseCase:  registeruserusecase.New(repos.TblUser, repos.JWTManager),
		LoginUserUseCase:     loginuserusecase.New(repos.TblUser, repos.JWTManager),
		PaginateUserUseCase:  paginateuserusecase.New(repos.TblUser),
		CheckUsernameUseCase: checkusernameusecase.New(repos.TblUser),

		AmiUseCase: amiusecase.New(repos.TblAmi),
	}
}
