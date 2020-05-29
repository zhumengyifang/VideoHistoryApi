package ServiceModel

type ListHistoryParameter struct {
	//用户唯一标识
	OpenId string `json:"openId" binding:"required"`

	PageCount int `json:"pageCount"  binding:"min=1"`
	PageSize  int `json:"pageSize"  binding:"min=1"`
}

type ListHistoryResponse struct {
	PageCount int `json:"pageCount"`
	PageSize  int `json:"pageSize"`

	Videos []*InfoHistoryResponse `json:"videos"`
}
