package servertemplateusecase

import (
	"context"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/model/response/responsedata"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase"
)

type ServerTemplateUseCaseImpl struct {
	tblServerTemplate database.TblServerTemplate
}

func New(tblServerTemplate database.TblServerTemplate) usecase.ServerTemplateUseCase {
	return &ServerTemplateUseCaseImpl{
		tblServerTemplate: tblServerTemplate,
	}
}

func (usecase *ServerTemplateUseCaseImpl) PaginateServerTemplate(ctx context.Context, req request.PaginateRequest, resp *response.APIResponse) {
	_, serverTemplates, err := usecase.tblServerTemplate.PaginateServerTemplate(ctx, req.Page, req.Size)
	if err != nil {
		resp.Code = responsecode.CodeTblServerTemplateError
		resp.Error = err.Error()
		return
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.PaginateServerTemplateResponseData{
		ServerTemplates: serverTemplates,
	}
}
