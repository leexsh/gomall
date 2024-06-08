package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type PaymentLog struct {
	gorm.Model
	UserId int32 `json:"user_id"`	
	OrderId string `json:"orderId"`	
	TransactionId string `json:"transaction_id"`	
	Amount float32 `json:"amount"`	
	PayAt time.Time `json:"pay_at"`	
}


func (PaymentLog)TableName() string {
	return "payment_log"
}


func CreatePaymentLog(ctx context.Context, db *gorm.DB, payment *PaymentLog)  error {
	return db.WithContext(ctx).Model(&PaymentLog{}).Create(payment).Error
}