package usecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
)

type ServerTemplateUseCase interface {
	PaginateServerTemplate(ctx context.Context, req request.PaginateRequest, resp *response.APIResponse)
}
