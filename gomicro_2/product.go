package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	httptransport "github.com/go-kit/kit/transport/http"
	consulapi "github.com/hashicorp/consul/api"
	. "gomicro_2/Service"
	"io"
	"net/url"
	"os"
)


func main_2() {
	//httptransport.NewClient(method,target, enc,dec)
	tgt,_:= url.Parse("http://127.0.0.1:9090")
	// 第一步：创建一个直连client 必须传入两个func:如何请求及响应怎麼处理
	client := httptransport.NewClient("GET",tgt, GetUserInfo_req,GetUserInfo_res)
	// 第二步：暴露 endpoint
	get_user_info := client.Endpoint()
	// 第三步：创建一个上下文对象
	ctx := context.Background()
	// 第四步：执行
	res,err  := get_user_info(ctx,UserReuest{Uid: 101})
	if err != nil{
		fmt.Println(err)
	}
	// 第五步： 断言 得到响应值
	userinfo :=res.(UserResponse)
	fmt.Println(userinfo.Result)
	fmt.Println(userinfo)
}

func main(){

	config := consulapi.DefaultConfig()
	config.Address = "127.0.0.1:8500"

	api_client,_  := consulapi.NewClient(config)
	client := consul.NewClient(api_client)

	var logger log.Logger
	{
		logger =  log.NewLogfmtLogger(os.Stdout)
	}
	{
		tags:= []string{"dev"}
		instancer := consul.NewInstancer(client,logger,"hks",tags,true)

		{
			f := func(service_url string )(endpoint.Endpoint , io.Closer, error){
				tgt,_:= url.Parse("http://" + service_url)
				return httptransport.NewClient("GET",tgt, GetUserInfo_req,GetUserInfo_res).Endpoint(),nil,nil
			}

			endpointer := sd.NewEndpointer(instancer,f,logger)
			endpoints,_ := endpointer.Endpoints()
			fmt.Println(len(endpoints))
			//
			get_user_info := endpoints[0]
			// 第三步：创建一个上下文对象
			ctx := context.Background()
			// 第四步：执行
			res,err  := get_user_info(ctx,UserReuest{Uid: 101})
			if err != nil{
				fmt.Println(err)
			}
			// 第五步： 断言 得到响应值
			userinfo :=res.(UserResponse)
			fmt.Println(userinfo.Result)
			fmt.Println(userinfo)

		}
	}


}