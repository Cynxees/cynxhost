package amiusecase

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase"
	"strconv"
)

type AmiUseCaseImpl struct {
	tblAmi database.TblAmi
}

func New(tblAmi database.TblAmi) usecase.AmiUseCase {
	return &AmiUseCaseImpl{
		tblAmi: tblAmi,
	}
}

func (usecase *AmiUseCaseImpl) GetAllAmi(ctx context.Context) (context.Context, []entity.TblAmi, error) {
	return usecase.tblAmi.GetAllAmi(ctx)
}

func (usecase *AmiUseCaseImpl) GetAmi(ctx context.Context, amiId int) (context.Context, entity.TblAmi, error) {
	return usecase.tblAmi.GetAmi(ctx, "id", strconv.Itoa(amiId))
}
