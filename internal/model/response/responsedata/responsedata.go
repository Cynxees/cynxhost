package responsedata

import (
	"cynxhost/internal/model/entity"
)

type PaginateServerTemplateResponseData struct {
	ServerTemplates []entity.TblServerTemplate `json:"server_templates"`
}

type PaginateUserResponseData struct {
	Users []entity.TblUser `json:"users"`
}

type AuthResponseData struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type GetAllPersistentNodeResponseData struct {
	PersistentNodes []entity.TblPersistentNode `json:"persistent_nodes"`
}

type PaginatePersistentNodeResponseData struct {
	PersistentNodes []entity.TblPersistentNode `json:"persistent_nodes"`
}

type PaginateInstanceTypeResponseData struct {
	InstanceTypes []entity.TblInstanceType `json:"instance_types"`
}

type LaunchCallbackPersistentNodeResponseData struct {
	Script entity.TblScript `json:"script"`
	Type   string           `json:"type"`
}
