package Weblib

import (
	"fmt"
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
func ErrorMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		if r := recover(); r != nil {
			context.JSON(500, gin.H{"status": fmt.Sprintf("%s", r)})
			context.Abort()
		}
		context.Next()
	}
}
