// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	frontend_auth "gomall/app/frontend/biz/router/frontend/auth"
	frontend_home "gomall/app/frontend/biz/router/frontend/home"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	frontend_auth.Register(r)

	frontend_home.Register(r)
}