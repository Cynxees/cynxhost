package tblparameters

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"database/sql"
)

type TblParametersImpl struct {
	DB *sql.DB
}

func New(DB *sql.DB) database.TblParameters {
	return &TblParametersImpl{
		DB: DB,
	}
}

func (database *TblParametersImpl) GetAllParameters(ctx context.Context) (context.Context, []entity.TblParameters, error) {
	// Implementation for getting all parameters
	return nil, nil, nil
}

func (database *TblParametersImpl) GetParameters(ctx context.Context, key, value string) (context.Context, entity.TblParameters, error) {
	// Implementation for getting parameters by key and value
	return nil, entity.TblParameters{}, nil
}