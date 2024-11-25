package usecase

import (
	"context"
	"cynxhost/internal/model/entity"
)

type InstanceTypeUseCase interface {
	GetAllInstanceType(ctx context.Context) (context.Context, []entity.TblInstanceType, error)
	GetInstanceType(ctx context.Context, instanceTypeId int) (context.Context, entity.TblInstanceType, error)
}