package cart

import (
	"context"
	"gomall/app/frontend/biz/utils"

	"gomall/app/frontend/biz/service"
	cart "gomall/app/frontend/hertz_gen/frontend/cart"
	common "gomall/app/frontend/hertz_gen/frontend/common"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// GetCart .
// @router /cart [GET]
func GetCart(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewGetCartService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}
	c.HTML(consts.StatusOK, "cart", utils.WrapResp(ctx, c, resp))
}

// AddCartItem .
// @router /cart [POST]
func AddCartItem(ctx context.Context, c *app.RequestContext) {
	var err error
	var req cart.AddCartItemReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	_, err = service.NewAddCartItemService(ctx, c).Run(&req)

	if err != nil {
		c.HTML(consts.StatusOK, "cart", err)
		return
	}
	c.Redirect(consts.StatusFound, []byte("/cart"))
}
