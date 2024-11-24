package database

import (
	"context"
	"cynxhost/internal/model/entity"
)

type TblHostTemplate interface {
	CreateHostTemplate(ctx context.Context, hostTemplate entity.TblHostTemplate) (context.Context, int, error)
	GetHostTemplate(ctx context.Context, key, value string) (context.Context, entity.TblHostTemplate, error)
	GetAllUserOwnedHostTemplate(ctx context.Context, userId int) (context.Context, []entity.TblHostTemplate, error)
	UpdateHostTemplate(ctx context.Context, hostTemplate entity.TblHostTemplate) (context.Context, int, error)
	DeleteHostTemplate(ctx context.Context, hostTemplateId int) (context.Context, error)
}