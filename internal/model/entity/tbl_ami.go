package entity

import "time"

type TblAmi struct {
	Id               int
	Name             string
	ImageId          string
	MinecraftEdition string
	MinecraftVersion string
	ModLoader        string
	ModLoaderVersion string
	MinimumRam       int
	MinimumVcpu      int
	CreatedDate      time.Time `gorm:"autoCreateTime"`
	ModifiedDate     time.Time `gorm:"autoUpdateTime"`
}
