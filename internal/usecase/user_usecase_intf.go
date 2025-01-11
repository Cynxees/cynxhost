package usecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
)

type UserUseCase interface {
	PaginateUser(ctx context.Context, req request.PaginateRequest, resp *response.APIResponse)
	CheckUsername(ctx context.Context, req request.CheckUsernameRequest, resp *response.APIResponse)
	RegisterUser(ctx context.Context, req request.RegisterUserRequest, resp *response.APIResponse)
	LoginUser(ctx context.Context, req request.LoginUserRequest, resp *response.APIResponse)
	GetProfile(ctx context.Context, resp *response.APIResponse) context.Context
}
