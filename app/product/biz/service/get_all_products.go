package service

import (
	"context"
	product "gomall/app/product/kitex_gen/gomall/product"
)

type GetAllProductsService struct {
	ctx context.Context
} // NewGetAllProductsService new GetAllProductsService
func NewGetAllProductsService(ctx context.Context) *GetAllProductsService {
	return &GetAllProductsService{ctx: ctx}
}

// Run create note info
func (s *GetAllProductsService) Run(req *product.GetAllProductsReq) (resp *product.GetAllProductsResp, err error) {
	// Finish your business logic.

	return
}
