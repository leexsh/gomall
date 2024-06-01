package rpc_client

import (
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
	r, err := consul.NewConsulResolver("47.104.98.71:8500")
	if err != nil {
		return
	}
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		return
	}

}
