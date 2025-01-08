package usercontroller

import (
	"cynxhost/internal/helper"
	"cynxhost/internal/model/request"
	"cynxhost/internal/model/response"
	"cynxhost/internal/model/response/responsecode"
	"cynxhost/internal/usecase"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userUsecase usecase.UserUseCase
	validator   *validator.Validate
}

func New(
	userUseCase usecase.UserUseCase,
	validate *validator.Validate,
) *UserController {
	return &UserController{
		userUsecase: userUseCase,
		validator:   validate,
	}
}

func (controller *UserController) CheckUsername(w http.ResponseWriter, r *http.Request) response.APIResponse {
	var requestBody request.CheckUsernameRequest
	var apiResponse response.APIResponse
	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	controller.userUsecase.CheckUsername(r.Context(), requestBody, &apiResponse)

	return apiResponse
}

func (controller *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) response.APIResponse {
	var requestBody request.RegisterUserRequest
	var apiResponse response.APIResponse

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	controller.userUsecase.RegisterUser(r.Context(), requestBody, &apiResponse)

	return apiResponse
}

func (controller *UserController) PaginateUser(w http.ResponseWriter, r *http.Request) response.APIResponse {
	var requestBody request.PaginateRequest
	var apiResponse response.APIResponse

	fmt.Println(4)
	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	controller.userUsecase.PaginateUser(r.Context(), requestBody, &apiResponse)

	return apiResponse
}

func (controller *UserController) LoginUser(w http.ResponseWriter, r *http.Request) response.APIResponse {
	var requestBody request.LoginUserRequest
	var apiResponse response.APIResponse

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return apiResponse
	}

	controller.userUsecase.LoginUser(r.Context(), requestBody, &apiResponse)

	return apiResponse
}

func (controller *UserController) GetProfile(w http.ResponseWriter, r *http.Request) response.APIResponse {
	var apiResponse response.APIResponse

	controller.userUsecase.GetProfile(r.Context(), &apiResponse)

	return apiResponse
}
