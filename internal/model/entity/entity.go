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
