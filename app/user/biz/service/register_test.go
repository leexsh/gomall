package service

import (
	"context"
	"gomall/app/user/biz/dal/mysql"
	"gomall/app/user/kitex_gen/gomall/user"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "t1@qq.com",
		Password:        "wqeqgag",
		PasswordConfrim: "wqeqgag",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
