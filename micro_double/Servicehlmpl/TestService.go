package Servicehlmpl

import (
	"context"
	"micro_double/Services"
	"strconv"
)

type TestService struct {
}

func (this *TestService) Call(ctx context.Context, req *Services.TestRequest, res *Services.TestResponse) error {
	res.Data = "text" + strconv.Itoa(int(req.Id))
	return nil
}
