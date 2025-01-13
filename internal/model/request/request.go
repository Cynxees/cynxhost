package request

type PaginateRequest struct {
	Page int `json:"page" validate:"required"`
	Size int `json:"size" validate:"required"`
}

type GetServerTemplateCategoryRequest struct {
	Id *int `json:"id"`
}

type IdRequest struct {
	Id int `json:"id" validate:"required"`
}