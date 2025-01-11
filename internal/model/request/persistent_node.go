package request

type CreatePersistentNodeRequest struct {
	ServerTemplateId int    `json:"server_template_id" validate:"required"`
	StorageSizeMb    int    `json:"storage_size_mb" validate:"required"`
	InstanceTypeId   int    `json:"instance_type_id" validate:"required"`
	Name             string `json:"name" validate:"required"`
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
