package persistentnodeusecase

import (
	"context"
	"cynxhost/internal/helper"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/model/response/responsedata"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase"
	"strconv"
)

type PersistentNodeUseCaseImpl struct {
	tblPersistentNode database.TblPersistentNode
}

func New(tblPersistentNode database.TblPersistentNode) usecase.PersistentNodeUseCase {
	return &PersistentNodeUseCaseImpl{
		tblPersistentNode: tblPersistentNode,
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
