package Service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func GetUserInfo_req(c context.Context, req *http.Request, r interface{}) error {
	user_req := r.(UserReuest)
	req.URL.Path += "/user/" + strconv.Itoa(user_req.Uid)
	fmt.Println(req.URL)
	return nil
}

//
func GetUserInfo_res(c context.Context, res *http.Response) (response interface{}, err error) {
	if res.StatusCode > 400 {
		return nil, errors.New("no data")
	}
	var user_res UserResponse
	fmt.Println(res)
	err = json.NewDecoder(res.Body).Decode(&user_res)
	if err != nil {
		return nil, err
	}
	fmt.Println(user_res)
	return user_res, nil
}
