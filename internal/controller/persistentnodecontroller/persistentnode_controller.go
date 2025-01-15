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

func (controller *PersistentNodeController) LaunchCallbackPersistentNode(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.LaunchCallbackPersistentNodeRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	requestBody.ClientIp = helper.GetClientIP(r)

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	ctx = controller.persistentNodeUsecase.LaunchCallbackPersistentNode(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}

func (controller *PersistentNodeController) StatusCallbackPersistentNode(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.StatusCallbackPersistentNodeRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	requestBody.ClientIp = helper.GetClientIP(r)

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	ctx = controller.persistentNodeUsecase.StatusCallbackPersistentNode(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}

func (controller *PersistentNodeController) SendCommandPersistentNode(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.SendCommandPersistentNodeRequest
	var apiResponse response.APIResponse

	ctx := r.Context()
	sessionUser, ok := helper.GetUserFromContext(ctx)
	if !ok {
		apiResponse.Code = responsecode.CodeAuthenticationError
		apiResponse.Error = "User not found in context"
		return ctx, apiResponse
	}

	requestBody.SessionUser = sessionUser
	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	ctx = controller.persistentNodeUsecase.SendCommandPersistentNode(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}

func (controller *PersistentNodeController) ShutdownCallbackPersistentNode(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.ShutdownCallbackPersistentNodeRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	requestBody.ClientIp = helper.GetClientIP(r)
	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	controller.persistentNodeUsecase.ShutdownCallbackPersistentNode(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}

func (controller *PersistentNodeController) ForceShutdownPersistentNode(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.ForceShutdownPersistentNodeRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	sessionUser, ok := helper.GetUserFromContext(ctx)
	if !ok {
		apiResponse.Code = responsecode.CodeAuthenticationError
		apiResponse.Error = "User not found in context"
		return ctx, apiResponse
	}

	requestBody.SessionUser = sessionUser
	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	controller.persistentNodeUsecase.ForceShutdownPersistentNode(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}

func (controller *PersistentNodeController) GetPersistentNodeDetail(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var requestBody request.IdRequest
	var apiResponse response.APIResponse

	ctx := r.Context()

	if err := helper.DecodeAndValidateRequest(r, &requestBody, controller.validator); err != nil {
		apiResponse.Code = responsecode.CodeValidationError
		apiResponse.Error = err.Error()
		return ctx, apiResponse
	}

	controller.persistentNodeUsecase.GetPersistentNode(ctx, requestBody, &apiResponse)

	return ctx, apiResponse
}

func (controller *PersistentNodeController) GetAllPersistentNodesFromUser(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var apiResponse response.APIResponse

	ctx := r.Context()

	controller.persistentNodeUsecase.GetPersistentNodes(ctx, &apiResponse)

	return ctx, apiResponse
}
