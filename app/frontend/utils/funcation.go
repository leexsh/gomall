package myutils

import "context"

func GetUidByCtx(ctx context.Context) string {
	userId := ctx.Value(SessionUidKey)
	if userId == nil {
		return ""
	}
	return userId.(string)
}
