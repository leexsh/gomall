package service

import (
	"context"
	"strconv"

	rpccart "gomall/app/cart/kitex_gen/gomall/cart"
	cart "gomall/app/frontend/hertz_gen/frontend/cart"
	common "gomall/app/frontend/hertz_gen/frontend/common"
	"gomall/app/frontend/infra/rpc_client"
	myutils "gomall/app/frontend/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type AddCartItemService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddCartItemService(Context context.Context, RequestContext *app.RequestContext) *AddCartItemService {
	return &AddCartItemService{RequestContext: RequestContext, Context: Context}
}

func (h *AddCartItemService) Run(req *cart.AddCartItemReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	id, _ := strconv.Atoi(myutils.GetUidByCtx(h.Context))
	_, err = rpc_client.CartClient.AddItem(h.Context, &rpccart.AddItemReq{
		UserId: int32(id),
		Item: &rpccart.CartItem{
			ProductId: req.ProductID,
			Quantity: req.ProductNum,
		},
	})
	if err != nil {
		return nil, err
	}
	return
}
