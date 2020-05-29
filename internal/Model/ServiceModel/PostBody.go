package ServiceModel

type Header struct {
	Version   int    `json:"version" binding:"min=1"`
	AppID     string `json:"appId" binding:"required"`
	AppSecret string `json:"appSecret" binding:"required"`
}

type PostBody struct {
	Header Header      `json:"header" binding:"required,dive"`
	Body   interface{} `json:"body" binding:"required,dive"`
}
