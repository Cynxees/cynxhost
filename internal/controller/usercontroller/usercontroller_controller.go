package usercontroller

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

func (controller *UserController) CheckUsername(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.CheckUsernameRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	controller.userUsecase.CheckUsername(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}

func (controller *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.RegisterUserRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	resp := controller.userUsecase.RegisterUser(ctx, requestBody, &apiResponse)

	if apiResponse.Code != responsecode.CodeSuccess || resp == nil {
		return ctx, apiResponse
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "AuthToken",
		Value:    resp.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,                    // Set to true if using HTTPS
		SameSite: http.SameSiteStrictMode, // Adjust based on your requirements
	})

	return ctx, apiResponse
}

func (controller *UserController) PaginateUser(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.PaginateRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	controller.userUsecase.PaginateUser(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}

func (controller *UserController) LoginUser(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.LoginUserRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	resp := controller.userUsecase.LoginUser(ctx, requestBody, &apiResponse)

	if apiResponse.Code != responsecode.CodeSuccess || resp == nil {
		return ctx, apiResponse
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "AuthToken",
		Value:    resp.AccessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,                    // Set to true if using HTTPS
		SameSite: http.SameSiteStrictMode, // Adjust based on your requirements
	})

	return ctx, apiResponse
}

func (controller *UserController) GetProfile(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var apiResponse response.APIResponse

	ctx := r.Context()

	ctx = controller.userUsecase.GetProfile(ctx, &apiResponse)

	return ctx, apiResponse
}
