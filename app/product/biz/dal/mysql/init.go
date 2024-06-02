package mysql

import (
	"gomall/app/product/biz/model"
	"gomall/app/product/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB.AutoMigrate(&model.Product{})
	DB.AutoMigrate(&model.Category{})
	if err != nil {
		panic(err)
	}
}
