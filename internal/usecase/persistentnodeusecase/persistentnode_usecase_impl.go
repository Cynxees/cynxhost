package persistentnodeusecase

import (
	"context"
	"cynxhost/internal/constant/types"
	"cynxhost/internal/dependencies"
	"cynxhost/internal/helper"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/model/response/responsedata"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase"
	"strconv"
)

type PersistentNodeUseCaseImpl struct {
	tblPersistentNode database.TblPersistentNode
	tblStorage        database.TblStorage
	tblScript         database.TblServerTemplate

	awsClient *dependencies.AWSClient
}

func New(tblPersistentNode database.TblPersistentNode, tblStorage database.TblStorage, tblScript database.TblServerTemplate, awsClient *dependencies.AWSClient) usecase.PersistentNodeUseCase {

	return &PersistentNodeUseCaseImpl{
		tblPersistentNode: tblPersistentNode,
		tblStorage:        tblStorage,
		tblScript:         tblScript,

		awsClient: awsClient,
	}
}

func (usecase *PersistentNodeUseCaseImpl) GetPersistentNodes(ctx context.Context, resp *response.APIResponse) {

	contextUser, ok := helper.GetUserFromContext(ctx)
	if !ok {
		resp.Code = responsecode.CodeAuthenticationError
		resp.Error = "User not found in context"
		return
	}

	_, persistentNodes, err := usecase.tblPersistentNode.GetPersistentNodes(ctx, "owner_id", strconv.Itoa(contextUser.Id))
	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = err.Error()
		return
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.PaginatePersistentNodeResponseData{
		PersistentNodes: persistentNodes,
	}
}

func (usecase *PersistentNodeUseCaseImpl) GetPersistentNode(ctx context.Context, req request.GetPersistentNodeRequest, resp *response.APIResponse) context.Context {
	_, persistentNodes, err := usecase.tblPersistentNode.GetPersistentNodes(ctx, "id", strconv.Itoa(req.PersistentNodeId))
	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = err.Error()
		return ctx
	}

	if len(persistentNodes) == 0 {
		resp.Code = responsecode.CodeNotFound
		resp.Error = "Persistent node not found"
		return ctx
	}

	contextUser, ok := helper.GetUserFromContext(ctx)
	if !ok {
		resp.Code = responsecode.CodeAuthenticationError
		resp.Error = "User not found in context"
		return ctx
	}

	persistentNode := persistentNodes[0]
	if persistentNode.OwnerId != contextUser.Id {
		resp.Code = responsecode.CodeForbidden
		resp.Error = "You are not allowed to access this persistent node"
		return ctx
	}

	ctx = helper.SetVisibilityLevelToContext(ctx, types.VisibilityLevelPrivate)

	resp.Code = responsecode.CodeSuccess
	resp.Data = persistentNode
	return ctx
}

func (usecase *PersistentNodeUseCaseImpl) CreatePersistentNode(ctx context.Context, req request.CreatePersistentNodeRequest, resp *response.APIResponse) {

	contextUser, ok := helper.GetUserFromContext(ctx)
	if !ok {
		resp.Code = responsecode.CodeAuthenticationError
		resp.Error = "User not found in context"
		return
	}

	// storage := entity.TblStorage{
	// 	Name:   req.Name,
	// 	SizeMb: req.StorageSizeMb,
	// 	Status: types.StorageStatusNew,
	// }

	persistentNode := entity.TblPersistentNode{
		Name:             req.Name,
		OwnerId:          contextUser.Id,
		ServerTemplateId: req.ServerTemplateId,
		InstanceTypeId:   req.InstanceTypeId,
	}

	ctx, _, err := usecase.tblPersistentNode.CreatePersistentNode(ctx, persistentNode)
	if err != nil {
		resp.Code = responsecode.CodeTblPersistentNodeError
		resp.Error = err.Error()
		return
	}

	resp.Code = responsecode.CodeSuccess
}

func (usecase *PersistentNodeUseCaseImpl) RunPersistentNodeScript(ctx context.Context, req request.RunPersistentNodeScriptRequest, resp *response.APIResponse) {

}
