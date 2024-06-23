package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)


var Registry *prometheus.Registry
func InitMetric(serviceName, metricsPort, registerAddr string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	r, _ := consul.NewConsulRegister(registerAddr)
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)
	registerInfo := &registry.Info{
		ServiceName: serviceName,
		Addr: addr,
		Weight: 1,
		Tags: map[string]string{"service":serviceName},
	}
	_ = r.Register(registerInfo)
	server.RegisterShutdownHook(func ()  {
		r.Deregister(registerInfo)
	})
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))
	go http.ListenAndServe(metricsPort, nil)
	return r, registerInfo 
}