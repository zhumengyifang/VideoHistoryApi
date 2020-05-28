package ServiceModel

type Header struct {
	Version   int    `json:"version" binding:"required,version"`
	AppID     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type PostBody struct {
	Header Header      `json:"header" binding:"required,dive"`
	Body   interface{} `json:"body" binding:"required,dive"`
}
