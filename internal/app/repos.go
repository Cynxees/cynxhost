package app

import (
	"cynxhost/internal/dependencies"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/repository/database/tblscript"
	"cynxhost/internal/repository/database/tblservertemplate"
	"cynxhost/internal/repository/database/tbluser"
)

type Repos struct {
	TblUser           database.TblUser
	TblScript         database.TblScript
	TblServerTemplate database.TblServerTemplate
	JWTManager        *dependencies.JWTManager
}

func NewRepos(dependencies *Dependencies) *Repos {

	return &Repos{
		TblUser:           tbluser.New(dependencies.DatabaseClient.Db),
		TblScript:         tblscript.New(dependencies.DatabaseClient.Db),
		TblServerTemplate: tblservertemplate.New(dependencies.DatabaseClient.Db),
		JWTManager:        dependencies.JWTManager,
	}
}
