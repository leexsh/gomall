package main

import (
	"gomall/demo/demo_proto/biz/dal"
	"gomall/demo/demo_proto/biz/dal/mysql"
	"gomall/demo/demo_proto/biz/model"
)

func main() {
	dal.Init()
	mysql.DB.Create(&model.User{
		Email:    "test@qq.com",
		Password: "sadasf",
	})
}
