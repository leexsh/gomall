package service

import (
	"context"
	"fmt"
	"strconv"

	"gomall/app/cart/kitex_gen/gomall/cart"
	common "gomall/app/frontend/hertz_gen/frontend/common"
	"gomall/app/frontend/infra/rpc_client"
	myutils "gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/gomall/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutService(Context context.Context, RequestContext *app.RequestContext) *CheckoutService {
	return &CheckoutService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	var items []map[string]string
	userId := myutils.GetUidByCtx(h.Context)
	userI, err := strconv.Atoi(userId)
	carts, err := rpc_client.CartClient.GetCart(h.Context, &cart.GetCartReq{UserId: int32(userI)})
	if err != nil {
		return nil, err
	}
	var price float64
	for _, item := range carts.Items {
		prod, err := rpc_client.ProductClient.GetProduct(h.Context, &product.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if prod.Product == nil {
			continue
		}
		p := prod.Product
		items = append(items, map[string]string{
			"Name": p.Name,
			"Price": fmt.Sprintf("%.2f", p.Price),
			"Picture": p.Picture,
			"Qty": strconv.Itoa(int(item.Quantity)),
		})
		price += prod.Product.Price * float64(item.Quantity)
	}
	return utils.H{
		"title": "Checkout",
		"items": items,
		"total": fmt.Sprintf("%.2f", price),
	}, nil
}
