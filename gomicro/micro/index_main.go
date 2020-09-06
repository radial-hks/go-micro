package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/web"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Handle("GET", "/", func(context *gin.Context) {
		data := make([]int, 0)
		context.JSON(http.StatusOK, gin.H{
			"result": data,
		})
	})

	//r.Run(":9090")
	server := web.NewService(
		web.Address(":9091"),
		web.Handler(r),
	)
	server.Run()
}
