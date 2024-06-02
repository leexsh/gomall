package product

import (
	"context"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
	product "gomall/rpc_gen/kitex_gen/gomall/product"
)

func ListProducts(ctx context.Context, req *product.ListProductReq, callOptions ...callopt.Option) (resp *product.ListProductResp, err error) {
	resp, err = defaultClient.ListProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (resp *product.GetProductResp, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *product.SearchProductReq, callOptions ...callopt.Option) (resp *product.SearchProductResp, err error) {
	resp, err = defaultClient.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
