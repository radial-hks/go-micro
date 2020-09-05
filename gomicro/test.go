package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	//"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
	"time"
)

func main_wait() {
	r := rate.NewLimiter(1, 5)
	ctx := context.Background()

	for {
		err := r.WaitN(ctx, 2)
		if err != nil {
			log.Fatal("err")
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}

}

func main_allow() {
	r := rate.NewLimiter(1, 5)
	//ctx := context.Background()
	for {
		bool_ := r.AllowN(time.Now(), 2)
		if bool_ {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		} else {
			fmt.Println("too many request ")
		}
		// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(time.Second)
	}

}

var r = rate.NewLimiter(1, 5)

func mylimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !r.AllowN(time.Now(), 2) {
			http.Error(writer, "too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(writer, request)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if !r.AllowN(time.Now(), 2) {
			http.Error(writer, "too many requests", http.StatusTooManyRequests)
		} else {
			writer.Write([]byte("ok"))
		}

	})
	http.ListenAndServe(":8080", mux)
	http.ListenAndServe(":8080", mylimit(mux))
}
