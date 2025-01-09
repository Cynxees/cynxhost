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

	var convertedPersistentNodes []responsedata.PersistentNode

	// TODO: setup entity so that there are level of data, like 0 to show all, 1 to hide sensitive data, 2 show brief data, 3 show only a little data
	for _, persistentNode := range persistentNodes {
		convertedPersistentNodes = append(convertedPersistentNodes, responsedata.PersistentNode{
			Id:             persistentNode.Id,
			Name:           persistentNode.Name,
			CreatedDate:    persistentNode.CreatedDate,
			UpdatedDate:    persistentNode.UpdatedDate,
			Owner:          persistentNode.Owner,
			ServerTemplate: persistentNode.ServerTemplate,
			InstanceType:   persistentNode.InstanceType,
			Instance:       persistentNode.Instance,
			Storage:        persistentNode.Storage,
			Status:         persistentNode.Status,
		})
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.PaginatePersistentNodeResponseData{
		PersistentNodes: convertedPersistentNodes,
	}
}
