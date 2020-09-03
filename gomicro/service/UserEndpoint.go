package service

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
)

type UserReuest struct {
	Uid int `json:"uid"`
	Method string `json:"method"`
}

type UserResponse struct {
	Result string `json:"result"`
}

func GenUserEnpoint(service UserService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserReuest)

		reslut := "NOTHONING"
		if r.Method == "GET" {
			reslut = service.GetName(r.Uid)
		} else if r.Method == "DELETE"  {
			err :=service.DeleteUser(r.Uid)
			if err != nil{
				reslut = err.Error()
			}
		} else {
			reslut =  fmt.Sprintf("success")
		}
		//reslut := service.GetName(r.Uid)
		return UserResponse{Result:reslut},nil
	}
}