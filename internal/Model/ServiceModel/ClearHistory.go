package ServiceModel

type ClearHistoryParameter struct {
	//用户唯一标识
	OpenId string `json:"openId" binding:"required"`
}

