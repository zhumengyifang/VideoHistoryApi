package Controllers

import (
	"encoding/json"
	"gindemo/api/ServiceModel"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func Convert(c *gin.Context, postBody *ServiceModel.PostBody) (*ServiceModel.PostBody, error) {
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

func HandelError(c *gin.Context, responseBody ServiceModel.ResponseBody) {
	if !responseBody.Header.IsSuccess {
		c.JSON(200, gin.H{
			"body": responseBody,
		})
		return
	}

	c.JSON(200, gin.H{
		"body": responseBody,
	})
}
