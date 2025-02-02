package responsedata

import (
	"cynxhost/internal/constant/types"
	"cynxhost/internal/model/entity"
	"time"
)

type PaginateServerTemplateResponseData struct {
	ServerTemplates []entity.TblServerTemplate `json:"server_templates"`
}

type ServerTemplate struct {
	Id          int              `json:"id"`
	Name        string           `json:"name"`
	Description *string          `json:"description"`
	MinimumRam  int              `json:"minimum_ram"`
	MinimumCpu  int              `json:"minimum_cpu"`
	MinimumDisk int              `json:"minimum_disk"`
	ImageUrl    string           `json:"image_url"`
	Variables   []ScriptVariable `json:"variables"`
}

type ScriptVariable struct {
	Name    string                   `json:"name"`
	Type    types.ScriptVariableType `json:"type"`
	Content []ScriptVariableContent  `json:"content"`
}

type ScriptVariableContent struct {
	Name string `json:"name"`
}

type PaginateServerTemplateCategoriesResponseData struct {
	ServerTemplateCategories []ServerTemplateCategory `json:"server_template_categories"`
}

type ServerTemplateCategory struct {
	Id               int       `json:"id"`
	Name             string    `json:"name"`
	Description      *string   `json:"description"`
	ParentId         int       `json:"parent_id"`
	ImageUrl         *string   `json:"image_url"`
	ServerTemplateId *int      `json:"server_template_id"`
	CreatedDate      time.Time `json:"created_date"`
	UpdatedDate      time.Time `json:"updated_date"`
}

type PaginateUserResponseData struct {
	Users []entity.TblUser `json:"users"`
}

type AuthResponseData struct {
	AccessToken string
	TokenType   string
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
	PersistentNodeId int    `json:"persistent_node_id"`
	Script           string `json:"script"`
}

type GetServerTemplateResponseData struct {
	ServerTemplate ServerTemplate `json:"server_template"`
}
