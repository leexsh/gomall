package serversuite

import (
	"gomall/common/mtl"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	prom "github.com/kitex-contrib/monitor-prometheus"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonServerSuite struct{
	CurrentServerName string
	RegisteryAddr string
}

func (s *CommonServerSuite)Options() []server.Option {
	opts := []server.Option {
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServerName,
		}),
		server.WithTracer(prom.NewServerTracer("", "", prom.WithDisableServer(true), prom.WithRegistry(mtl.Registry))),
		server.WithSuite(tracing.NewServerSuite()),
	}
	r, _ := consul.NewConsulRegister(s.RegisteryAddr)
	opts = append(opts, server.WithRegistry(r))
	return opts
}