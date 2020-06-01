package Middleware

import (
	"gindemo/internal/Config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if Config.GetAPIBearerToken() != c.Request.Header.Get("Authorization") {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
		}
	}
}
