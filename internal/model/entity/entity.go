package entity

import "time"

type TblUser struct {
	Id          int       `gorm:"primaryKey"`
	Username    string    `gorm:"size:255;not null;unique"`
	Password    string    `gorm:"size:255;not null"`
	Coin        int       `gorm:"default:0"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
}

type TblScript struct {
	Id             int       `gorm:"primaryKey"`
	CreatedDate    time.Time `gorm:"autoCreateTime"`
	UpdatedDate    time.Time `gorm:"autoUpdateTime"`
	Name           string    `gorm:"size:255;not null"`
	Variables      string    `gorm:"type:json;not null"`
	SetupScript    string    `gorm:"type:text;not null"`
	StartScript    string    `gorm:"type:text;not null"`
	StopScript     string    `gorm:"type:text;not null"`
	ShutdownScript string    `gorm:"type:text;not null"`
}

type TblServerTemplate struct {
	Id          int       `gorm:"primaryKey"`
	CreatedDate time.Time `gorm:"autoCreateTime"`
	UpdatedDate time.Time `gorm:"autoUpdateTime"`
	Name        string    `gorm:"size:255;not null"`
	MinimumRam  int       `gorm:"not null"`
	MinimumCpu  int       `gorm:"not null"`
	MinimumDisk int       `gorm:"not null"`
	ScriptId    int       `gorm:"not null"`
	Script      TblScript `gorm:"foreignKey:ScriptId"`
}

type TblInstanceType struct {
	Id           int       `gorm:"primaryKey"`
	CreatedDate  time.Time `gorm:"autoCreateTime"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime"`
	Name         string    `gorm:"size:255;not null"`
	VcpuCount    int       `gorm:"not null"`
	MemorySizeMb int       `gorm:"not null"`
	SpotPrice    float64   `gorm:"type:decimal(10,2);not null"`
	SellPrice    float64   `gorm:"type:decimal(10,2);not null"`
	Status       string    `gorm:"type:enum('ACTIVE', 'INACTIVE', 'HIDDEN');not null"`
}

type TblInstance struct {
	Id             int             `gorm:"primaryKey"`
	CreatedDate    time.Time       `gorm:"autoCreateTime"`
	UpdatedDate    time.Time       `gorm:"autoUpdateTime"`
	Name           string          `gorm:"size:255;not null"`
	AwsInstanceId  string          `gorm:"size:255;not null"`
	PublicIp       string          `gorm:"size:255;not null"`
	PrivateIp      string          `gorm:"size:255;not null"`
	InstanceTypeId int             `gorm:"not null"`
	Status         string          `gorm:"size:255;not null"`
	InstanceType   TblInstanceType `gorm:"foreignKey:InstanceTypeId"`
}

type TblStorage struct {
	Id               int       `gorm:"primaryKey"`
	CreatedDate      time.Time `gorm:"autoCreateTime"`
	UpdatedDate      time.Time `gorm:"autoUpdateTime"`
	Name             string    `gorm:"size:255;not null"`
	SizeMb           int       `gorm:"not null"`
	AwsEbsId         string    `gorm:"size:255"`
	AwsEbsSnapshotId string    `gorm:"size:255"`
	Status           string    `gorm:"size:255;not null"`
}

type TblPersistentNode struct {
	Id               int               `gorm:"primaryKey"`
	CreatedDate      time.Time         `gorm:"autoCreateTime"`
	UpdatedDate      time.Time         `gorm:"autoUpdateTime"`
	Name             string            `gorm:"size:255;not null"`
	OwnerId          int               `gorm:"not null"`
	ServerTemplateId int               `gorm:"not null"`
	InstanceTypeId   int               `gorm:"not null"`
	StorageId        int               `gorm:"not null"`
	Status           string            `gorm:"size:255;not null"`
	Owner            TblUser           `gorm:"foreignKey:OwnerId"`
	ServerTemplate   TblServerTemplate `gorm:"foreignKey:ServerTemplateId"`
	InstanceType     TblInstanceType   `gorm:"foreignKey:InstanceTypeId"`
	Storage          TblStorage        `gorm:"foreignKey:StorageId"`
}
