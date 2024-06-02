package service

import (
	"context"
	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/biz/model"
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
	its, err := model.NewProductQuery(s.ctx, mysql.DB).GetAllProducts()
	if err != nil {
		return nil, err
	}
	resp = product.NewGetAllProductsResp()
	for _, i := range its {
		resp.Results = append(resp.Results, &product.Product{
			Id:          int32(i.ID),
			Name:        i.Name,
			Description: i.Description,
			Picture:     i.Picture,
			Price:       i.Price,
		})
	}
	return
}
