package request

import contextmodel "cynxhost/internal/model/context"

type CreatePersistentNodeRequest struct {
	ServerTemplateId int                            `json:"server_template_id" validate:"required"`
	StorageSizeGb    int                            `json:"storage_size_gb" validate:"required"`
	InstanceTypeId   int                            `json:"instance_type_id" validate:"required"`
	ServerAlias      string                         `json:"server_alias" validate:"required"`
	Name             string                         `json:"name" validate:"required"`
	Variables        []ServerTemplateScriptVariable `json:"variables"`
}

type ValidateServerTemplateVariablesRequest struct {
	ServerTemplateId int                            `json:"server_template_id" validate:"required"`
	Variables        []ServerTemplateScriptVariable `json:"variables"`
}

type ServerTemplateScriptVariable struct {
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type LaunchCallbackPersistentNodeRequest struct {
	ClientIp string `validate:"required"`

	AwsInstanceId string `json:"aws_instance_id" validate:"required"`
	PublicIp      string `json:"public_ip" validate:"required"`
	EbsVolumeId   string `json:"ebs_volume_id" validate:"required"`
	Type          string `json:"type" validate:"required"`
}

type StatusCallbackPersistentNodeRequest struct {
	ClientIp string `validate:"required"`

	PersistentNodeId int    `json:"persistent_node_id" validate:"required"`
	Type             string `json:"type" validate:"required"`
}

type ShutdownCallbackPersistentNodeRequest struct {
	ClientIp string `validate:"required"`

	PersistentNodeId int `json:"persistent_node_id" validate:"required"`
}

type ForceShutdownPersistentNodeRequest struct {
	SessionUser contextmodel.User `validate:"required"`

	PersistentNodeId int `json:"persistent_node_id" validate:"required"`
}

type PersistentNodeScript string

const (
	Setup    PersistentNodeScript = "SETUP"
	Start    PersistentNodeScript = "START"
	Stop     PersistentNodeScript = "STOP"
	Shutdown PersistentNodeScript = "SHUTDOWN"
)

type RunPersistentNodeScriptRequest struct {
	PersistentNodeId int                  `json:"persistent_node_id" validate:"required"`
	Script           PersistentNodeScript `json:"script" validate:"required"`
}

type GetPersistentNodeRequest struct {
	PersistentNodeId int `json:"persistent_node_id" validate:"required"`
}

type SendCommandPersistentNodeRequest struct {
	SessionUser contextmodel.User `validate:"required"`

	PersistentNodeId int    `json:"persistent_node_id" validate:"required"`
	Command          string `json:"command" validate:"required"`
}
