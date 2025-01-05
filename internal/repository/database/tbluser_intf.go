package database

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/model/request"
)

type TblUser interface {
	InsertUser(ctx context.Context, user entity.TblUser) (context.Context, int, error)
	GetUser(ctx context.Context, key string, value string) (context.Context, entity.TblUser, error)
	PaginateUser(ctx context.Context, req request.PaginateRequest) (context.Context, []entity.TblUser, error)
	CheckUserExists(ctx context.Context, key string, value string) (context.Context, bool, error)
}
