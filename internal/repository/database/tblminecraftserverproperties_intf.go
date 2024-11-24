package database

import (
	"context"
	"cynxhost/internal/model/entity"
)

type TblMinecraftServerProperties interface {
	InitializeMinecraftServerProperties(ctx context.Context, hostTemplateId int) (context.Context, error)
	CreateMinecraftServerProperties(ctx context.Context, data entity.TblMinecraftServerProperties) (context.Context, entity.TblMinecraftServerProperties, error)

	GetMinecraftServerProperties(ctx context.Context, key, value string) (context.Context, entity.TblMinecraftServerProperties, error)
	GetHostTemplateMinecraftServerProperties(ctx context.Context, hostTemplateId int) (context.Context, []entity.TblMinecraftServerProperties, error)

	UpdateMinecraftServerProperties(ctx context.Context, data entity.TblMinecraftServerProperties) (context.Context, entity.TblMinecraftServerProperties, error)

	DeleteMinecraftServerProperties(ctx context.Context, id int) (context.Context, error)
	DeleteAllByHostTemplateId(ctx context.Context, hostTemplateId int) (context.Context, error)
}
