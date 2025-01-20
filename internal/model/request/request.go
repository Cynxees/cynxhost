package request

type PaginateRequest struct {
	Page    int     `json:"page" validate:"required"`
	Size    int     `json:"size" validate:"required"`
	Sort    *string `json:"sort" validate:"omitempty,oneof=ASC DESC"`
	Keyword *string `json:"keyword"`
}

type PaginateServerTemplateCategoryRequest struct {
	PaginateRequest
	Id     *int   `json:"id"`
	Filter string `json:"filter"`
}

type IdRequest struct {
	Id int `json:"id" validate:"required"`
}
