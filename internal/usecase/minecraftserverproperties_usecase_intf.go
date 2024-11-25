package usecase

import (
	"context"
	"cynxhost/internal/model/entity"
)

type MinecraftServerPropertiesUseCase interface {
	GetHostTemplateMinecraftServerProperties(ctx context.Context, hostTemplateId int) (context.Context, entity.TblMinecraftServerProperties, error)
	
	InitializeMinecraftServerProperties(ctx context.Context, hostTemplateId int) (context.Context, error)
	CreateMinecraftServerProperties(ctx context.Context, data entity.TblMinecraftServerProperties) (context.Context, entity.TblMinecraftServerProperties, error)
	
	UpdateMinecraftServerProperties(ctx context.Context, data entity.TblMinecraftServerProperties) (context.Context, entity.TblMinecraftServerProperties, error)
	UpdateMultipleMinecraftServerProperties(ctx context.Context, data []entity.TblMinecraftServerProperties) (context.Context, []entity.TblMinecraftServerProperties, error)

	DeleteMinecraftServerProperties(ctx context.Context, id int) (context.Context, error)
	DeleteAllByHostTemplateId(ctx context.Context, hostTemplateId int) (context.Context, error)
}