package service

import (
	"context"
	"errors"
	"gomall/app/product/biz/model"
	product "gomall/app/product/kitex_gen/gomall/product"
	"gomall/app/product/biz/dal/mysql"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, errors.New("req is is error")
	}
	prodQuery := model.NewProductQuery(s.ctx, mysql.DB)
	prod, err := prodQuery.GetProductById(int(req.Id))
	if err != nil {
		return nil, err
	}
	return &product.GetProductResp{
		Product: &product.Product{
			Id: int32(prod.ID),
			Picture: prod.Picture,
			Name: prod.Name,
			Description: prod.Description,
			Price: prod.Price,
		},
	}, nil

}
