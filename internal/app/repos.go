package app

import (
	"cynxhost/internal/dependencies"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/repository/database/tblinstance"
	"cynxhost/internal/repository/database/tblinstancetype"
	"cynxhost/internal/repository/database/tblparameter"
	"cynxhost/internal/repository/database/tblpersistentnode"
	"cynxhost/internal/repository/database/tblscript"
	"cynxhost/internal/repository/database/tblservertemplate"
	"cynxhost/internal/repository/database/tblstorage"
	"cynxhost/internal/repository/database/tbluser"
)

type Repos struct {
	TblUser           database.TblUser
	TblScript         database.TblScript
	TblServerTemplate database.TblServerTemplate
	TblInstance       database.TblInstance
	TblInstanceType   database.TblInstanceType
	TblPersistentNode database.TblPersistentNode
	TblStorage        database.TblStorage
	TblParameter      database.TblParameter
	JWTManager        *dependencies.JWTManager
}

func NewRepos(dependencies *Dependencies) *Repos {

	return &Repos{
		TblUser:           tbluser.New(dependencies.DatabaseClient.Db),
		TblScript:         tblscript.New(dependencies.DatabaseClient.Db),
		TblServerTemplate: tblservertemplate.New(dependencies.DatabaseClient.Db),
		TblInstance:       tblinstance.New(dependencies.DatabaseClient.Db),
		TblInstanceType:   tblinstancetype.New(dependencies.DatabaseClient.Db),
		TblPersistentNode: tblpersistentnode.New(dependencies.DatabaseClient.Db),
		TblStorage:        tblstorage.New(dependencies.DatabaseClient.Db),
		TblParameter:      tblparameter.New(dependencies.DatabaseClient.Db),
		JWTManager:        dependencies.JWTManager,
	}
}
