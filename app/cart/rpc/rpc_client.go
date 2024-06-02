package rpc_client

import (
	"gomall/app/cart/conf"
	"gomall/app/product/kitex_gen/gomall/product/productservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)
var (
	ProductClient productservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		initProductRpcClient()
	})
}


func initProductRpcClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Kitex.RegistryAddr)
	MustHandleErr(err)
	ProductClient, err = productservice.NewClient("product", client.WithResolver(r))
	MustHandleErr(err)
}