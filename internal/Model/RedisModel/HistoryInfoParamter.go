package RedisModel

import "time"

type HistoryInfo struct {
	////用户唯一标识
	OpenId string `json:"openId"`
	//视频唯一标识
	VideoId string `json:"videoId"`
	//看到了哪里
	UseTime int `json:"useTime"`
	//视频作者
	AuthorName string `json:"authorName"`
	//视频标签
	Title *[]string `json:"title"`
	//封面url
	CoverUrl string `json:"coverUrl"`

	//提交时间
	SubmitDate time.Time `json:"submitDate"`
	//是否删除
	IsDelete bool `json:"isDelete"`
}

type HistoryInfos []*HistoryInfo

func (v HistoryInfos) Len() int {
	return len(v)
}

func (v HistoryInfos) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v HistoryInfos) Less(i, j int) bool {
	return v[i].SubmitDate.Unix() > v[j].SubmitDate.Unix()
}
