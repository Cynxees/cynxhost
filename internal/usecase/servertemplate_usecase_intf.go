package usecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
)

type ServerTemplateUseCase interface {
	PaginateServerTemplate(ctx context.Context, req request.PaginateRequest, resp *response.APIResponse)
	GetServerTemplate(ctx context.Context, req request.IdRequest, resp *response.APIResponse)

	PaginateServerTemplateCategories(ctx context.Context, req request.PaginateServerTemplateCategoryRequest, resp *response.APIResponse)

	ValidateServerTemplateVariables(ctx context.Context, req request.ValidateServerTemplateVariablesRequest, resp *response.APIResponse)
}
