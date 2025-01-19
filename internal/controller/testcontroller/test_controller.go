package testcontroller

import (
	"context"
	"cynxhost/internal/dependencies"
	"cynxhost/internal/model/response"
	"cynxhost/internal/usecase/testusecase"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type TestController struct {
	testUsecase *testusecase.TestUseCaseImpl
	validator   *validator.Validate
	config      *dependencies.Config
}

func New(
	testUsecase *testusecase.TestUseCaseImpl,
	validate *validator.Validate,
	config *dependencies.Config,
) *TestController {
	return &TestController{
		testUsecase: testUsecase,
		validator:   validate,
		config:      config,
	}
}

func (c *TestController) CreateDNS(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var apiResponse response.APIResponse

	ctx := r.Context()

	c.testUsecase.CreateDNS(&apiResponse)
	return ctx, apiResponse
}

func (c *TestController) RetrieveDNS(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var apiResponse response.APIResponse

	ctx := r.Context()

	c.testUsecase.RetrieveDNS(&apiResponse)
	return ctx, apiResponse
}

func (c *TestController) UpdateDNS(w http.ResponseWriter, r *http.Request) (context.Context, response.APIResponse) {
	var apiResponse response.APIResponse

	ctx := r.Context()

	c.testUsecase.UpdateDNS(&apiResponse)
	return ctx, apiResponse
}