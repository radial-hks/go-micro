package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
	"gomicro/util"
	"strconv"
)

type UserReuest struct {
	Uid    int    `json:"uid"`
	Method string `json:"method"`
}

type UserResponse struct {
	Result string `json:"result"`
}

func RateLimit(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if limit.Allow() {
				return nil, errors.New("too many requests")
			}
			return next(ctx, request)
		}
	}
}

func GenUserEnpoint(service UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserReuest)

		reslut := "NOTHONING"
		if r.Method == "GET" {
			reslut = service.GetName(r.Uid) + strconv.Itoa(util.ServicePort)
		} else if r.Method == "DELETE" {
			err := service.DeleteUser(r.Uid)
			if err != nil {
				reslut = err.Error()
			}
		} else {
			reslut = fmt.Sprintf("success")
		}
		//reslut := service.GetName(r.Uid)
		return UserResponse{Result: reslut}, nil
	}
}
