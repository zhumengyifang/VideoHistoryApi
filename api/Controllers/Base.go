package Controllers

import (
	"gindemo/internal/Model/ServiceModel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandelError(c *gin.Context, responseBody ServiceModel.ResponseBody) {
	if !responseBody.Header.IsSuccess {
		c.JSON(http.StatusOK, responseBody)
		c.Abort()
	}

	c.JSON(http.StatusOK, responseBody)
}
