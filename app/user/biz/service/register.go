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

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" || req.PasswordConfrim == "" {
		return nil, errors.New("email or password is empty")
	}
	if req.Password != req.PasswordConfrim {
		return nil, err
	}
	passwdHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	resp = user.NewRegisterResp()
	user := &model.User{
		Email:          req.Email,
		PasswordHashed: string(passwdHashed),
	}
	model.CreateUser(mysql.DB, user)
	if err != nil {
		return nil, err
	}

	resp.UserId = strconv.Itoa(int(user.ID))
	return
}
