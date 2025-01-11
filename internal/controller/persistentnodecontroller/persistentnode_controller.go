package persistentnodecontroller

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

type PersistentNodeController struct {
	persistentNodeUsecase usecase.PersistentNodeUseCase
	validator             *validator.Validate
}

func New(
	persistentNodeUseCase usecase.PersistentNodeUseCase,
	validate *validator.Validate,
) *PersistentNodeController {
	return &PersistentNodeController{
		persistentNodeUsecase: persistentNodeUseCase,
		validator:             validate,
	}
}

func (controller *PersistentNodeController) CreatePersistentNode(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.CreatePersistentNodeRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	ctx = controller.persistentNodeUsecase.CreatePersistentNode(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}
