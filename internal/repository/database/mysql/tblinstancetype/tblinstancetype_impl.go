package tblinstancetype

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"database/sql"
)

type TblInstanceTypeImpl struct {
	DB *sql.DB
}

func New(DB *sql.DB) database.TblInstanceType {
	return &TblInstanceTypeImpl{
		DB: DB,
	}
}

func (database *TblInstanceTypeImpl) GetInstanceType(ctx context.Context, key, value string) (context.Context, entity.TblInstanceType, error) {
	// Implementation for getting an instance type by key and value
	return nil, entity.TblInstanceType{}, nil
}

func (database *TblInstanceTypeImpl) GetAllInstanceType(ctx context.Context) (context.Context, []entity.TblInstanceType, error) {
	// Implementation for getting all instance types
	return nil, []entity.TblInstanceType{}, nil
}

func (database *TblInstanceTypeImpl) UpdateInstanceType(ctx context.Context, instance entity.TblInstanceType) (context.Context, entity.TblInstanceType, error) {
	// Implementation for updating an instance type
	return nil, entity.TblInstanceType{}, nil
}