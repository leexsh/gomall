package service

import (
	"context"
	"strconv"

	"gomall/app/cart/kitex_gen/gomall/cart"
	common "gomall/app/frontend/hertz_gen/frontend/common"
	"gomall/app/frontend/infra/rpc_client"
	myutils "gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/gomall/product"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	id, _ := strconv.Atoi(myutils.GetUidByCtx(h.Context))
	res, err := rpc_client.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: int32(id),
	})	
	if err != nil {
		return nil, err
	}
	var items []map[string]string
	var total float64
	for _, v := range res.Items {
		prod, err := rpc_client.ProductClient.GetProduct(h.Context, &product.GetProductReq{Id: v.ProductId})
		if err != nil {
			return nil, err
		}
		p := prod.Product
		items = append(items, map[string]string{
			"Name":p.Name,
			"Description":p.Description,
			"Price": strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Picture":p.Picture,
			"Qty":strconv.Itoa(int(v.Quantity)),
		})
		total += float64(v.Quantity) * p.Price
	}
	return utils.H{
	"title": "Cart",
	"items": items,
	"total": strconv.FormatFloat(total, 'f', 2, 64),
}, nil
}
