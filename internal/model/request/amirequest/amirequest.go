package amirequest

type GetAmiRequest struct {
	AmiId int `json:"ami_id" validate:"omitempty"`
}