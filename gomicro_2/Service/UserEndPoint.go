package Service

type UserReuest struct {
	Uid int `json:"uid"`
	Method string `json:"method"`
}

type UserResponse struct {
	Result string `json:"result"`
}


