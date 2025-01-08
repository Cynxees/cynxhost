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

	var convertedTemplates []responsedata.ServerTemplate
	for _, template := range serverTemplates {
		convertedTemplates = append(convertedTemplates, responsedata.ServerTemplate{
			Id:  template.Id,
			Name: template.Name,
			CreatedDate: template.CreatedDate,
			UpdatedDate: template.UpdatedDate,
			MinimumRam: template.MinimumRam,
			MinimumCpu: template.MinimumCpu,
			MinimumDisk: template.MinimumDisk,
		})
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.PaginateServerTemplateResponseData{
		ServerTemplates: convertedTemplates,
	}
}
