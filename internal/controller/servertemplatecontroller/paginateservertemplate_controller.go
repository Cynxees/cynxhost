package servertemplatecontroller

import (
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

func (controller *ServerTemplateController) PaginateServerTemplate(w http.ResponseWriter, r *http.Request) response.APIResponse {
	var requestBody request.PaginateRequest
	var apiResponse response.APIResponse

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	controller.uc.PaginateServerTemplate(r.Context(), requestBody, &apiResponse)

	return apiResponse
}
