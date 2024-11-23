package entity

import "time"

type TblUser struct {
	Id           int
	Username     string
	Password     string
	CreatedDate  time.Time
	ModifiedDate time.Time
}
