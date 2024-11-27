package tblminecraftserverproperties

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"

	"gorm.io/gorm"
)

type TblMinecraftServerPropertiesImpl struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) database.TblMinecraftServerProperties {
	return &TblMinecraftServerPropertiesImpl{
		DB: DB,
	}
}

func (database *TblMinecraftServerPropertiesImpl) InitializeMinecraftServerProperties(ctx context.Context, hostTemplateId int) (context.Context, error) {
	// TODO: implement this
	// properties := []entity.TblMinecraftServerProperties{
	// 	{HostTemplateId: hostTemplateId, Name: "property1", Value: "value1"},
	// 	{HostTemplateId: hostTemplateId, Name: "property2", Value: "value2"},
	// 	{HostTemplateId: hostTemplateId, Name: "property3", Value: "value3"},
	// }

	// err := database.DB.WithContext(ctx).Create(&properties).Error
	// if err != nil {
	// 	return ctx, err
	// }

	return ctx, nil
}

func (database *TblMinecraftServerPropertiesImpl) CreateMinecraftServerProperties(ctx context.Context, data entity.TblMinecraftServerProperties) (context.Context, entity.TblMinecraftServerProperties, error) {
	err := database.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return ctx, entity.TblMinecraftServerProperties{}, err
	}

	return ctx, data, nil
}

func (database *TblMinecraftServerPropertiesImpl) GetMinecraftServerProperties(ctx context.Context, key, value string) (context.Context, entity.TblMinecraftServerProperties, error) {
	var properties entity.TblMinecraftServerProperties

	err := database.DB.WithContext(ctx).Where(key+" = ?", value).First(&properties).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx, entity.TblMinecraftServerProperties{}, nil
		}
		return ctx, entity.TblMinecraftServerProperties{}, err
	}

	return ctx, properties, nil
}

func (database *TblMinecraftServerPropertiesImpl) GetHostTemplateMinecraftServerProperties(ctx context.Context, hostTemplateId int) (context.Context, []entity.TblMinecraftServerProperties, error) {
	var propertiesList []entity.TblMinecraftServerProperties

	err := database.DB.WithContext(ctx).Where("host_template_id = ?", hostTemplateId).Find(&propertiesList).Error
	if err != nil {
		return ctx, nil, err
	}

	return ctx, propertiesList, nil
}

func (database *TblMinecraftServerPropertiesImpl) UpdateMinecraftServerProperties(ctx context.Context, data entity.TblMinecraftServerProperties) (context.Context, entity.TblMinecraftServerProperties, error) {
	err := database.DB.WithContext(ctx).Save(&data).Error
	if err != nil {
		return ctx, entity.TblMinecraftServerProperties{}, err
	}

	return ctx, data, nil
}

func (database *TblMinecraftServerPropertiesImpl) UpdateMultipleMinecraftServerProperties(ctx context.Context, data []entity.TblMinecraftServerProperties) (context.Context, []entity.TblMinecraftServerProperties, error) {
	tx := database.DB.WithContext(ctx).Begin()

	for _, prop := range data {
		err := tx.Save(&prop).Error
		if err != nil {
			tx.Rollback()
			return ctx, nil, err
		}
	}

	err := tx.Commit().Error
	if err != nil {
		return ctx, nil, err
	}

	return ctx, data, nil
}

func (database *TblMinecraftServerPropertiesImpl) DeleteMinecraftServerProperties(ctx context.Context, id int) (context.Context, error) {
	err := database.DB.WithContext(ctx).Delete(&entity.TblMinecraftServerProperties{}, id).Error
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}

func (database *TblMinecraftServerPropertiesImpl) DeleteAllByHostTemplateId(ctx context.Context, hostTemplateId int) (context.Context, error) {
	err := database.DB.WithContext(ctx).Where("host_template_id = ?", hostTemplateId).Delete(&entity.TblMinecraftServerProperties{}).Error
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}
