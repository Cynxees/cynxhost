package usecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
)

type ServerTemplateUseCase interface {
	PaginateServerTemplate(ctx context.Context, req request.PaginateRequest, resp *response.APIResponse)
	GetServerTemplate(ctx context.Context, req request.IdRequest, resp *response.APIResponse)

	GetServerTemplateCategories(ctx context.Context, req request.GetServerTemplateCategoryRequest, resp *response.APIResponse)
}
