package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
	//"github.com/micro/go-micro/v3"
	"github.com/micro/go-micro/registry"
	"gomicro_grpc/Servicelmpl"
	services "gomicro_grpc/Services"
)

//func main(){
//	// consul_url: https://segmentfault.com/a/1190000023529475?utm_source=tag-newest
//	consulreg := consul.NewRegistry(
//		// docker inspect ipAddress + port
//		//registry.Addrs("120.79.44.169:8500"),
//		registry.Addrs("127.0.0.4:8500"),
//	)
//
//	Prodservice := micro.NewService(
//		micro.Name("prod_service"),
//		micro.Address(":8083"),
//		micro.Registry(consulreg),
//		)
//	Prodservice.Init()
//	services.RegisterProdServiceHandler(Prodservice.Server(),new(Servicelmpl.ProdSercives))
//	Prodservice.Run()
//}

func main() {
	//consul注册中心
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.4:8500"),
	)

	//创建web服务器
	prodService := micro.NewService(
		micro.Name("prod_service"), // 服务名
		micro.Address(":8011"),     //端口号
		micro.Registry(consulReg),  // 注册服务
	)
	//初始化服务
	prodService.Init()
	//绑定handler -->
	services.RegisterProdServiceHandler(prodService.Server(), new(Servicelmpl.ProdServices))
	//运行服务
	prodService.Run()
}
