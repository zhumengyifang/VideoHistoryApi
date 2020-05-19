package ServiceModel

type ResponseHeader struct {
	Version   int            `json:"version"`
	IsSuccess bool           `json:"is_success"`
	Error     *ResponseError `json:"error"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseBody struct {
	Header ResponseHeader `json:"header"`
	Body   interface{}    `json:"body"`
}
