package tblparameters

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"

	"gorm.io/gorm"
)

type TblParametersImpl struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) database.TblParameters {
	return &TblParametersImpl{
		DB: DB,
	}
}

func (database *TblParametersImpl) GetAllParameters(ctx context.Context) (context.Context, []entity.TblParameters, error) {
	var parameters []entity.TblParameters
	err := database.DB.WithContext(ctx).Find(&parameters).Error
	if err != nil {
		return ctx, nil, err
	}

	return ctx, parameters, nil
}

func (database *TblParametersImpl) GetParameters(ctx context.Context, key, value string) (context.Context, entity.TblParameters, error) {
	var param entity.TblParameters
	err := database.DB.WithContext(ctx).Where(key+" = ?", value).First(&param).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx, entity.TblParameters{}, nil
		}
		return ctx, entity.TblParameters{}, err
	}

	return ctx, param, nil
}
