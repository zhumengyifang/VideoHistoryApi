package RedisModel

type ClearHistoryParameter struct {
	//用户唯一标识
	OpenId string `json:"openId"`
}

type ClearHistoryResponse struct {
	Count string `json:"count"`
}
