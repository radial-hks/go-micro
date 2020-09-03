package util

import (
	consul "github.com/hashicorp/consul/api"
	"log"
)

var CLIENT *consul.Client

func init(){
	config :=  consul.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client,err := consul.NewClient(config)
	if err != nil{
		log.Fatal("wrong")
	}
	CLIENT = client
}

//func Regservice(){
//	config :=  consul.DefaultConfig()
//	config.Address = "127.0.0.1:8500"
//	//
//	reg := consul.AgentServiceRegistration{}
//	reg.ID = "15"
//	reg.Name = "hks"
//	reg.Address = "127.0.0.1"
//	reg.Port = 80
//	reg.Tags = []string{"dev"}
//
//	check := consul.AgentServiceCheck{}
//	check.Interval =  "5s"
//	check.HTTP = "http://127.0.0.1:9090/health"
//
//	reg.Check = &check
//
//	client,err := consul.NewClient(config)
//	CLIENT = client
//	if err != nil{
//		log.Fatal("wrong")
//	}
//	if err:=client.Agent().ServiceRegister(&reg);err != nil {
//		log.Fatal("wrong")
//	}
//}

// node.json
//{
//"id":"14",
//"name":"radial_node",
//"tags":["dev"],
//"port":80,
//"address":"127.0.0.1",
//"check":{
//"HTTP":"http://127.0.0.1:9090/health",
//"interval":"5s"
//}
//}

func Regservice(){
	reg := consul.AgentServiceRegistration{}
	reg.ID = "15"
	reg.Name = "hks"
	reg.Address = "127.0.0.1"
	reg.Port = 80
	reg.Tags = []string{"dev"}

	check := consul.AgentServiceCheck{}
	check.Interval =  "5s"
	check.HTTP = "http://127.0.0.1:9090/health"

	reg.Check = &check

	if err:=CLIENT.Agent().ServiceRegister(&reg);err != nil {
		log.Fatal("wrong")
	}
}

func UnregService(){
	CLIENT.Agent().ServiceDeregister("15")
}