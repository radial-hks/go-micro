package Weblib

import (
	"github.com/gin-gonic/gin"
	services "micro/Services"
)

func InitMiddleWare(service services.ProdService) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Keys = make(map[string]interface{})
		context.Keys["prodService"] = service
		context.Next()
	}
}
