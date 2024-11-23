package database

import (
	"context"
	"cynxhost/internal/model/entity"
)

type TblUser interface {
	InsertUser(ctx context.Context, user entity.TblUser) (context.Context, int, error)
	GetUser(ctx context.Context, key string, value string) (context.Context, entity.TblUser, error)
	PaginateUser(ctx context.Context, page int, size int) (context.Context, []entity.TblUser, error)
	CheckUserExists(ctx context.Context, key string, value string) (context.Context, bool, error)
}
