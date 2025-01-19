package responsedata

import (
	"cynxhost/internal/model/entity"
	"time"
)

type PaginateServerTemplateResponseData struct {
	ServerTemplates []ServerTemplate
}

type ServerTemplate struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	MinimumRam  int     `json:"minimum_ram"`
	MinimumCpu  int     `json:"minimum_cpu"`
	MinimumDisk int     `json:"minimum_disk"`
	ImageUrl    *string `json:"image_url"`
}

type PaginateServerTemplateCategoriesResponseData struct {
	ServerTemplateCategory []ServerTemplateCategory
}

type ServerTemplateCategory struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	ParentId         int       `json:"parent_id"`
	ImageUrl         *string   `json:"image_url"`
	ServerTemplateId *int      `json:"server_template_id"`
	CreatedDate      time.Time `json:"created_date"`
	UpdatedDate      time.Time `json:"updated_date"`
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

type GetServerTemplateResponseData struct {
	ServerTemplate entity.TblServerTemplate
}
