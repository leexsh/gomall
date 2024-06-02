package service

import (
	"context"

	product "gomall/app/frontend/hertz_gen/frontend/product"
	"gomall/app/frontend/infra/rpc_client"
	productRPC "gomall/rpc_gen/kitex_gen/gomall/product"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code

	res, err := rpc_client.ProductClient.GetProduct(h.Context, &productRPC.GetProductReq{
		Id: req.ID,
	})
	if err != nil {
		return nil,err
	}
	return utils.H{
		"item":res.Product,
	}, nil
}
