package usecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
)

type PersistentNodeUseCase interface {
	CreatePersistentNode(ctx context.Context, req request.CreatePersistentNodeRequest, resp *response.APIResponse) context.Context
	GetPersistentNodes(ctx context.Context, resp *response.APIResponse)
	GetPersistentNode(ctx context.Context, req request.GetPersistentNodeRequest, resp *response.APIResponse) context.Context
	RunPersistentNodeScript(ctx context.Context, req request.RunPersistentNodeScriptRequest, resp *response.APIResponse)

	// Callback
	LaunchCallbackPersistentNode(ctx context.Context, req request.LaunchCallbackPersistentNodeRequest, resp *response.APIResponse) context.Context
	StatusCallbackPersistentNode(ctx context.Context, req request.StatusCallbackPersistentNodeRequest, resp *response.APIResponse) context.Context
}
