package rpc_client

import (
	"gomall/app/frontend/conf"
	"gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/gomall/user/userservice"
	"gomall/rpc_gen/kitex_gen/gomall/product/productservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client
	ProductClient productservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		initUserRpcClient()
		initProductRpcClient()
	})
}

func initUserRpcClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	myutils.MustHandleErr(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	myutils.MustHandleErr(err)
}


func initProductRpcClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	myutils.MustHandleErr(err)
	ProductClient, err = productservice.NewClient("product", client.WithResolver(r))
	myutils.MustHandleErr(err)
}