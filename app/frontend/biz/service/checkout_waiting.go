package service

import (
	"context"
	"strconv"

	checkout "gomall/app/frontend/hertz_gen/frontend/checkout"
	"gomall/app/frontend/infra/rpc_client"
	myutils "gomall/app/frontend/utils"
	rpccheckout "gomall/rpc_gen/kitex_gen/gomall/checkout"
	rpcpayment "gomall/rpc_gen/kitex_gen/gomall/payment"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CheckoutWaitingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckoutWaitingService(Context context.Context, RequestContext *app.RequestContext) *CheckoutWaitingService {
	return &CheckoutWaitingService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckoutWaitingService) Run(req *checkout.CheckoutReq) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := myutils.GetUidByCtx(h.Context)
	userI, err := strconv.Atoi(userId)
	_, err = rpc_client.CheckoutClient.Checkout(h.Context, &rpccheckout.CheckoutReq{
		UserId: int32(userI),
		Email: req.Email,
		FirstName: req.FirstName,
		LastName: req.LastName,
		Address: &rpccheckout.Address{
			Country: req.Country,
			City: req.City,
			ZipCode: req.ZipCode,
			State: req.Province,
			StreetAddress: req.Street,
		},
		CreditCard: &rpcpayment.CreditCardInfo{
			CreditCardNumber: req.CardNum,
			CreditCardCvv: req.Cvv,
			CreditCardExpirationYear: req.ExpirationYear,
			CreditCardExpirationMonth: req.ExpirationMonth,
		},
	})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"title":"waiting",
		"redirect":"/checkout/result",
	}, nil
}
