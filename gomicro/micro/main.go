package main

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/registry/consul"
	"micro/Helper"
	"micro/ProdService"

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
	// Add Route Group
	v1group := r.Group("/v1")
	{
		v1group.Handle("POST", "/prod", func(context *gin.Context) {
			var pr Helper.ProdsRequest
			err := context.Bind(&pr)
			if err != nil || pr.Size <= 0 {
				pr = Helper.ProdsRequest{
					Size: 2,
				}
			}
			context.JSON(http.StatusOK, gin.H{
				"data": ProdService.NewProdList(pr.Size),
			})
		})
	}
	//r.Run(":9090")
	server := web.NewService(
		// name
		web.Name("product_service"),
		// go run main.go --server_address :9090
		//web.Address(":9090"),
		web.Handler(r),
		//  register consul
		web.Registry(consulreg),
	)
	server.Init()
	server.Run()
}
