package main

import (
	"context"
	"errors"
	"fmt"
	"gomall/demo/demo_proto/kitex_gen/pbapi"
	"gomall/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"gomall/demo/demo_proto/middleware"
	"log"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func main() {
	// r, err := consul.NewConsulResolver("47.104.98.71:8500")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	c, err := echoservice.NewClient("demo_proto",
		client.WithHostPorts("localhost:8888"),
		client.WithMiddleware(middleware.MiddleWare))
	if err != nil {
		log.Fatal(err)
	}
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto_client")
	res, err := c.Echo(ctx, &pbapi.Request{Message: "test"})
	var bizerr *kerrors.GRPCBizStatusError
	if err != nil {
		_ = errors.As(err, bizerr)
		fmt.Printf("%#v", bizerr)
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
}
