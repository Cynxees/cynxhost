package entity

import (
	"cynxhost/internal/constant/types"
	"time"
)

type TblUser struct {
	Id          int       `gorm:"primaryKey" json:"id" visibility:"1"`
	Username    string    `gorm:"size:255;not null;unique" json:"username" visibility:"2"`
	Password    string    `gorm:"size:255;not null" json:"password" visibility:"10"`
	Coin        int       `gorm:"default:0" json:"coin" visibility:"2"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date" visibility:"1"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date" visibility:"1"`
}

type TblScript struct {
	Id             int       `gorm:"primaryKey" json:"id" visibility:"1"`
	CreatedDate    time.Time `gorm:"autoCreateTime" json:"created_date" visibility:"1"`
	UpdatedDate    time.Time `gorm:"autoUpdateTime" json:"updated_date" visibility:"1"`
	Name           string    `gorm:"size:255;not null" json:"name" visibility:"1"`
	Variables      string    `gorm:"type:json;not null" json:"variables" visibility:"2"`
	SetupScript    string    `gorm:"type:text;not null" json:"setup_script" visibility:"2"`
	StartScript    string    `gorm:"type:text;not null" json:"start_script" visibility:"2"`
	StopScript     string    `gorm:"type:text;not null" json:"stop_script" visibility:"2"`
	ShutdownScript string    `gorm:"type:text;not null" json:"shutdown_script" visibility:"2"`
	ConfigPath     string    `gorm:"size:255;not null" json:"config_path" visibility:"2"`
}

type TblServerTemplate struct {
	Id          int       `gorm:"primaryKey" json:"id" visibility:"1"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date" visibility:"1"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date" visibility:"1"`
	Name        string    `gorm:"size:255;not null" json:"name" visibility:"1"`
	Description *string   `gorm:"type:text" json:"description" visibility:"2"`
	MinimumRam  int       `gorm:"not null" json:"minimum_ram" visibility:"1"`
	MinimumCpu  int       `gorm:"not null" json:"minimum_cpu" visibility:"1"`
	MinimumDisk int       `gorm:"not null" json:"minimum_disk" visibility:"1"`
	ImagePath   *string   `gorm:"size:255" json:"image_path" visibility:"10"`
	ScriptId    int       `gorm:"not null" json:"script_id" visibility:"1"`
	Script      TblScript `gorm:"foreignKey:ScriptId" json:"script" visibility:"1"`
}

type TblServerTemplateCategory struct {
	Id               int                `gorm:"primaryKey" json:"id" visibility:"1"`
	CreatedDate      time.Time          `gorm:"autoCreateTime" json:"created_date" visibility:"1"`
	UpdatedDate      time.Time          `gorm:"autoUpdateTime" json:"updated_date" visibility:"1"`
	Name             string             `gorm:"size:255;not null" json:"name" visibility:"1"`
	Description      *string            `gorm:"type:text" json:"description" visibility:"2"`
	ParentId         int                `gorm:"default:null" json:"parent_id"`
	ImagePath        *string            `gorm:"size:255" json:"image_path" visibility:"10"`
	ServerTemplateId *int               `gorm:"default:null" json:"server_template_id"`
	ServerTemplate   *TblServerTemplate `gorm:"foreignKey:ServerTemplateId" json:"server_template"`
}

type TblInstanceType struct {
	Id           int       `gorm:"primaryKey" json:"id" visibility:"1"`
	CreatedDate  time.Time `gorm:"autoCreateTime" json:"created_date" visibility:"1"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime" json:"updated_date" visibility:"1"`
	Name         string    `gorm:"size:255;not null" json:"name" visibility:"1"`
	VcpuCount    int       `gorm:"not null" json:"vcpu_count" visibility:"1"`
	MemorySizeMb int       `gorm:"not null" json:"memory_size_mb" visibility:"1"`
	ImagePath    *string   `gorm:"size:255" json:"image_path" visibility:"10"`
	SpotPrice    float64   `gorm:"type:decimal(10,2);not null" json:"spot_price" visibility:"10"`
	SellPrice    float64   `gorm:"type:decimal(10,2);not null" json:"sell_price" visibility:"1"`
	Status       string    `gorm:"type:enum('ACTIVE', 'INACTIVE', 'HIDDEN');not null" json:"status" visibility:"1"`
}

type TblInstance struct {
	Id             int                  `gorm:"primaryKey" json:"id" visibility:"1"`
	CreatedDate    time.Time            `gorm:"autoCreateTime" json:"created_date" visibility:"1"`
	UpdatedDate    time.Time            `gorm:"autoUpdateTime" json:"updated_date" visibility:"1"`
	Name           string               `gorm:"size:255;not null" json:"name" visibility:"1"`
	AwsInstanceId  string               `gorm:"size:255;not null" json:"aws_instance_id" visibility:"10"`
	PublicIp       string               `gorm:"size:255;not null" json:"public_ip" visibility:"2"`
	PrivateIp      string               `gorm:"size:255;not null" json:"private_ip" visibility:"10"`
	InstanceTypeId int                  `gorm:"not null" json:"instance_type_id" visibility:"1"`
	Status         types.InstanceStatus `gorm:"size:255;not null" json:"status" visibility:"1"`
	InstanceType   TblInstanceType      `gorm:"foreignKey:InstanceTypeId" json:"instance_type" visibility:"1"`
}

type TblStorage struct {
	Id               int                 `gorm:"primaryKey" json:"id"`
	CreatedDate      time.Time           `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate      time.Time           `gorm:"autoUpdateTime" json:"updated_date"`
	Name             string              `gorm:"size:255;not null" json:"name"`
	SizeMb           int                 `gorm:"not null" json:"size_mb"`
	AwsEbsId         string              `gorm:"size:255" json:"aws_ebs_id" visibility:"10"`
	AwsEbsSnapshotId string              `gorm:"size:255" json:"aws_ebs_snapshot_id" visibility:"10"`
	Status           types.StorageStatus `gorm:"size:255;not null" json:"status"`
}

type TblPersistentNode struct {
	Id               int                        `gorm:"primaryKey" json:"id"`
	CreatedDate      time.Time                  `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate      time.Time                  `gorm:"autoUpdateTime" json:"updated_date"`
	Name             string                     `gorm:"size:255;not null" json:"name"`
	OwnerId          int                        `gorm:"not null" json:"owner_id"`
	ServerTemplateId int                        `gorm:"not null" json:"server_template_id"`
	InstanceId       *int                       `gorm:"default:null" json:"instance_id"`
	InstanceTypeId   int                        `gorm:"not null" json:"instance_type_id"`
	StorageId        int                        `gorm:"not null" json:"storage_id"`
	ServerAlias      string                     `gorm:"size:255;not null" json:"server_alias"`
	DnsRecordId      *string                    `gorm:"size:255" json:"dns_record_id" visibility:"10"`
	Status           types.PersistentNodeStatus `gorm:"size:255;not null" json:"status"`
	Owner            TblUser                    `gorm:"foreignKey:OwnerId" json:"owner"`
	ServerTemplate   TblServerTemplate          `gorm:"foreignKey:ServerTemplateId" json:"server_template"`
	Instance         *TblInstance               `gorm:"foreignKey:InstanceId" json:"instance"`
	InstanceType     TblInstanceType            `gorm:"foreignKey:InstanceTypeId" json:"instance_type"`
	Storage          TblStorage                 `gorm:"foreignKey:StorageId" json:"storage"`
}

type TblParameter struct {
	Id          string    `gorm:"primaryKey" json:"id"`
	Value       string    `gorm:"type:text;not null" json:"value"`
	Desc        string    `gorm:"type:text;not null" json:"desc"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
}
