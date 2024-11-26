package entity

import "time"

type TblUser struct {
	Id           int       `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Coin         int       `json:"coin"`
	CreatedDate  time.Time `json:"created_date"`
	ModifiedDate time.Time `json:"modified_date"`
}
