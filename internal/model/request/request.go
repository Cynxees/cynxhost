package request

type PaginateRequest struct {
	Page int `json:"page" validate:"required"`
	Size int `json:"size" validate:"required"`
}
