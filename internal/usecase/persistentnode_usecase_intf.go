package usecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
)

type PersistentNodeUseCase interface {
	CreatePersistentNode(ctx context.Context, req request.CreatePersistentNodeRequest, resp *response.APIResponse) context.Context
	GetPersistentNodes(ctx context.Context, resp *response.APIResponse)
	GetPersistentNode(ctx context.Context, req request.IdRequest, resp *response.APIResponse) context.Context
	ForceShutdownPersistentNode(ctx context.Context, req request.ForceShutdownPersistentNodeRequest, resp *response.APIResponse) context.Context

	// Callback
	LaunchCallbackPersistentNode(ctx context.Context, req request.LaunchCallbackPersistentNodeRequest, resp *response.APIResponse) context.Context
	StatusCallbackPersistentNode(ctx context.Context, req request.StatusCallbackPersistentNodeRequest, resp *response.APIResponse) context.Context
	ShutdownCallbackPersistentNode(ctx context.Context, req request.ShutdownCallbackPersistentNodeRequest, resp *response.APIResponse) context.Context

	// Admin Dashboard
	SendCommandPersistentNode(ctx context.Context, req request.SendCommandPersistentNodeRequest, resp *response.APIResponse) context.Context
}
