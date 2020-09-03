package service

import (
	"context"
	"encoding/json"
	"errors"
	mymux "github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context,r *http.Request) (interface{},error){
	//if r.URL.Query().Get("uid") != ""{
	//	uid,_ := strconv.Atoi(r.URL.Query().Get("uid"))
	//	return  UserReuest{Uid: uid},nil
	//}
	vars := mymux.Vars(r)
	if uid,ok := vars["uid"];ok{
		uid,_ := strconv.Atoi(uid)
		return  UserReuest{
			Uid: uid,
			Method: r.Method,
		},nil
	}

	return nil,errors.New("Wrong")

}

func  EncodeUserResponse(ctx context.Context,w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-type","appliication/json")
	return json.NewEncoder(w).Encode(response)
}