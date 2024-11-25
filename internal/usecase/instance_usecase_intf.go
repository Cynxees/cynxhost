package usecase

import (
	"context"
	"cynxhost/internal/model/entity"
)

type InstanceUseCase interface {
	CreateInstance(ctx context.Context, instance entity.TblInstance) (context.Context, entity.TblInstance, error)
	GetInstance(ctx context.Context, key, value string) (context.Context, entity.TblInstance, error)
	UpdateInstance(ctx context.Context, instance entity.TblInstance) (context.Context, entity.TblInstance, error)
	DeleteInstance(ctx context.Context, instanceId int) (context.Context, error)
}