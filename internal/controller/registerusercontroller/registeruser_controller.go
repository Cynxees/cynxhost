package registerusercontroller

import (
	"cynxhost/internal/model/entity"
	"cynxhost/internal/model/request"
	"cynxhost/internal/usecase"
	"encoding/json"
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

func (controller *RegisterUserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	
	var userRequest request.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := controller.validator.Struct(userRequest); err != nil {
		http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, _, err := controller.uc.RegisterUser(r.Context(), entity.TblUser{
		Username: userRequest.Username,
		Password: userRequest.Password,
	})
	if err != nil {
		http.Error(w, "Failed to register user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write([]byte("User registered successfully")); err != nil {
		http.Error(w, "Failed to write response: "+err.Error(), http.StatusInternalServerError)
	}

}
