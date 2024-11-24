package userusecase

import (
	"context"
	"cynxhost/internal/model/entity"
)

type PaginateUserUseCase interface {
	PaginateUser(ctx context.Context, page int, size int) (context.Context, []entity.TblUser, error)
}
