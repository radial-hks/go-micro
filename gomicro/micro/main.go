package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	services "micro/Services"

	//"net/http"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// consul_url: https://segmentfault.com/a/1190000023529475?utm_source=tag-newest
	consulreg := consul.NewRegistry(
		// docker inspect ipAddress + port
		//registry.Addrs("120.79.44.169:8500"),
		registry.Addrs("127.0.0.4:8500"),
	)
	//addr := func(o *web.Options) {
	//	o.Address = ":8001"
	//}

	// go-micro/web
	//server := web.NewService(web.Address(":9090"))
	//server.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("radial"))
	//})
	//server.Run()
	//  gin +  go-micro/web
	r := gin.Default()
	r.Handle("GET", "/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	})

	httpserver := web.NewService(
		// name
		web.Name("product_service"),
		// go run main.go --server_address :9090
		web.Address(":8086"),
		web.Handler(r),
		//  register consul
		web.Registry(consulreg),
	)
	myService := micro.NewService(
		micro.Name("product_service_client"),
	)
	//product_service
	ProdService := services.NewProdService("product_service", myService.Client())

	// Add Route Group
	v1group := r.Group("/v1")
	{
		v1group.Handle("POST", "/prod", func(gin_context *gin.Context) {
			var prodReq services.ProdsRequest
			err := gin_context.Bind(&prodReq)
			fmt.Println(prodReq, err)
			if err != nil {
				gin_context.JSON(500, gin.H{
					"status": err.Error(),
				})
			} else {
				ProdRes, err := ProdService.GetProdsList(context.Background(), &prodReq)
				fmt.Println(err)
				if err != nil {
					gin_context.JSON(500, gin.H{
						"status": err.Error(),
					})
				} else {
					gin_context.JSON(200, gin.H{
						"data": ProdRes.Data,
					})
				}

			}

		})
	}

	httpserver.Init()
	httpserver.Run()
}
