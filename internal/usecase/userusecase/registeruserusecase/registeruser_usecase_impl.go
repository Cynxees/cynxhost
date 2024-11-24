package registeruserusecase

import (
	"context"
	"cynxhost/internal/dependencies"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase/userusecase"
	"errors"
)

type RegisterUserUseCaseImpl struct {
	tblUser    database.TblUser
	jwtManager *dependencies.JWTManager
}

func New(tblUser database.TblUser, jwtManager *dependencies.JWTManager) userusecase.RegisterUserUseCase {
	return &RegisterUserUseCaseImpl{
		tblUser:    tblUser,
		jwtManager: jwtManager,
	}
}

func (usecase *RegisterUserUseCaseImpl) RegisterUser(ctx context.Context, user entity.TblUser) (context.Context, string, error) {

	ctx, exist, err := usecase.tblUser.CheckUserExists(ctx, "username", user.Username)
	if err != nil {
		return ctx, "", err
	}

	if exist {
		return ctx, "", errors.New("username already exists")
	}

	ctx, id, err := usecase.tblUser.InsertUser(ctx, user)
	if err != nil {
		return ctx, "", err
	}

	token, err := usecase.jwtManager.GenerateToken(id)
	if err != nil {
		return ctx, "", err
	}

	return ctx, token.AccessToken, nil
}
