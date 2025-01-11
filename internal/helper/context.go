package helper

import (
	"context"
	"cynxhost/internal/constant"
	contextmodel "cynxhost/internal/model/context"
)

func GetUserFromContext(ctx context.Context) (contextmodel.User, bool) {
	user, ok := ctx.Value(constant.ContextKeyUser).(contextmodel.User)
	return user, ok
}

func GetVisibilityLevelFromContext(ctx context.Context) (int, bool) {
	level, ok := ctx.Value(constant.ContextKeyVisibility).(constant.VisibilityLevel)
	return int(level), ok
}

func SetVisibilityLevelToContext(ctx context.Context, level constant.VisibilityLevel) context.Context {
	return context.WithValue(ctx, constant.ContextKeyVisibility, level)
}
