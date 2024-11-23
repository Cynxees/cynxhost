package checkusernameusecase

import (
	"context"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase"
	"errors"
)

type CheckUsernameUseCaseImpl struct {
	tblUser database.TblUser
}

func New(tblUser database.TblUser) usecase.CheckUsernameUseCase {
	return &CheckUsernameUseCaseImpl{
		tblUser: tblUser,
	}
}

func (usecase *CheckUsernameUseCaseImpl) CheckUsername(ctx context.Context, username string) (context.Context, error) {

	if len(username) < 4 {
		return ctx, errors.New("username must be at least 4 characters")
	}

	if len(username) > 20 {
		return ctx, errors.New("username must be at most 20 characters")
	}


	ctx, isExist, err := usecase.tblUser.CheckUserExists(ctx, "username", username)
	if err != nil {
		return ctx, err
	}

	if isExist {
		return ctx, errors.New("username already exist")
	}

	return ctx, nil
}
