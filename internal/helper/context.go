package helper

import (
	"context"
	"cynxhost/internal/constant"
	contextmodel "cynxhost/internal/model/context"
)

func GetUserFromContext(ctx context.Context) (contextmodel.User, bool) {
	user, ok := ctx.Value(constant.UserContextKey).(contextmodel.User)
	return user, ok
}
