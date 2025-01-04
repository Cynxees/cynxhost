package entity

import "time"

type TblUser struct {
	Id           int       `gorm:"primaryKey"`
	Username     string    `gorm:"size:255;not null;unique"`
	Password     string    `gorm:"size:255;not null"`
	Coin         int       `gorm:"default:0"`
	CreatedDate  time.Time `gorm:"autoCreateTime"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime"`
}
