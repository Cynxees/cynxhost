package entity

import "time"

type TblHostTemplate struct {
	Id             int
	OwnerId        int
	AmiId          int
	InstanceTypeId int
	CreatedDate    time.Time `gorm:"autoCreateTime"`
	ModifiedDate   time.Time `gorm:"autoUpdateTime"`
}
