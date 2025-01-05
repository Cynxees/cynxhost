package userusecase

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/model/request"
)

type PaginateUserUseCase interface {
	PaginateUser(ctx context.Context, paginateRequest request.PaginateRequest) (context.Context, []entity.TblUser, error)
}
