package database

import (
	"context"
	"cynxhost/internal/model/entity"
)

type TblInstance interface {
	InsertInstance(ctx context.Context, instance entity.TblInstance) (context.Context, entity.TblInstance, error)
	GetInstance(ctx context.Context, key, value string) (context.Context, entity.TblInstance, error)
}
