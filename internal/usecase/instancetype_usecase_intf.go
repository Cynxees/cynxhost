package usecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
)

type InstanceTypeUseCase interface {
	PaginateInstanceType(ctx context.Context, req request.PaginateRequest, resp *response.APIResponse)
}
