package database

import (
	"context"
	"cynxhost/internal/model/entity"
)

type TblServerTemplate interface {
	CreateServerTemplate(ctx context.Context, serverTemplate entity.TblServerTemplate) (context.Context, int, error)
	GetServerTemplate(ctx context.Context, key string, value string) (context.Context, entity.TblServerTemplate, error)
	PaginateServerTemplate(ctx context.Context, page int, size int) (context.Context, []entity.TblServerTemplate, error)
	DeleteServerTemplate(ctx context.Context, key string, value string) (context.Context, error)

	GetServerTemplateCategories(ctx context.Context, key string, value string) (context.Context, []entity.TblServerTemplateCategory, error)
	GetServerTemplateCategoryChildren(ctx context.Context, parentId *int) (context.Context, []entity.TblServerTemplateCategory, error)
}
