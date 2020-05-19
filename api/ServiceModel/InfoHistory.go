package ServiceModel

type InfoHistoryParameter struct {
	//用户唯一标识
	OpenId string `json:"openId"`
	//视频唯一标识
	VideoId string `json:"videoId"`
}

type InfoHistoryResponse struct {
	////用户唯一标识
	OpenId string `json:"openId"`
	//视频唯一标识
	VideoId string `json:"videoId"`
	//看到了哪里
	UseTime int `json:"useTime"`
	//视频作者
	AuthorName string `json:"authorName"`
	//视频标签
	Title []string `json:"title"`
	//封面url
	CoverUrl string `json:"coverUrl"`
}