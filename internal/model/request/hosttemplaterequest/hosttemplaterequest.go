package hosttemplaterequest

type CreateHostTemplateRequest struct {
	AmiId          int `json:"ami_id" validate:"required"`
	InstanceTypeId int `json:"instance_type_id" validate:"required"`
}

type GetHostTemplateRequest struct {
	HostTemplateId int `json:"host_template_id" validate:"required"`
}

type UpdateHostTemplateRequest struct {
	HostTemplateId int `json:"host_template_id" validate:"required"`
	AmiId          *int `json:"ami_id" validate:"omitempty"`
	InstanceTypeId *int `json:"instance_type_id" validate:"omitempty"`
}

type DeleteHostTemplateRequest struct {
	HostTemplateId int `json:"host_template_id" validate:"required"`
}

type StartHostTemplateRequest struct {
	HostTemplateId int `json:"host_template_id" validate:"required"`
}

type StopHostTemplateRequest struct {
	HostTemplateId int `json:"host_template_id" validate:"required"`
}