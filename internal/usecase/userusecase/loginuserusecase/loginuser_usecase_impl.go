package loginuserusecase

import (
	"context"
	"cynxhost/internal/dependencies"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase/userusecase"
	"errors"
)

type LoginUserUseCaseImpl struct {
	tblUser database.TblUser
	jwtManager *dependencies.JWTManager
}

func New(tblUser database.TblUser, jwtManager *dependencies.JWTManager) userusecase.LoginUserUseCase {
	return &LoginUserUseCaseImpl{
		tblUser: tblUser,
		jwtManager: jwtManager,
	}
}

func (usecase *LoginUserUseCaseImpl) LoginUser(ctx context.Context, username string, password string) (context.Context, string, error) {

	ctx, user, err := usecase.tblUser.GetUser(ctx, "username", username)
	if err != nil {
		return ctx, "", err
	}

	if user.Password != password {
		return ctx, "", errors.New("invalid password")
	}

	token, err := usecase.jwtManager.GenerateToken(user.Id)
	if err != nil {
		return ctx, "", err
	}

	return ctx, token.AccessToken, nil
}
