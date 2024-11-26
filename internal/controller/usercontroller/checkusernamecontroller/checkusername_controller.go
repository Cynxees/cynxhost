package checkusernamecontroller

import (
	"cynxhost/internal/helper"
	"cynxhost/internal/model/request/userrequest"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/usecase/userusecase"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CheckusernameController struct {
	uc        userusecase.CheckUsernameUseCase
	validator *validator.Validate
}

func New(checkUsernameUseCase userusecase.CheckUsernameUseCase, validate *validator.Validate) *CheckusernameController {
	return &CheckusernameController{
		uc:        checkUsernameUseCase,
		validator: validate,
	}
}

func (controller *CheckusernameController) CheckUsername(w http.ResponseWriter, r *http.Request) response.APIResponse {

	var requestBody userrequest.CheckUsernameRequest
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
