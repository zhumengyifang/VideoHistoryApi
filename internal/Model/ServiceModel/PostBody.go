package ServiceModel

type Header struct {
	Version   int    `form:"version" json:"version" binding:"required"`
	AppID     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type PostBody struct {
	Header Header      `json:"header" binding:"required,dive"`
	Body   interface{} `json:"body" binding:"required,dive"`
}
