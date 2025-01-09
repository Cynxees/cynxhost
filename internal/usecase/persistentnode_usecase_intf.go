package usecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
)

type PersistentNodeUseCase interface {
	CreatePersistentNode(ctx context.Context, req request.CreatePersistentNodeRequest, resp *response.APIResponse)
	GetPersistentNodes(ctx context.Context, resp *response.APIResponse)
	GetPersistentNode(ctx context.Context, req request.GetPersistentNodeRequest, resp *response.APIResponse)
	RunPersistentNodeScript(ctx context.Context, req request.RunPersistentNodeScriptRequest, resp *response.APIResponse)
}