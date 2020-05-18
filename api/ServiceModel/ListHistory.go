package ServiceModel

type ListHistoryParameter struct {
	//用户唯一标识
	OpenID string `json:"openId"`

	PageCount int `json:"pageCount"`
	PageSize  int `json:"pageSize"`
}

type HistoryListResponse struct {
	PageCount int `json:"pageCount"`
	PageSize  int `json:"pageSize"`

	Videos []HistoryInfoResponse `json:"videos"`
}
