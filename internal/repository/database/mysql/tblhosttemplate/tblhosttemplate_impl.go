package tblhosttemplate

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"strconv"

	"gorm.io/gorm"
)

type TblHostTemplateImpl struct {
	DB *gorm.DB
}

func New(DB *gorm.DB) database.TblHostTemplate {
	return &TblHostTemplateImpl{
		DB: DB,
	}
}

func (database *TblHostTemplateImpl) CreateHostTemplate(ctx context.Context, hostTemplate entity.TblHostTemplate) (context.Context, entity.TblHostTemplate, error) {
	err := database.DB.WithContext(ctx).Create(&hostTemplate).Error
	if err != nil {
		return ctx, entity.TblHostTemplate{}, err
	}

	ctx, createdData, err := database.GetHostTemplate(ctx, "id", strconv.Itoa(hostTemplate.Id))
	if err != nil {
		return ctx, entity.TblHostTemplate{}, err
	}

	return ctx, createdData, nil
}

func (database *TblHostTemplateImpl) GetHostTemplate(ctx context.Context, key, value string) (context.Context, entity.TblHostTemplate, error) {
	var hostTemplate entity.TblHostTemplate

	err := database.DB.WithContext(ctx).Where(key+" = ?", value).First(&hostTemplate).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx, entity.TblHostTemplate{}, nil
		}
		return ctx, entity.TblHostTemplate{}, err
	}

	return ctx, hostTemplate, nil
}

func (database *TblHostTemplateImpl) GetAllUserOwnedHostTemplate(ctx context.Context, userId int) (context.Context, []entity.TblHostTemplate, error) {
	var hostTemplates []entity.TblHostTemplate

	err := database.DB.WithContext(ctx).Where("owner_id = ?", userId).Find(&hostTemplates).Error
	if err != nil {
		return ctx, nil, err
	}

	return ctx, hostTemplates, nil
}

func (database *TblHostTemplateImpl) UpdateHostTemplate(ctx context.Context, hostTemplate entity.TblHostTemplate) (context.Context, entity.TblHostTemplate, error) {
	err := database.DB.WithContext(ctx).Save(&hostTemplate).Error
	if err != nil {
		return ctx, entity.TblHostTemplate{}, err
	}

	ctx, finalData, err := database.GetHostTemplate(ctx, "id", strconv.Itoa(hostTemplate.Id))
	if err != nil {
		return ctx, entity.TblHostTemplate{}, err
	}

	return ctx, finalData, nil
}

func (database *TblHostTemplateImpl) DeleteHostTemplate(ctx context.Context, hostTemplateId int) (context.Context, error) {
	err := database.DB.WithContext(ctx).Delete(&entity.TblHostTemplate{}, hostTemplateId).Error
	if err != nil {
		return ctx, err
	}

	return ctx, nil
}
