package paginateusercontroller

import (
	"cynxhost/internal/helper"
	"cynxhost/internal/model/request/userrequest"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/usecase/userusecase"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type PaginateUserController struct {
	uc        userusecase.PaginateUserUseCase
	validator *validator.Validate
}

func New(paginateUserUseCase userusecase.PaginateUserUseCase, validate *validator.Validate) *PaginateUserController {
	return &PaginateUserController{
		uc:        paginateUserUseCase,
		validator: validate,
	}
}

func (controller *PaginateUserController) PaginateUser(w http.ResponseWriter, r *http.Request) response.APIResponse {

	var requestBody userrequest.PaginateUserRequest
	var apiResponse response.APIResponse

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	_, users, err := controller.uc.PaginateUser(r.Context(), requestBody.Page, requestBody.Size)
	if err != nil {
		apiResponse.Code = responsecode.CodeAuthenticationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	apiResponse.Data = map[string]any{
		"users": users,
	}
	apiResponse.Code = responsecode.CodeSuccess
	return apiResponse
}
