package ServiceModel

type ListHistoryParameter struct {
	//用户唯一标识
	OpenId string `form:"openId" json:"openId" binding:"required"`

	PageCount int `json:"pageCount"`
	PageSize  int `json:"pageSize"`
}

type ListHistoryResponse struct {
	PageCount int `json:"pageCount"`
	PageSize  int `json:"pageSize"`

	Videos []*InfoHistoryResponse `json:"videos"`
}
