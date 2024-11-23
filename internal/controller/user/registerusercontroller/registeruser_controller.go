package registerusercontroller

import (
	"cynxhost/internal/helper"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/usecase"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type RegisterUserController struct {
	uc        usecase.RegisterUserUseCase
	validator *validator.Validate
}

func New(registerUserUseCase usecase.RegisterUserUseCase, validate *validator.Validate) *RegisterUserController {
	return &RegisterUserController{
		uc:        registerUserUseCase,
		validator: validate,
	}
}

func (controller *RegisterUserController) RegisterUser(w http.ResponseWriter, r *http.Request) response.APIResponse {

	var requestBody request.RegisterUserRequest
	var apiResponse response.APIResponse

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	_, accessToken, err := controller.uc.RegisterUser(r.Context(), entity.TblUser{
		Username: requestBody.Username,
		Password: requestBody.Password,
	})
	if err != nil {
		apiResponse.Code = responsecode.CodeInternalError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	apiResponse.Code = responsecode.CodeSuccess
	apiResponse.Data = map[string]string{
		"access_token": accessToken,
		"token_type": "Bearer",
	}
	return apiResponse
}
