package service

import (
	"context"
	"gomall/app/cart/biz/dal/mysql"
	"gomall/app/cart/biz/model"
	cart "gomall/app/cart/kitex_gen/gomall/cart"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	err = model.EmptyItem(s.ctx, mysql.DB, req.UserId)
	resp = cart.NewEmptyCartResp()
	return
}
