package service

import (
	"context"
	"gomall/app/payment/biz/dal/mysql"
	"gomall/app/payment/biz/model"
	payment "gomall/app/payment/kitex_gen/gomall/payment"
	"strconv"
	"time"

	creditcard "github.com/durango/go-credit-card"
	"github.com/google/uuid"
)

type ChargeService struct {
	ctx context.Context
} // NewChargeService new ChargeService
func NewChargeService(ctx context.Context) *ChargeService {
	return &ChargeService{ctx: ctx}
}

// Run create note info
func (s *ChargeService) Run(req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	// Finish your business logic.
	card := creditcard.Card{
		Number: req.CreditInfo.CreditCardNumber,
		Cvv: strconv.Itoa(int(req.CreditInfo.CreditCardCvv)),
		Year: strconv.Itoa(int(req.CreditInfo.CreditCardExpirationYear)),
		Month: strconv.Itoa(int(req.CreditInfo.CreditCardExpirationMonth)),
	}
	err = card.Validate(true)
	if err != nil {
		return nil, err
	}
	transId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	err = model.CreatePaymentLog(s.ctx, mysql.DB, &model.PaymentLog{
		UserId: req.UserId,
		OrderId: req.OrderId,
		TransactionId: transId.String(),
		Amount: float32(req.Amount),
		PayAt: time.Now(),

	})
	if err != nil {
		return nil, err
	}
	resp = payment.NewChargeResp()
	resp.TransactionId = transId.String()
	return
}
