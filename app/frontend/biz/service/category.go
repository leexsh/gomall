package service

import (
	"context"

	product "gomall/app/frontend/hertz_gen/frontend/product"
	"gomall/app/frontend/infra/rpc_client"
	productRPC "gomall/rpc_gen/kitex_gen/gomall/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *product.CategoryReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	res, err := rpc_client.ProductClient.ListProducts(h.Context, &productRPC.ListProductReq{
		CategoryName: req.Category,
	})
	if err != nil {
		return nil, err
	}
	return utils.H {
		"title": "Category",
		"items": res.Products,
	}, nil
}
