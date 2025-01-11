package responsedata

import (
	"cynxhost/internal/model/entity"
)

type PaginateServerTemplateResponseData struct {
	ServerTemplates []entity.TblServerTemplate
}

type PaginateUserResponseData struct {
	Users []entity.TblUser
}

type AuthResponseData struct {
	AccessToken string
	TokenType   string
}

type GetAllPersistentNodeResponseData struct {
	PersistentNodes []entity.TblPersistentNode
}

type PaginatePersistentNodeResponseData struct {
	PersistentNodes []entity.TblPersistentNode
}

type PaginateInstanceTypeResponseData struct {
	InstanceTypes []entity.TblInstanceType
}

type LaunchCallbackPersistentNodeResponseData struct {
	PersistentNodeId int
	Script           string
}
