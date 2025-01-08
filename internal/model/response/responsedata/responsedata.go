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

type User struct {
	Id          int       `json:"id"`
	Username    string    `json:"username"`
	Coin        int       `json:"coin"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}

type PaginateServerTemplateResponseData struct {
	ServerTemplates []ServerTemplate `json:"server_templates"`
}

type PaginateUserResponseData struct {
	Users []User `json:"users"`
}

type AuthResponseData struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
