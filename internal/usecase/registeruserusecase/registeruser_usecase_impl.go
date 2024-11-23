package registeruserusecase

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase"
)

type RegisterUserUseCaseImpl struct {
	tblUser database.TblUser
}

func New(tblUser database.TblUser) usecase.RegisterUserUseCase {
	return &RegisterUserUseCaseImpl{
		tblUser: tblUser,
	}
}

func (usecase *RegisterUserUseCaseImpl) RegisterUser(ctx context.Context, user entity.TblUser) (context.Context, entity.TblUser, error) {
	ctx, user, err := usecase.tblUser.InsertUser(ctx, user)
	if err != nil {
		return ctx, user, err
	}

	return ctx, user, nil
}