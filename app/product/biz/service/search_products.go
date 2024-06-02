package service

import (
	"context"
	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/biz/model"
	product "gomall/app/product/kitex_gen/gomall/product"
)

type SearchProductsService struct {
	ctx context.Context
} // NewSearchProductsService new SearchProductsService
func NewSearchProductsService(ctx context.Context) *SearchProductsService {
	return &SearchProductsService{ctx: ctx}
}

// Run create note info
func (s *SearchProductsService) Run(req *product.SearchProductReq) (resp *product.SearchProductResp, err error) {
	// Finish your business logic.
	pquery := model.NewProductQuery(s.ctx, mysql.DB)
	ps, err := pquery.SearchProducts(req.Query)
	var res []*product.Product
	for _, v := range ps {
		res = append(res, &product.Product{
			Id: int32(v.ID),
			Name: v.Name,
			Picture: v.Picture,
			Descrption: v.Description,
		})
	}
	resp = product.NewSearchProductResp()
	resp.Results = res
	return 
}
