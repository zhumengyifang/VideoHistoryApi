package ServiceModel

type DelHistoryParameter struct {
	//用户唯一标识
	OpenId string `form:"openId" json:"openId" binding:"required"`
	//视频唯一标识
	VideoIds []string `json:"videoIds" binding:"required"`
}

type DelHistoryResponse struct {
	//用户唯一标识
	OpenId string `json:"openId"`
	//删除数量
	DeleteInfo map[string]bool `json:"deleteInfo"`
}
