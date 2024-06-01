package service

import (
	"context"
	"errors"
	"gomall/app/user/biz/dal/mysql"
	"gomall/app/user/biz/model"
	"gomall/app/user/kitex_gen/gomall/user"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Password == "" || req.Password == "" {
		return nil, errors.New("email or password is empty")
	}
	u, err := model.GetByEmail(mysql.DB, req.Email)
	if u == nil || err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	resp = &user.LoginResp{UserId: strconv.Itoa(int(u.ID))}
	return
}
