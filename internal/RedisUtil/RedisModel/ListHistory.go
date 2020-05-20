package RedisModel

type ListHistoryResponse struct {
	PageCount int `json:"pageCount"`
	PageSize  int `json:"pageSize"`

	Videos []HistoryInfoParameter `json:"videos"`
}
