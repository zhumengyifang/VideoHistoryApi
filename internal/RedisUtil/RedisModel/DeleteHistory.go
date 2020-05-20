package RedisModel

type DeleteHistoryParameter struct {
	//用户唯一标识
	OpenId string `json:"openId"`
	//视频唯一标识
	VideoIds []string `json:"videoIds"`
}

type DeleteHistoryResponse struct {
	//用户唯一标识
	OpenId string `json:"openId"`
	//删除数量
	DeleteInfo map[string]bool `json:"deleteInfo"`
}
