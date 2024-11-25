package usecase

import (
	"context"
	"cynxhost/internal/model/entity"
)

type AmiUseCase interface {
	GetAllAmi(ctx context.Context) (context.Context, []entity.TblAmi, error)
	GetAmi(ctx context.Context, amiId int) (context.Context, entity.TblAmi, error)
}
