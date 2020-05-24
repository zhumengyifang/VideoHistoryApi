package RedisModel

import "time"

type HistoryInfoParameter struct {
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

type HistoryInfoParameters []*HistoryInfoParameter

func (v HistoryInfoParameters) Len() int {
	return len(v)
}

func (v HistoryInfoParameters) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v HistoryInfoParameters) Less(i, j int) bool {
	return v[i].SubmitDate.Unix() < v[j].SubmitDate.Unix()
}
