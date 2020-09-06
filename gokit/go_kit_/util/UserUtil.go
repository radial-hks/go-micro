package util

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
	. "gomicro_2/Service"
	"io"
	"net/url"
	"os"
	"time"
)

func GetUser() (string, error) {
	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"

	api_client, err := consulapi.NewClient(config)
	if err != nil {
		return "", err
	}
	client := consul.NewClient(api_client)

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
	}
	{
		tags := []string{"dev"}
		instancer := consul.NewInstancer(client, logger, "user_service", tags, true)

		{
			f := func(service_url string) (endpoint.Endpoint, io.Closer, error) {
				tgt, _ := url.Parse("http://" + service_url)
				return httptransport.NewClient("GET", tgt, GetUserInfo_req, GetUserInfo_res).Endpoint(), nil, nil
			}

			endpointer := sd.NewEndpointer(instancer, f, logger)

			endpoints, err := endpointer.Endpoints()
			if err != nil {
				return "", err
			}
			fmt.Println(len(endpoints))

			myld := lb.NewRoundRobin(endpointer)
			//myld := lb.NewRandom(endpointer,time.Now().UnixNano())

			//get_user_info := endpoints[0]
			get_user_info, err := myld.Endpoint()
			if err != nil {
				return "", err
			}
			// 第三步：创建一个上下文对象
			ctx := context.Background()
			// 第四步：执行
			res, err := get_user_info(ctx, UserReuest{Uid: 101})
			if err != nil {
				return "", err
			}
			// 第五步： 断言 得到响应值
			userinfo := res.(UserResponse)
			fmt.Println(userinfo.Result)
			//fmt.Println(userinfo)
			time.Sleep(3 * time.Second)
			return userinfo.Result, nil

		}
	}
}
