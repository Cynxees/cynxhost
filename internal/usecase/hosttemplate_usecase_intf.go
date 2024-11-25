package usecase

import (
	"context"
	"cynxhost/internal/model/entity"
)

type HostTemplateUseCase interface {
	GetAlluserOwnedHostTemplate(ctx context.Context, userId int) (context.Context, []entity.TblHostTemplate, error)
	GetHostTemplate(ctx context.Context, hostTemplateId int) (context.Context, entity.TblHostTemplate, error)
	CreateHostTemplate(ctx context.Context, hostTemplate entity.TblHostTemplate) (context.Context, entity.TblHostTemplate, error)
	UpdateHostTemplate(ctx context.Context, hostTemplate entity.TblHostTemplate) (context.Context, entity.TblHostTemplate, error)
	DeleteHostTemplate(ctx context.Context, hostTemplateId int) (context.Context, error)
}