package testcontroller

import (
	"cynxhost/internal/dependencies"
	"cynxhost/internal/usecase/testusecase"

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
