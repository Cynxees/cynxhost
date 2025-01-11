package instancetypeusecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/model/response/responsedata"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase"
)

type InstanceTypeUseCaseImpl struct {
	tblInstanceType database.TblInstanceType
}

func New(tblInstanceType database.TblInstanceType) usecase.InstanceTypeUseCase {
	return &InstanceTypeUseCaseImpl{
		tblInstanceType: tblInstanceType,
	}
}

func (usecase *InstanceTypeUseCaseImpl) PaginateInstanceType(ctx context.Context, req request.PaginateRequest, resp *response.APIResponse) {
	_, instanceTypes, err := usecase.tblInstanceType.PaginateInstanceType(ctx, req)
	if err != nil {
		resp.Code = responsecode.CodeTblInstanceTypeError
		resp.Error = err.Error()
		return
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.PaginateInstanceTypeResponseData{
		InstanceTypes: instanceTypes,
	}
}
