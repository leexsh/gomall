package product

import (
	"context"
	product "gomall/rpc_gen/kitex_gen/gomall/product"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"gomall/rpc_gen/kitex_gen/gomall/product/productservice"
)

type RPCClient interface {
	KitexClient() productservice.Client
	Service() string
	ListProducts(ctx context.Context, req *product.ListProductReq, callOptions ...callopt.Option) (r *product.ListProductResp, err error)
	GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error)
	SearchProducts(ctx context.Context, req *product.SearchProductReq, callOptions ...callopt.Option) (r *product.SearchProductResp, err error)
	GetAllProducts(ctx context.Context, req *product.GetAllProductsReq, callOptions ...callopt.Option) (r *product.GetAllProductsResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ListProducts(ctx context.Context, req *product.ListProductReq, callOptions ...callopt.Option) (r *product.ListProductResp, err error) {
	return c.kitexClient.ListProducts(ctx, req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, req *product.GetProductReq, callOptions ...callopt.Option) (r *product.GetProductResp, err error) {
	return c.kitexClient.GetProduct(ctx, req, callOptions...)
}

func (c *clientImpl) SearchProducts(ctx context.Context, req *product.SearchProductReq, callOptions ...callopt.Option) (r *product.SearchProductResp, err error) {
	return c.kitexClient.SearchProducts(ctx, req, callOptions...)
}

func (c *clientImpl) GetAllProducts(ctx context.Context, req *product.GetAllProductsReq, callOptions ...callopt.Option) (r *product.GetAllProductsResp, err error) {
	return c.kitexClient.GetAllProducts(ctx, req, callOptions...)
}
