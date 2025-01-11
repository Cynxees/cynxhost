package helper

import (
	"context"
	"cynxhost/internal/constant/types"
	contextmodel "cynxhost/internal/model/context"
)

func GetUserFromContext(ctx context.Context) (contextmodel.User, bool) {
	user, ok := ctx.Value(types.ContextKeyUser).(contextmodel.User)
	return user, ok
}

func GetVisibilityLevelFromContext(ctx context.Context) (int, bool) {
	level, ok := ctx.Value(types.ContextKeyVisibility).(types.VisibilityLevel)
	return int(level), ok
}

func SetVisibilityLevelToContext(ctx context.Context, level types.VisibilityLevel) context.Context {
	return context.WithValue(ctx, types.ContextKeyVisibility, level)
}
