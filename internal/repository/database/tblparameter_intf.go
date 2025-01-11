package database

import (
	"context"
	"cynxhost/internal/model/entity"
)

type TblParameter interface {
	SelectTblParameters(ctx context.Context, ids []string) (context.Context, []entity.TblParameter, error)
}
