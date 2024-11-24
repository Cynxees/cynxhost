package database

import (
	"context"
	"cynxhost/internal/model/entity"
)

type TblParameters interface {
	GetAllParameters(ctx context.Context) (context.Context, []entity.TblParameters, error)
	GetParameters(ctx context.Context, key, value string) (context.Context, entity.TblParameters, error)
}
