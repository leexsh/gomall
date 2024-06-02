package rpc_client

import (
	"gomall/app/frontend/conf"
	"gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/gomall/user/userservice"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func Init() {
	once.Do(func() {
		initUserRpcClient()
	})
}

func initUserRpcClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddr)
	myutils.MustHandleErr(err)
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	myutils.MustHandleErr(err)
}
