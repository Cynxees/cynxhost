package paginateuserusecase

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/model/request"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase/userusecase"
)

type PaginateUserUseCaseImpl struct {
	tblUser database.TblUser
}

func New(tblUser database.TblUser) userusecase.PaginateUserUseCase {
	return &PaginateUserUseCaseImpl{
		tblUser: tblUser,
	}
}

func (usecase *PaginateUserUseCaseImpl) PaginateUser(ctx context.Context, paginateRequest request.PaginateRequest) (context.Context, []entity.TblUser, error) {

	ctx, users, err := usecase.tblUser.PaginateUser(ctx, paginateRequest)
	if err != nil {
		return ctx, nil, err
	}

	for i := range users {
		users[i].Password = "HIDDEN"
	}

	return ctx, users, nil
}
