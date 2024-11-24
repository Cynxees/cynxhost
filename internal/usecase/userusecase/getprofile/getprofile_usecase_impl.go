package getprofile

import (
	"context"
	"cynxhost/internal/model/entity"
	"cynxhost/internal/repository/database"
	"cynxhost/internal/usecase/userusecase"
	"strconv"
)

type getProfileUseCaseImpl struct {
	tblUser database.TblUser
}

func NewGetProfileUseCase(tblUser database.TblUser) userusecase.GetProfileUseCase {
	return &getProfileUseCaseImpl{
		tblUser: tblUser,
	}
}

func (uc *getProfileUseCaseImpl) GetProfile(ctx context.Context, userId int) (context.Context, entity.TblUser, error) {
	ctx, user, err := uc.tblUser.GetUser(ctx, "id", strconv.Itoa(userId))
	if err != nil {
		return ctx, entity.TblUser{}, err
	}

	return ctx, user, nil
}
