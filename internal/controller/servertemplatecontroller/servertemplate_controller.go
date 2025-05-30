package servertemplatecontroller

import (
	"context"
	"cynxhost/internal/helper"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ServerTemplateController struct {
	uc        usecase.ServerTemplateUseCase
	validator *validator.Validate
}

func New(serverTemplateUseCase usecase.ServerTemplateUseCase, validate *validator.Validate) *ServerTemplateController {
	return &ServerTemplateController{
		uc:        serverTemplateUseCase,
		validator: validate,
	}
}

func (controller *ServerTemplateController) PaginateServerTemplate(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.PaginateRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	controller.uc.PaginateServerTemplate(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}

func (controller *ServerTemplateController) PaginateServerTemplateCategories(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.PaginateServerTemplateCategoryRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	controller.uc.PaginateServerTemplateCategories(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}

func (controller *ServerTemplateController) GetServerTemplate(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.IdRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	controller.uc.GetServerTemplate(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}

func (controller *ServerTemplateController) ValidateServerTemplateVariables(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.ValidateServerTemplateVariablesRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	controller.uc.ValidateServerTemplateVariables(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}