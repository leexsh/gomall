package service

import (
	"context"

	product "gomall/app/frontend/hertz_gen/frontend/product"
	"gomall/app/frontend/infra/rpc_client"
	productRPC "gomall/rpc_gen/kitex_gen/gomall/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	res, err := rpc_client.ProductClient.SearchProducts(h.Context, &productRPC.SearchProductReq{
		Query: req.Q,
	})
	if err != nil {
		return nil, err
	}
	return utils.H {
		"items": res.Results,
		"q": req.Q,
	}, nil
}
