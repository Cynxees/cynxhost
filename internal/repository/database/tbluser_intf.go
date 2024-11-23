package database

import (
	"context"
	"cynxhost/internal/model/entity"
)

type TblUser interface {
	InsertUser(ctx context.Context, user entity.TblUser) (context.Context, entity.TblUser, error)
	GetUser(ctx context.Context, key string, value string) (context.Context, entity.TblUser, error)
}
