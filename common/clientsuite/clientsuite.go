package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonClientSuite struct{
	CurrentServerName string
	RegisterAddr string
}

func (s *CommonClientSuite)Options() []client.Option {
	opts := []client.Option {
		client.WithMetaHandler(transmeta.ServerHTTP2Handler),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServerName,
		}),
		client.WithTransportProtocol(transport.GRPC),
		client.WithSuite(tracing.NewClientSuite()),
	}
	r, err := consul.NewConsulResolver(s.RegisterAddr)
	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	return opts
}