package Controllers

import (
	"gindemo/internal/Model/ServiceModel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(c *gin.Context, responseBody ServiceModel.ResponseBody)  {
	if !responseBody.Header.IsSuccess {
		c.JSON(http.StatusOK, responseBody)
		c.Abort()
	} else {
		c.JSON(http.StatusOK, responseBody)
	}
}
