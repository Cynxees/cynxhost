package database

import (
	"context"
	"cynxhost/internal/model/entity"
)

type TblScript interface {
	CreateScript(ctx context.Context, script entity.TblScript) (context.Context, int, error)
	GetScript(ctx context.Context, key string, value string) (context.Context, entity.TblScript, error)
	DeleteScript(ctx context.Context, key string, value string) (context.Context, error)
}
