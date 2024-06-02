package model

import (
	"context"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Picture     string     `json:"picture"`
	Price       float64    `json:"price"`
	Categories  []Category `json:"categories" gorm:"many2many:product_category"`
}

func (p Product) TableName() string {
	return "product"
}

type ProductQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewProductQuery(ctx context.Context, db *gorm.DB)*ProductQuery  {
	return &ProductQuery{
		ctx: ctx,
		db: db,
	}
}
func (p ProductQuery) GetProductById(pid int) (product Product, err error) {
	err = p.db.WithContext(p.ctx).Model(&Product{}).First(&product, pid).Error
	return
}

func (p ProductQuery) SearchProducts(q string) (products []*Product, err error)  {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Where("name like ? or description like ?", "%"+q+"%", "%"+q+"%").Find(&products).Error
	return
}

func (p ProductQuery) GetAllProducts() (products []*Product, err error)  {
	err = p.db.WithContext(p.ctx).Model(&Product{}).Find(&products).Limit(100).Error
	return
}