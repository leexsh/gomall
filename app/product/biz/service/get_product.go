package service

import (
	"context"
	"errors"
	"gomall/app/product/biz/dal/mysql"
	"gomall/app/product/biz/model"
	product "gomall/app/product/kitex_gen/gomall/product"

	"github.com/redis/go-redis/v9"
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
	proquery := model.NewCacheProdcutQuery(s.ctx, mysql.DB, redis.RedisClient)
	prod, err := proquery.GetById(int(req.Id))
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
