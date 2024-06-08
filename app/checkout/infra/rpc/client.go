package rpc

import (
	"gomall/app/cart/kitex_gen/gomall/cart/cartservice"
	"gomall/app/checkout/conf"
	"gomall/app/payment/kitex_gen/gomall/payment/paymentservice"
	"gomall/app/product/kitex_gen/gomall/product/productservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	ProductClient productservice.Client
	CartClient cartservice.Client
	PaymentClient paymentservice.Client
	once sync.Once
	err error
)

func Init() {
	once.Do(func ()  {
		initCartRpcClient()
		initPaymentRpcClient()
		initProductRpcClient()
	})
}

func initProductRpcClient() {
	r, _ := consul.NewConsulResolver(conf.GetConf().Kitex.RegistryAddr)
	ProductClient, _ = productservice.NewClient("product", client.WithResolver(r))
}


func initCartRpcClient()  {
	r, _ := consul.NewConsulResolver(conf.GetConf().Kitex.RegistryAddr)
	CartClient, _ = cartservice.NewClient("cart", client.WithResolver(r))
}

func initPaymentRpcClient()  {
	r, _ := consul.NewConsulResolver(conf.GetConf().Kitex.RegistryAddr)
	PaymentClient, _ = paymentservice.NewClient("payment", client.WithResolver(r))
}