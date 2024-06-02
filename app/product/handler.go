package main

import (
	"context"
	"gomall/app/product/biz/service"
	product "gomall/app/product/kitex_gen/gomall/product"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// ListProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ListProducts(ctx context.Context, req *product.ListProductReq) (resp *product.ListProductResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductReq) (resp *product.SearchProductResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}
