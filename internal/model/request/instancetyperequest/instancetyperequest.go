package instancetyperequest

type GetInstanceTypeRequest struct {
	InstanceTypeId int `json:"instance_type_id" validate:"required"`
}