package rpc_client

import (
	"gomall/app/cart/kitex_gen/gomall/cart/cartservice"
	"gomall/app/frontend/conf"
	"gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/gomall/checkout/checkoutservice"
	"gomall/rpc_gen/kitex_gen/gomall/product/productservice"
	"gomall/rpc_gen/kitex_gen/gomall/user/userservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient     userservice.Client
	ProductClient  productservice.Client
	CartClient     cartservice.Client
	CheckoutClient checkoutservice.Client
	once           sync.Once
)

func Init() {
	once.Do(func() {
		initUserRpcClient()
		initProductRpcClient()
		initCartRpcClient()
		initCheckoutRpcClient()
	})
}

func initUserRpcClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	myutils.MustHandleErr(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	myutils.MustHandleErr(err)
}

func initProductRpcClient() {
	r, _ := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	ProductClient, _ = productservice.NewClient("product", client.WithResolver(r))
}

func initCartRpcClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	myutils.MustHandleErr(err)
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	myutils.MustHandleErr(err)
}

func initCheckoutRpcClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	myutils.MustHandleErr(err)
	CheckoutClient, err = checkoutservice.NewClient("checkout", client.WithResolver(r))
	myutils.MustHandleErr(err)
}
