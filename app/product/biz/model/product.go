package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
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

type CacheProductQuery struct {
	p ProductQuery
	cacheClient *redis.Client
	profix string
}

func NewCacheProdcutQuery(ctx context.Context, db *gorm.DB, cacheclient *redis.Client) *CacheProductQuery {
	return &CacheProductQuery{
		*NewProductQuery(ctx, db),
		cacheclient,
		"product_"
	}
}

func (c CacheProductQuery) GetById(pid int) (pro Product, err error) {
	cacheKey := fmt.Sprintf("%s_%s_%d", c.profix, "product_by_id", pid)
	cacheRes := c.cacheClient.Get(c.p.ctx, cacheKey)
	err = func () error {
		if err := cacheRes.Err(); err != redis.Nil {
			return err
		}
		cacheByte, err := cacheRes.Bytes()
		if err != nil {
			return err
		}
		err = json.Unmarshal(cacheByte, &pro)
		if err != nil {
			return err
		}
		return nil
	}()
	if err != nil {
		pro, err = c.p.GetProductById(pid)
		if err != nil {
			return Product{}, err
		}
		str, err := json.Marshal(pro)
		if err != nil {
			return pro, nil
		}
		_ = c.cacheClient.Set(c.p.ctx, cacheKey, str, time.Minute*30)
	}
	return
}


func (c CacheProductQuery) SearchProducts(q string) (pros []*Product, err error) {
	return c.p.SearchProducts(q)
}


type ProductMutation struct{
	
}