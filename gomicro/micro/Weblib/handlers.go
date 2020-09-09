package Weblib

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	services "micro/Services"
)

func GetProdslist(gin_context *gin.Context) {
	ProdService := gin_context.Keys["prodService"].(services.ProdService)
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

}
