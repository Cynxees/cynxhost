package request

type PaginateRequest struct {
	Page int `json:"page" validate:"required"`
	Size int `json:"size" validate:"required"`
}

type PaginateServerTemplateCategoryRequest struct {
	PaginateRequest
	Id     *int   `json:"id"`
	Filter string `json:"filter"`
}

type IdRequest struct {
	Id int `json:"id" validate:"required"`
}
