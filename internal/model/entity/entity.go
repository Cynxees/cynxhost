package entity

import (
	"cynxhost/internal/constant/types"
	"cynxhost/internal/model/request"
	"database/sql/driver"
	"encoding/json"
	"errors"
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

type TblInstanceType struct {
	Id               int       `gorm:"primaryKey" json:"id" visibility:"1"`
	CreatedDate      time.Time `gorm:"autoCreateTime" json:"created_date" visibility:"1"`
	UpdatedDate      time.Time `gorm:"autoUpdateTime" json:"updated_date" visibility:"1"`
	Name             string    `gorm:"size:255;not null" json:"name" visibility:"1"`
	AwsKey           string    `gorm:"size:255;not null" json:"aws_key" visibility:"10"`
	VcpuCount        int       `gorm:"not null" json:"vcpu_count" visibility:"1"`
	MemorySizeGb     int       `gorm:"not null" visibility:"1"`
	NetworkSpeedMbps int       `gorm:"not null" visibility:"1"`
	ImagePath        *string   `gorm:"size:255" json:"image_path" visibility:"10"`
	SpotPrice        float64   `gorm:"type:decimal(10,2);not null" json:"spot_price" visibility:"10"`
	SellPrice        float64   `gorm:"type:decimal(10,2);not null" json:"sell_price" visibility:"1"`
	Status           string    `gorm:"type:enum('ACTIVE', 'INACTIVE', 'HIDDEN');not null" json:"status" visibility:"1"`
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
	Variables        JSONMap                    `gorm:"type:json" json:"variables"`
}

type TblParameter struct {
	Id          string    `gorm:"primaryKey" json:"id"`
	Value       string    `gorm:"type:text;not null" json:"value"`
	Desc        string    `gorm:"type:text;not null" json:"desc"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date"`
}

type JSONMap []map[string]string

func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func (j *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("failed to scan JSONMap: value is not []byte")
	}

	return json.Unmarshal(bytes, j)
}

func (j *JSONMap) ToScriptVariables() ([]request.ServerTemplateScriptVariable, error) {
	if j == nil {
		return nil, errors.New("JSONMap is nil")
	}

	var result []request.ServerTemplateScriptVariable

	// Iterate over the slice of maps
	for _, item := range *j {

		variable := request.ServerTemplateScriptVariable{
			Name:  item["name"],
			Value: item["value"],
		}
		result = append(result, variable)
	}

	return result, nil
}
