package controller

import (
	"cynxhost/internal/helper"
	"cynxhost/internal/model/request/amirequest"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AmiController struct {
	uc        usecase.AmiUseCase
	validator *validator.Validate
}

func NewAmiController(amiUseCase usecase.AmiUseCase, validate *validator.Validate) *AmiController {
	return &AmiController{
		uc:        amiUseCase,
		validator: validate,
	}
}

func (controller *AmiController) GetAllAmi(w http.ResponseWriter, r *http.Request) response.APIResponse {
	var apiResponse response.APIResponse

	_, amis, err := controller.uc.GetAllAmi(r.Context())
	if err != nil {
		apiResponse.Code = responsecode.CodeInternalError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	apiResponse.Code = responsecode.CodeSuccess
	apiResponse.Data = amis
	return apiResponse
}

func (controller *AmiController) GetAmi(w http.ResponseWriter, r *http.Request) response.APIResponse {
  
	var requestBody amirequest.GetAmiRequest
	var apiResponse response.APIResponse
  
	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	_, ami, err := controller.uc.GetAmi(r.Context(), requestBody.AmiId)
	if err != nil {
		apiResponse.Code = responsecode.CodeInternalError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	apiResponse.Code = responsecode.CodeSuccess
	apiResponse.Data = ami
	return apiResponse
}