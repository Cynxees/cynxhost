package checkusernamecontroller

import (
	"cynxhost/internal/helper"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CheckusernameController struct {
	uc        usecase.CheckUsernameUseCase
	validator *validator.Validate
}

func New(checkUsernameUseCase usecase.CheckUsernameUseCase, validate *validator.Validate) *CheckusernameController {
	return &CheckusernameController{
		uc:        checkUsernameUseCase,
		validator: validate,
	}
}

func (controller *CheckusernameController) CheckUsername(w http.ResponseWriter, r *http.Request) response.APIResponse {

	var requestBody request.CheckUsernameRequest
	var apiResponse response.APIResponse

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	_, err := controller.uc.CheckUsername(r.Context(), requestBody.Username)
	if err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	apiResponse.Code = responsecode.CodeSuccess
	return apiResponse
}
