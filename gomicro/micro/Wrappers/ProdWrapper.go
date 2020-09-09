package Wrappers

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/client"
	services "micro/Services"
	"strconv"
)

func NewProd(id int32, pname string) *services.ProdModel {
	return &services.ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

func defaultProds(rsp interface{}) {
	Models := make([]*services.ProdModel, 0)
	var i int32
	for i = 0; i < 1; i++ {
		name := "service" + strconv.Itoa(int(i))
		Models = append(Models, NewProd(100+i, name))
	}
	//res := &services.ProdResponse{}
	res := rsp.(*services.ProdResponse)
	res.Data = Models
	//return res, nil
}

type ProdWrapper struct {
	client.Client
}

func (l *ProdWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout: 1000,
	}
	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return l.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		defaultProds(rsp)
		return nil
	})
}

func NewProdWrapper(c client.Client) client.Client {
	return &ProdWrapper{c}
}
