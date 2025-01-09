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

type InstanceType struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	VcpuCount    int     `json:"vcpu_count"`
	MemorySizeMb int     `json:"memory_size_mb"`
	SellPrice    float64 `json:"sell_price"`
}

type Storage struct {
	Id     int    `json:"id"`
	Size   int    `json:"size"`
	Status string `json:"status"`
}

type Instance struct {
	Id           int          `json:"id"`
	Status       string       `json:"status"`
	PublicIp     string       `json:"public_ip"`
	InstanceType InstanceType `json:"instance_type"`
	CreatedDate  time.Time    `json:"created_date"`
}

type PersistentNode struct {
	Id             int            `json:"id"`
	CreatedDate    time.Time      `json:"created_date"`
	UpdatedDate    time.Time      `json:"updated_date"`
	Owner          User           `json:"owner"`
	ServerTemplate ServerTemplate `json:"server_template"`
	InstanceType   InstanceType   `json:"instance_type"`
	Instance       *Instance      `json:"instance"`
	Storage        Storage        `json:"storage"`
	Name           string         `json:"name"`
	Status         string         `json:"status"`
}

type GetAllPersistentNodeResponseData struct {
	PersistentNodes []PersistentNode `json:"persistent_nodes"`
}

type PaginateInstanceTypeResponseData struct {
	InstanceTypes []InstanceType `json:"instance_types"`
}
