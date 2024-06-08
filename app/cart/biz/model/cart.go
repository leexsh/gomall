package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserId int32 `gorm:"type:int(11); not null; index:idx_user_id"`
	ProductId int32 `gorm:"type:int(11); not null;"`
	Qty int32`gorm:"type:int(11); not null;"`
}

func (Cart)TableName() string {
	return "cart"
}

func AddItem(ctx context.Context, db *gorm.DB, ca *Cart) error {
	var row Cart
	err := db.WithContext(ctx).Model(&Cart{}).Where(&Cart{
		UserId:ca.UserId,
		ProductId: ca.ProductId,
	}).First(&row).Error
	if err != nil {
		fmt.Println(err)
	}
	if row.ID > 0 {
		return db.WithContext(ctx).Model(&Cart{}).Where(
			&Cart{
				UserId: ca.UserId, 
				ProductId: ca.ProductId}).Update("qty", gorm.Expr("qty+?", ca.Qty)).Error
	}
	return db.WithContext(ctx).Create(ca).Error
}

func EmptyItem(ctx context.Context, db *gorm.DB, id int32) error {
	return db.WithContext(ctx).Delete(&Cart{}, "user_id=?", id).Error
}

func GetCartByUserId(db *gorm.DB, ctx context.Context, userId uint32) (cartList []*Cart, err error) {
	err = db.Debug().WithContext(ctx).Model(&Cart{}).Find(&cartList, "user_id = ?", userId).Error
	return cartList, err
}