package userusecase

import (
	"context"
	"cynxhost/internal/model/entity"
)

type RegisterUserUseCase interface {
	RegisterUser(ctx context.Context, user entity.TblUser) (context.Context, string, error)
}
