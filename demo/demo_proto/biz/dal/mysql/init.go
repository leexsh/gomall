package mysql

import (
	"fmt"
	"gomall/demo/demo_proto/biz/model"
	"gomall/demo/demo_proto/conf"

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
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(model.User{})
	type Version struct {
		Version string
	}
	var v Version
	fmt.Println("test")
	err := DB.Raw("select version() as version").Scan(&v).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}