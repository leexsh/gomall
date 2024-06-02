package service

import (
	"context"

	common "gomall/app/frontend/hertz_gen/frontend/common"
	"gomall/app/frontend/infra/rpc_client"
	"gomall/rpc_gen/kitex_gen/gomall/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	// defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	// }()
	// todo edit your code
	// resp := make(map[string]any)
	// items := []map[string]any{
	// 	{"Name": "T-shirt-1", "Price": 20, "Picture": "/static/image/t-shrit1.png"},
	// 	{"Name": "T-shirt-2", "Price": 20, "Picture": "/static/image/t-shrit1.png"},
	// 	{"Name": "T-shirt-3", "Price": 20, "Picture": "/static/image/t-shrit1.png"},
	// 	{"Name": "T-shirt-4", "Price": 25, "Picture": "/static/image/t-shrit2.png"},
	// 	{"Name": "T-shirt-5", "Price": 25, "Picture": "/static/image/t-shrit2.png"},
	// 	{"Name": "T-shirt-6", "Price": 25, "Picture": "/static/image/t-shrit2.png"},
	// }
	// // pros, err := rpc_client.ProductClient.ListProducts(h.Context, &product.ListProductReq{})
	// resp["Title"] = "Hot Sales"
	// resp["Items"] = items
	ros, err := rpc_client.ProductClient.GetAllProducts(h.Context, &product.GetAllProductsReq{})
	if err != nil {
		klog.Error(err)
	}
	var cartNum int
	return utils.H{
		"Title":    "Hot sale",
		"cart_num": cartNum,
		"Items":    ros.Results,
	}, nil
}
