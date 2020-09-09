package Weblib

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gin-gonic/gin"
	services "micro/Services"
	"strconv"
)

func NewProd(id int32, pname string) *services.ProdModel {
	return &services.ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

func defaultProds() (*services.ProdResponse, error) {
	Models := make([]*services.ProdModel, 0)
	var i int32
	for i = 0; i < 1; i++ {
		name := "service" + strconv.Itoa(int(i))
		Models = append(Models, NewProd(100+i, name))
	}
	res := &services.ProdResponse{}
	res.Data = Models
	return res, nil
}

// gin
func GetProdslist(gin_context *gin.Context) {
	ProdService := gin_context.Keys["prodService"].(services.ProdService)
	var prodReq services.ProdsRequest
	err := gin_context.Bind(&prodReq)
	//fmt.Println(prodReq, err)
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
		},
			func(err error) error {
				ProdRes, err = defaultProds()
				return nil
			})

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
