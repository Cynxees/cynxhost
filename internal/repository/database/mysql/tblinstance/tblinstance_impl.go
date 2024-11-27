package tblinstance

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"strconv"

	"gorm.io/gorm"
)

type TblInstanceImpl struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) database.TblInstance {
	return &TblInstanceImpl{
		DB: DB,
	}
}

func (database *TblInstanceImpl) CreateInstance(ctx context.Context, instance entity.TblInstance) (context.Context, entity.TblInstance, error) {
	err := database.DB.WithContext(ctx).Create(&instance).Error
	if err != nil {
		return ctx, entity.TblInstance{}, err
	}

	ctx, createdData, err := database.GetInstance(ctx, "id", strconv.Itoa(instance.Id))
	if err != nil {
		return ctx, entity.TblInstance{}, err
	}

	return ctx, createdData, nil
}

func (database *TblInstanceImpl) GetInstance(ctx context.Context, key, value string) (context.Context, entity.TblInstance, error) {
	var instance entity.TblInstance

	err := database.DB.WithContext(ctx).Where(key+" = ?", value).First(&instance).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx, entity.TblInstance{}, nil
		}
		return ctx, entity.TblInstance{}, err
	}

	return ctx, instance, nil
}

func (database *TblInstanceImpl) UpdateInstance(ctx context.Context, instance entity.TblInstance) (context.Context, entity.TblInstance, error) {
	err := database.DB.WithContext(ctx).Save(&instance).Error
	if err != nil {
		return ctx, entity.TblInstance{}, err
	}

	ctx, updatedData, err := database.GetInstance(ctx, "id", strconv.Itoa(instance.Id))
	if err != nil {
		return ctx, entity.TblInstance{}, err
	}

	return ctx, updatedData, nil
}

func (database *TblInstanceImpl) DeleteInstance(ctx context.Context, instanceId int) (context.Context, error) {
	err := database.DB.WithContext(ctx).Delete(&entity.TblInstance{}, instanceId).Error
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}
