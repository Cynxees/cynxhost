package app

import (
	"cynxhost/internal/repository/database"
	"cynxhost/internal/repository/database/mysql/tbluser"
)

type Repos struct {
	TblUser database.TblUser
	TblInstance database.TblInstance
}

func NewRepos(dependencies *Dependencies) *Repos {

	tblUser := tbluser.New(dependencies.DatabaseClient.Db)

	return &Repos{
		TblUser: tblUser,
		TblInstance: nil,
	}
}
