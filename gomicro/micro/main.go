package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	services "micro/Services"
	"micro/Weblib"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	md, _ := metadata.FromContext(ctx)
	fmt.Printf("[Log Wrapper] ctx: %v service: %s method: %s\n", md, req.Service(), req.Endpoint())
	return l.Client.Call(ctx, req, rsp)
}

func NewLogWrapper(c client.Client) client.Client {
	return &logWrapper{c}
}

func main() {

	// consul_url: https://segmentfault.com/a/1190000023529475?utm_source=tag-newest
	consulreg := consul.NewRegistry(
		// docker inspect ipAddress + port
		//registry.Addrs("120.79.44.169:8500"),
		registry.Addrs("127.0.0.4:8500"),
	)
	myService := micro.NewService(
		micro.Name("product_service_client"),
		micro.WrapClient(NewLogWrapper))
	//product_service
	ProdService := services.NewProdService("product_service", myService.Client())
	httpserver := web.NewService(
		// name
		web.Name("product_service"),
		// go run main.go --server_address :9090
		web.Address(":8087"),
		web.Handler(Weblib.NewRouter(ProdService)),
		//  register consul
		web.Registry(consulreg),
	)

	httpserver.Init()
	httpserver.Run()
}
