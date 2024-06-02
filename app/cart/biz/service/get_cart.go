package service

import (
	"context"
	"gomall/app/cart/biz/dal/mysql"
	"gomall/app/cart/biz/model"
	cart "gomall/app/cart/kitex_gen/gomall/cart"
)

type GetCartService struct {
	ctx context.Context
} // NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	// Finish your business logic.
	res, err := model.GetItemByUserID(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, err
	}
	var ites []*cart.CartItem
	for _, v := range res {
		ites = append(ites, &cart.CartItem{
			ProductId: v.ProductId,
			Quantity: v.Qty,
		})
	}
	resp = cart.NewGetCartResp()
	resp.Items = ites
	return
}
