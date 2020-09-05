package main

import (
	"flag"
	"gomicro/util"
	"log"
)

func main() {
	name := flag.String("name", "", "serviceName")
	port := flag.Int("port", 0, "port_id")
	flag.Parse()
	//fmt.Println(*name)
	if *name == "" {
		log.Fatal("reinput servicename")
	}
	if *port == 0 {
		log.Fatal("reinput portid")
	}
	util.SetServiceNameAndPort(*name, *port)

}
