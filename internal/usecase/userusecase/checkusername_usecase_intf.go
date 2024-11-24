package userusecase

import "context"

type CheckUsernameUseCase interface {
	CheckUsername(ctx context.Context, username string) (context.Context, error)
}