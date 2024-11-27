package tblinstancetype

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"

	"gorm.io/gorm"
)

type TblInstanceTypeImpl struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) database.TblInstanceType {
	return &TblInstanceTypeImpl{
		DB: DB,
	}
}

func (database *TblInstanceTypeImpl) GetInstanceType(ctx context.Context, key, value string) (context.Context, entity.TblInstanceType, error) {
	var instance entity.TblInstanceType
	err := database.DB.WithContext(ctx).Where(key+" = ?", value).First(&instance).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx, entity.TblInstanceType{}, nil
		}
		return ctx, entity.TblInstanceType{}, err
	}

	return ctx, instance, nil
}

func (database *TblInstanceTypeImpl) GetAllInstanceType(ctx context.Context) (context.Context, []entity.TblInstanceType, error) {
	var instances []entity.TblInstanceType
	err := database.DB.WithContext(ctx).Find(&instances).Error
	if err != nil {
		return ctx, nil, err
	}

	return ctx, instances, nil
}

func (database *TblInstanceTypeImpl) UpdateInstanceType(ctx context.Context, instance entity.TblInstanceType) (context.Context, entity.TblInstanceType, error) {
	err := database.DB.WithContext(ctx).Save(&instance).Error
	if err != nil {
		return ctx, entity.TblInstanceType{}, err
	}

	return ctx, instance, nil
}
