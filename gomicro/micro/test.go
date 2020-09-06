package main

import (
	"fmt"
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func callAPI(address string, path string, method string) (string, error) {
	req, err := http.NewRequest(method, "http://"+address+path, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(buff), nil
}

func main() {
	consulreg := consul.NewRegistry(
		// docker inspect ipAddress + port
		registry.Addrs("120.79.44.169:8500"),
	)
	for {
		getService, err := consulreg.GetService("product_service")
		if err != nil {
			log.Fatal(err)
		}
		// next:= selector.Random(getService)
		next := selector.RoundRobin(getService)
		node, err := next()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(node.Id, node.Address, node.Metadata)
		time.Sleep(time.Second * 1)
		callRes, err := callAPI(node.Address, "/v1/prod", "GET")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(callRes)
	}
}
