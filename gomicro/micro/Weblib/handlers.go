package Weblib

import (
	"context"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	services "micro/Services"
)

// gin
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
		// hysterix code review
		// 1 set config
		config := hystrix.CommandConfig{
			Timeout: 1000,
		}
		// 2 ser command
		hystrix.ConfigureCommand("getProdService", config)
		// 3  Process Do
		var ProdRes *services.ProdResponse
		err := hystrix.Do("getProdService", func() error {
			ProdRes, err = ProdService.GetProdsList(context.Background(), &prodReq)
			return err
		}, nil)

		//  ProdRes, err := ProdService.GetProdsList(context.Background(), &prodReq)
		//fmt.Println(err)
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
