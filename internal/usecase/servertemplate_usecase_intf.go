package usecase

import (
	"context"
	"cynxhost/internal/model/entity"
)

type ServerTemplateUseCase interface {
	PaginateServerTemplate(ctx context.Context) (context.Context, []entity.TblServerTemplate, error)
}
