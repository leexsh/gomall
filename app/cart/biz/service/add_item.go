package service

import (
	"context"
	"gomall/app/cart/biz/dal/mysql"
	"gomall/app/cart/biz/model"
	cart "gomall/app/cart/kitex_gen/gomall/cart"
	rpc_client "gomall/app/cart/rpc"
	"gomall/app/product/kitex_gen/gomall/product"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	prods, err := rpc_client.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
		Id: req.Item.ProductId,
	})
	if err != nil {
		return nil, err
	}
	if prods == nil || prods.Product.Id == 0 {
		return nil, err
	}
	ca := &model.Cart{
		UserId: req.UserId,
		ProductId: req.Item.ProductId,
		Qty: req.Item.Quantity,
	}
	err = model.AddItem(s.ctx, mysql.DB, ca)
	if err != nil {
		return nil, err
	}
	resp = cart.NewAddItemResp()
	return 
}
