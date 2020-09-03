package main

import (
	"context"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	. "gomicro_2/Service"
	"net/url"
)

func main() {
	//httptransport.NewClient(method,target, enc,dec)
	tgt,_:= url.Parse("http://127.0.0.1:9090")
	client := httptransport.NewClient("GET",tgt, GetUserInfo_req,GetUserInfo_res)
	get_user_info := client.Endpoint()
	ctx := context.Background()

	res,err  := get_user_info(ctx,UserReuest{Uid: 101})
	if err != nil{
		fmt.Println(err)
	}
	// duanyan
	userinfo :=res.(UserResponse)
	fmt.Println(userinfo.Result)
	fmt.Println(userinfo)
}

