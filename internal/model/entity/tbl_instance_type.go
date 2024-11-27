package entity

import "time"

type TblInstanceType struct {
	Id           int
	VcpuCount    int
	MemoryCount  int
	SpotPrice    int
	SellPrice    int
	Name         string
	Status       string
	CreatedDate  time.Time `gorm:"autoCreateTime"`
	ModifiedDate time.Time `gorm:"autoUpdateTime"`
}
