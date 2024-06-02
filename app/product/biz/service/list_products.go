package service

import (
	"context"
	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/biz/model"
	product "gomall/app/product/kitex_gen/gomall/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductReq) (resp *product.ListProductResp, err error) {
	// Finish your business logic.
	cateQuery := model.NewCategoryQuery(s.ctx, mysql.DB)
	c, err := cateQuery.GetProductByCategoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}
	resp = product.NewListProductResp()
	for _, v := range c {
		for _, v1 := range v.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id: int32(v1.ID),
				Name: v1.Name,
				Picture: v1.Picture,
				Descrption: v1.Description,
			})
		}
	}
	return
}
