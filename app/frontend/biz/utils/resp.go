package utils

import (
	"context"
	"gomall/app/cart/kitex_gen/gomall/cart"
	"gomall/app/frontend/infra/rpc_client"
	myutils "gomall/app/frontend/utils"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}

func WrapResp(ctx context.Context, c *app.RequestContext, content map[string]any) map[string]any {
	uid := myutils.GetUidByCtx(ctx)
	content["user_id"] = uid
	iuid, _ := strconv.Atoi(uid)
	if uid != "" {
		cartResp, err := rpc_client.CartClient.GetCart(ctx, &cart.GetCartReq{
			UserId: int32(iuid),
		})
		if err == nil && cartResp != nil {
			content["cart_num"] = len(cartResp.Items)
		}
	}
	return content
}
