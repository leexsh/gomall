package service

import (
	"context"
	"gomall/app/frontend/infra/rpc_client"

	auth "gomall/app/frontend/hertz_gen/frontend/auth"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginRequest) (resp map[string]string, err error) {
	// defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	// }()
	// todo edit your code
	resp, err := rpc_client.UserClient.Log
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	resp = make(map[string]string)
	direct := "redirect"
	resp[direct] = "/"
	if req.Next != "" {
		resp[direct] = req.Next
	}
	return
}
