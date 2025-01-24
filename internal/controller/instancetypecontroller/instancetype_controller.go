package instancetypecontroller

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

type InstanceTypeController struct {
	uc        usecase.InstanceTypeUseCase
	validator *validator.Validate
}

func New(instanceTypeUseCase usecase.InstanceTypeUseCase, validate *validator.Validate) *InstanceTypeController {
	return &InstanceTypeController{
		uc:        instanceTypeUseCase,
		validator: validate,
	}
}

func (controller *InstanceTypeController) PaginateInstanceType(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.PaginateRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	controller.uc.PaginateInstanceType(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}
