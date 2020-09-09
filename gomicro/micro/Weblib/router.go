package Weblib

import (
	"github.com/gin-gonic/gin"
	services "micro/Services"
	"net/http"
)

func NewRouter(service services.ProdService) *gin.Engine {
	r := gin.Default()
	r.Use(InitMiddleWare(service))
	r.Handle("GET", "/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	})
	// Add Route Group
	v1group := r.Group("/v1")
	{
		v1group.Handle("POST", "/prod", GetProdslist)
	}
	return r

}
