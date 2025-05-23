package servertemplateusecase

import (
	"context"
	"cynxhost/internal/helper"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/model/response/responsedata"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/repository/externalapi/awsmanager"
	"cynxhost/internal/usecase"
	"strconv"
)

type ServerTemplateUseCaseImpl struct {
	tblServerTemplate database.TblServerTemplate
	awsManager        *awsmanager.AWSManager
}

func New(tblServerTemplate database.TblServerTemplate, awsManager *awsmanager.AWSManager) usecase.ServerTemplateUseCase {
	return &ServerTemplateUseCaseImpl{
		tblServerTemplate: tblServerTemplate,
		awsManager:        awsManager,
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

func (usecase *ServerTemplateUseCaseImpl) PaginateServerTemplateCategories(ctx context.Context, req request.PaginateServerTemplateCategoryRequest, resp *response.APIResponse) {

	_, categories, err := usecase.tblServerTemplate.PaginateServerTemplateCategoryChildren(ctx, req)
	if err != nil {
		resp.Code = responsecode.CodeTblServerTemplateCategoryError
		resp.Error = err.Error()
		return
	}

	var responseServerTemplates []responsedata.ServerTemplateCategory

	for _, template := range categories {

		var url *string
		if template.ImagePath != nil {
			result, err := usecase.awsManager.GetUnsignedURL(*template.ImagePath)
			if err != nil {
				resp.Code = responsecode.CodeAWSError
				resp.Error = err.Error()
				return
			}
			url = result
		}

		responseServerTemplates = append(responseServerTemplates, responsedata.ServerTemplateCategory{
			Id:               template.Id,
			Name:             template.Name,
			Description:      template.Description,
			ParentId:         template.ParentId,
			ImageUrl:         url,
			ServerTemplateId: template.ServerTemplateId,
			CreatedDate:      template.CreatedDate,
			UpdatedDate:      template.UpdatedDate,
		})
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.PaginateServerTemplateCategoriesResponseData{
		ServerTemplateCategories: responseServerTemplates,
	}
}

func (usecase *ServerTemplateUseCaseImpl) GetServerTemplate(ctx context.Context, req request.IdRequest, resp *response.APIResponse) {
	_, template, err := usecase.tblServerTemplate.GetServerTemplate(ctx, "id", strconv.Itoa(req.Id))
	if err != nil {
		resp.Code = responsecode.CodeTblServerTemplateCategoryError
		resp.Error = err.Error()
		return
	}

	var variablesResponse []responsedata.ScriptVariable

	if template.Script.Variables != nil {
		variables := template.Script.Variables

		for _, variable := range variables {
			var contentResponse []responsedata.ScriptVariableContent
			for _, content := range variable.Content {
				contentResponse = append(contentResponse, responsedata.ScriptVariableContent{
					Name: content.Name,
				})
			}
			variablesResponse = append(variablesResponse, responsedata.ScriptVariable{
				Name:    variable.Name,
				Type:    variable.Type,
				Content: contentResponse,
			})
		}
	}

	serverTemplateResponse := responsedata.ServerTemplate{
		Id:          template.Id,
		Name:        template.Name,
		Description: template.Description,
		MinimumRam:  template.MinimumRam,
		MinimumCpu:  template.MinimumCpu,
		MinimumDisk: template.MinimumDisk,
		ImageUrl:    template.ImageUrl,
		Variables:   variablesResponse,
	}

	resp.Code = responsecode.CodeSuccess
	resp.Data = responsedata.GetServerTemplateResponseData{
		ServerTemplate: serverTemplateResponse,
	}
}

func (uc *ServerTemplateUseCaseImpl) ValidateServerTemplateVariables(ctx context.Context, req request.ValidateServerTemplateVariablesRequest, resp *response.APIResponse) {

	// Get Script
	_, serverTemplate, err := uc.tblServerTemplate.GetServerTemplate(ctx, "id", strconv.Itoa(req.ServerTemplateId))
	if err != nil {
		resp.Code = responsecode.CodeTblServerTemplateError
		resp.Error = err.Error()
		return
	}

	if serverTemplate == nil {
		resp.Code = responsecode.CodeNotFound
		resp.Error = "Server template not found"
		return
	}

	// Change struct to map
	_, err = helper.StructToMapStringArray(req.Variables)
	if err != nil {
		resp.Code = responsecode.CodeFailJSON
		resp.Error = err.Error()
		return
	}

	// Check if server variable is valid
	_, err = helper.FormatScriptVariables(serverTemplate.Script.Variables, req.Variables)
	if err != nil {
		resp.Code = responsecode.CodeFailJSON
		resp.Error = err.Error()
		return
	}

	resp.Code = responsecode.CodeSuccess
	return
}
