package database

import (
	"context"
	"cynxhost/internal/model/entity"
)

type TblAmi interface {
	GetAmi(ctx context.Context, key, value string) (context.Context, entity.TblAmi, error)
	GetAllAmi(ctx context.Context) (context.Context, []entity.TblAmi, error)
}