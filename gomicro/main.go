package main

import (
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"gomicro/service"
	"gomicro/util"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	user := service.UserService{}
	endp := service.GenUserEnpoint(user)

	serverHandler := httptransport.NewServer(endp,service.DecodeUserRequest,service.EncodeUserResponse)

	r := mymux.NewRouter()

	//r.Handle(`/user/{uid:\d+}`,serverHandler)
	r.Methods("GET","DELETE").Path(`/user/{uid:\d+}`).Handler(serverHandler)
	r.Methods("GET").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`{"status":"OK"}`))
	})
	//serverHandler := httptransport.NewServer(endp,service.DecodeUserRequest,service.EncodeUserResponse)
	errChan := make(chan error)
	go (func(){
		util.Regservice()
		err := http.ListenAndServe(":9090",r)
		if err!=nil{
			errChan <- err
		}
	})()
	go (func(){
		sig_c := make(chan  os.Signal)
		signal.Notify(sig_c,syscall.SIGINT,syscall.SIGTERM)
		errChan <- fmt.Errorf("%s",<- sig_c)
	})()
	getErr := <- errChan
	util.UnregService()
	log.Println(getErr)
}