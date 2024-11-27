package tblami

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"

	"gorm.io/gorm"
)

type TblAmiImpl struct {
	DB *gorm.DB
}

// Constructor function for TblAmiImpl
func New(DB *gorm.DB) database.TblAmi {
	return &TblAmiImpl{
		DB: DB,
	}
}

// GetAmi retrieves a single AMI record based on a dynamic key-value filter
func (database *TblAmiImpl) GetAmi(ctx context.Context, key, value string) (context.Context, entity.TblAmi, error) {
	var ami entity.TblAmi

	// Use GORM's dynamic filter capabilities
	err := database.DB.WithContext(ctx).Where(key+" = ?", value).First(&ami).Error
	if err != nil {
		return ctx, entity.TblAmi{}, err
	}

	return ctx, ami, nil
}

// GetAllAmi retrieves all AMI records
func (database *TblAmiImpl) GetAllAmi(ctx context.Context) (context.Context, []entity.TblAmi, error) {
	var amis []entity.TblAmi

	// Use GORM's Find method to retrieve all rows
	err := database.DB.WithContext(ctx).Find(&amis).Error
	if err != nil {
		return ctx, nil, err
	}

	return ctx, amis, nil
}
