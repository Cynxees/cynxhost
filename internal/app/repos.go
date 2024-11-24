package app

import (
	"cynxhost/internal/dependencies"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/repository/database/mysql/tblami"
	"cynxhost/internal/repository/database/mysql/tblhosttemplate"
	"cynxhost/internal/repository/database/mysql/tblinstance"
	"cynxhost/internal/repository/database/mysql/tblinstancetype"
	"cynxhost/internal/repository/database/mysql/tblminecraftserverproperties"
	"cynxhost/internal/repository/database/mysql/tblparameters"
	"cynxhost/internal/repository/database/mysql/tbluser"
)

type Repos struct {
	TblUser                      database.TblUser
	TblInstance                  database.TblInstance
	TblHostTemplate              database.TblHostTemplate
	TblInstanceType              database.TblInstanceType
	TblParameters                database.TblParameters
	TblAmi                       database.TblAmi
	TblMinecraftServerProperties database.TblMinecraftServerProperties
	JWTManager                   *dependencies.JWTManager
}

func NewRepos(dependencies *Dependencies) *Repos {

	return &Repos{
		TblUser:                      tbluser.New(dependencies.DatabaseClient.Db),
		TblInstance:                  tblinstance.New(dependencies.DatabaseClient.Db),
		TblHostTemplate:              tblhosttemplate.New(dependencies.DatabaseClient.Db),
		TblInstanceType:              tblinstancetype.New(dependencies.DatabaseClient.Db),
		TblParameters:                tblparameters.New(dependencies.DatabaseClient.Db),
		TblAmi:                       tblami.New(dependencies.DatabaseClient.Db),
		TblMinecraftServerProperties: tblminecraftserverproperties.New(dependencies.DatabaseClient.Db),
		JWTManager:                   dependencies.JWTManager,
	}
}
