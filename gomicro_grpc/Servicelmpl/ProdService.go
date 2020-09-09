package Servicelmpl

import (
	"context"
	services "gomicro_grpc/Services"
	"strconv"
	"time"
)

type ProdServices struct {
}

func NewProd(id int32, pname string) *services.ProdModel {
	return &services.ProdModel{
		ProdID:   id,
		ProdName: pname,
	}
}

func (*ProdServices) GetProdsList(ctx context.Context, in *services.ProdsRequest, res *services.ProdResponse) error {
	time.Sleep(time.Second * 4)
	Models := make([]*services.ProdModel, 0)
	var i int32
	for i = 0; i < in.Size; i++ {
		name := "service" + strconv.Itoa(int(i))
		Models = append(Models, NewProd(100+i, name))
	}
	res.Data = Models
	return nil
}
