package ServiceModel

type Header struct {
	Version   int    `json:"version" binding:"required"`
	AppID     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type PostBody struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}
