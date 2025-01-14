package usecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsedata"
)

type UserUseCase interface {
	PaginateUser(ctx context.Context, req request.PaginateRequest, resp *response.APIResponse)
	CheckUsername(ctx context.Context, req request.CheckUsernameRequest, resp *response.APIResponse)
	RegisterUser(ctx context.Context, req request.RegisterUserRequest, resp *response.APIResponse) *responsedata.AuthResponseData
	LoginUser(ctx context.Context, req request.LoginUserRequest, resp *response.APIResponse) *responsedata.AuthResponseData
	GetProfile(ctx context.Context, resp *response.APIResponse) context.Context
}
