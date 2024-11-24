package loginusercontroller

import (
	"cynxhost/internal/helper"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/usecase/userusecase"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type LoginUserController struct {
	uc        userusecase.LoginUserUseCase
	validator *validator.Validate
}

func New(loginUserUseCase userusecase.LoginUserUseCase, validate *validator.Validate) *LoginUserController {
	return &LoginUserController{
		uc:        loginUserUseCase,
		validator: validate,
	}
}

func (controller *LoginUserController) LoginUser(w http.ResponseWriter, r *http.Request) response.APIResponse {

	var requestBody request.LoginUserRequest
	var apiResponse response.APIResponse

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	_, accessToken, err := controller.uc.LoginUser(r.Context(), requestBody.Username, requestBody.Password)
	if err != nil {
		apiResponse.Code = responsecode.CodeAuthenticationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	apiResponse.Data = map[string]string{
		"access_token": accessToken,
		"token_type":   "Bearer",
	}
	apiResponse.Code = responsecode.CodeSuccess
	return apiResponse
}
