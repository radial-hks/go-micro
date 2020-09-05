package main

import (
	"flag"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"gomicro/service"
	"gomicro/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

//func main() {
//	user := service.UserService{}
//	endp := service.GenUserEnpoint(user)
//
//	serverHandler := httptransport.NewServer(endp,service.DecodeUserRequest,service.EncodeUserResponse)
//	http.ListenAndServe(":9090",serverHandler)
//}

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

	user := service.UserService{}
	endp := service.GenUserEnpoint(user)

	serverHandler := httptransport.NewServer(endp, service.DecodeUserRequest, service.EncodeUserResponse)

	r := mymux.NewRouter()

	//r.Handle(`/user/{uid:\d+}`,serverHandler)
	r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)
	r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`{"status":"OK"}`))
	})
	//serverHandler := httptransport.NewServer(endp,service.DecodeUserRequest,service.EncodeUserResponse)
	errChan := make(chan error)
	go (func() {
		util.Regservice()
		id := strconv.Itoa(*port)
		err := http.ListenAndServe(":"+id, r)
		if err != nil {
			errChan <- err
		}
	})()
	go (func() {
		sig_c := make(chan os.Signal)
		signal.Notify(sig_c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sig_c)
	})()
	getErr := <-errChan
	util.UnregService()
	log.Println(getErr)
}
