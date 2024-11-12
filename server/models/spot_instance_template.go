package models

import (
	"time"

	"gorm.io/gorm"
)

type SpotInstanceTemplate struct {
	gorm.Model
	Id        int `gorm:"uniqueIndex"`
	FleetRequestId *string
	InstanceId *string
	UserId    int
	
	Name      string
	Status string
	InstanceType string

	AmiId string
	IpId uint64
	Region string

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
