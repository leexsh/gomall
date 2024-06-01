package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

const SessionUidKey string = "user_id"

func GlobalAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		s := sessions.Default(ctx)
		c = context.WithValue(c, SessionUidKey, s.Get(SessionUidKey))
		ctx.Next(c)
	}
}

func Auth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		s := sessions.Default(ctx)
		userId := s.Get(SessionUidKey)
		if userId == nil {
			ctx.Redirect(consts.StatusFound, []byte("/sign-in?next="+ctx.FullPath()))
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
