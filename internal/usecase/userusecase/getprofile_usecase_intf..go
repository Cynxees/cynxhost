package userusecase

import (
	"context"
	"cynxhost/internal/model/entity"
)

type GetProfileUseCase interface {
	GetProfile(ctx context.Context, userId int) (context.Context, entity.TblUser, error)	
}