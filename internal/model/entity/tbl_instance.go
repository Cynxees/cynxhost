package entity

import "time"

type TblInstance struct {
	Id             int
	HostTemplateId int
	OwnerId        int
	Status         string
	Address        string
	CreatedDate    time.Time `gorm:"autoCreateTime"`
	ModifiedDate   time.Time `gorm:"autoUpdateTime"`
}
