package middleware

import (
	"context"
	myutils "gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		s := sessions.Default(ctx)
		c = context.WithValue(c, myutils.SessionUidKey, s.Get(myutils.SessionUidKey))
		ctx.Next(c)
	}
}

func Auth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		s := sessions.Default(ctx)
		userId := s.Get(myutils.SessionUidKey)
		if userId == nil {
			ctx.Redirect(consts.StatusFound, []byte("/sign-in?next="+ctx.FullPath()))
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
