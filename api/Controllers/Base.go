package Controllers

import (
	"encoding/json"
	"gindemo/internal/Model/ServiceModel"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
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
		c.JSON(http.StatusOK, responseBody)
		return
	}

	c.JSON(http.StatusOK, responseBody)
}
