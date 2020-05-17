package ServiceModel

type Header struct {
	Version   int
	AppId     string
	AppSecret string
}

type PostBody struct {
	Header Header
	Body   interface{}
}
