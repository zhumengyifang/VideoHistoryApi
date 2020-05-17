package ServiceModel

type HistoryInfoParameter struct {
	//用户唯一标识
	OpenId string
	//视频唯一标识
	VideoId string
}

type HistoryInfoResponse struct {
	////用户唯一标识
	OpenId string
	//视频唯一标识
	VideoId string
	//看到了哪里
	UseTime int
	//视频作者
	AuthorName string
	//视频标签
	Title []string
	//封面url
	CoverUrl string
}
