package ServiceModel

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Header struct {
	Version   int    `json:"version"`
	AppID     string `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type PostBody struct {
	Header Header      `json:"header"`
	Body   interface{} `json:"body"`
}

func Convert(c *gin.Context, postBody *PostBody) (*PostBody, error) {
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, postBody)
	if err != nil {
		return nil, err
	}
	return postBody, nil
}
