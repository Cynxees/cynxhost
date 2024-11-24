package tblami

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"database/sql"
)

type TblAmiImpl struct {
	DB *sql.DB
}

func New(DB *sql.DB) database.TblAmi {
	return &TblAmiImpl{
		DB: DB,
	}
}

func (database *TblAmiImpl) GetAmi(ctx context.Context, key, value string) (context.Context, entity.TblAmi, error) {
	return nil, entity.TblAmi{}, nil
}

func (database *TblAmiImpl) GetAllAmi(ctx context.Context) ([]entity.TblAmi, error) {
	return nil, nil
}
