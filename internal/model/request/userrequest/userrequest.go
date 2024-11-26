package userrequest

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type PaginateUserRequest struct {
	Page int `json:"page" validate:"required"`
	Size int `json:"size" validate:"required"`
}

type CheckUsernameRequest struct {
	Username string `json:"username" validate:"required"`
}