package service

import (
	"context"
	"encoding/json"
	"gomall/app/cart/kitex_gen/gomall/cart"
	"gomall/app/checkout/infra/consumer/mq"
	"gomall/app/checkout/infra/rpc"
	checkout "gomall/app/checkout/kitex_gen/gomall/checkout"
	"gomall/app/email/kitex_gen/frontend/email"
	"gomall/app/payment/kitex_gen/gomall/payment"
	"gomall/app/product/kitex_gen/gomall/product"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

type CheckoutService struct {
	ctx context.Context
} // NewCheckoutService new CheckoutService
func NewCheckoutService(ctx context.Context) *CheckoutService {
	return &CheckoutService{ctx: ctx}
}

// Run create note info
func (s *CheckoutService) Run(req *checkout.CheckoutReq) (resp *checkout.CheckoutResp, err error) {
	// Finish your business logic.
	// 1. 获取用户购物车
	cartRes, err := rpc.CartClient.GetCart(s.ctx, &cart.GetCartReq{
		UserId: req.UserId,
	})
	if err != nil || cartRes.Items == nil || len(cartRes.Items) == 0{
		return nil, err
	}
	var totalPrice float64
	// 2. 计算商品总价
	for _, item := range cartRes.Items {
		prod, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{
			Id: item.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if prod.Product == nil {
			continue
		}

		totalPrice += prod.Product.Price * float64(item.Quantity)
	}
	// 3.生产订单id
	var orderId string
	u, _ := uuid.NewRandom()
	orderId = u.String()
	payReq := &payment.ChargeReq{
		UserId: req.UserId,
		OrderId: orderId,
		Amount: totalPrice,
		CreditInfo: &payment.CreditCardInfo{
			CreditCardNumber: req.CreditCard.CreditCardNumber,
			CreditCardCvv: req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear: req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
	}
	// 4. 清空购物车
	_, err = rpc.CartClient.EmptyCart(s.ctx, &cart.EmptyCartReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}
	
	// 5.支付
	paymentRes, err := rpc.PaymentClient.Charge(s.ctx, payReq)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(&email.EmailReq{
		From: "gomall@gamall.com",
		To: req.Email,
		ContentType: "text/plain",
		Subject: "You have created an order on Gomall",
		Content: "You have create an order on Gomall",
		
	})
	msg := nats.Msg{
		Subject: "email",
		Data: data,
	}
	_ = mq.Nc.PublishMsg(&msg)
	resp = &checkout.CheckoutResp{
		OrderId: orderId,
		TransactionId: paymentRes.TransactionId,
	}
	return
}
