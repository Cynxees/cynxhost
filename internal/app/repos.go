package app

import (
	"cynxhost/internal/dependencies"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/repository/database/mysql/tbluser"
)

type Repos struct {
	TblUser    database.TblUser
	JWTManager *dependencies.JWTManager
}

func NewRepos(dependencies *Dependencies) *Repos {

	return &Repos{
		TblUser:    tbluser.New(dependencies.DatabaseClient.Db),
		JWTManager: dependencies.JWTManager,
	}
}
