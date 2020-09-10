package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"micro_double/Servicehlmpl"
	"micro_double/Services"
)

func main() {
	consulreg := consul.NewRegistry(
		// docker inspect ipAddress + port
		//registry.Addrs("120.79.44.169:8500"),
		registry.Addrs("127.0.0.4:8500"),
	)
	myservice := micro.NewService(
		micro.Name("test_service"),
		micro.Address(":8081"),
		micro.Registry(consulreg),
	)
	Services.RegisterTestServiceHandler(myservice.Server(), new(Servicehlmpl.TestService))
	myservice.Run()
}
