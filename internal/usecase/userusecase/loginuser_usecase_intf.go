package userusecase

import "context"

type LoginUserUseCase interface {
	LoginUser(ctx context.Context, username string, password string) (context.Context, string, error)
}