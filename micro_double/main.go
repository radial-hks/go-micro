package main

import (
	//_ "github.com/golang/protobuf/protoc-gen-go"
	//_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway"
	//_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	//"github.com/micro/go-plugins/registry/consul"
	"micro_double/Servicehlmpl"
	"micro_double/Services"
)

func main() {
	etcdreg := etcd.NewRegistry(registry.Addrs("127.0.0.4:2379"))
	//consulreg := consul.NewRegistry(
	//	// docker inspect ipAddress + port
	//	//registry.Addrs("120.79.44.169:8500"),
	//	registry.Addrs("127.0.0.4:8500"),
	//)
	myservice := micro.NewService(
		micro.Name("api.radial.cool"),
		micro.Address(":8081"),
		micro.Registry(etcdreg),
	)
	Services.RegisterTestServiceHandler(myservice.Server(), new(Servicehlmpl.TestService))
	myservice.Run()
}
