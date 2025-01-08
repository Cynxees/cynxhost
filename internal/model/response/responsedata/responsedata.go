package responsedata

import (
	"time"
)

type ServerTemplate struct {
	Id          int       `json:"id"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
	Name        string    `json:"name"`
	MinimumRam  int       `json:"minimum_ram"`
	MinimumCpu  int       `json:"minimum_cpu"`
	MinimumDisk int       `json:"minimum_disk"`
}

type PaginateServerTemplateResponseData struct {
	ServerTemplates []ServerTemplate `json:"server_templates"`
}
