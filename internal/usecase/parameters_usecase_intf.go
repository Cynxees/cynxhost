package usecase

import (
	"context"
	"cynxhost/internal/model/entity"
)

type ParametersUseCase interface {
	GetAllParameters(ctx context.Context) (context.Context, []entity.TblParameters, error)
}