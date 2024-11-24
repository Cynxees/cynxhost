package database

import (
	"context"
	"cynxhost/internal/model/entity"
)

type TblInstanceType interface {
	GetInstanceType(ctx context.Context, key, value string) (context.Context, entity.TblInstanceType, error)
	GetAllInstanceType(ctx context.Context) (context.Context, []entity.TblInstanceType, error)
	UpdateInstanceType(ctx context.Context, instance entity.TblInstanceType) (context.Context, entity.TblInstanceType, error)
}
