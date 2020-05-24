package api

import (
	"bytes"
	"gindemo/internal/Model/MongodbModel"
	"gindemo/internal/MongoDbUtil"
	"github.com/gin-gonic/gin"
	"time"
)

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//requestBody := c.Request.Body
		//body, _ := ioutil.ReadAll(c.Request.Body)

		//响应body
		responseBody := blw.body.String()

		mongoLog := MongodbModel.BuildMongoLog(startTime, endTime, latencyTime, reqMethod, reqUri, statusCode, clientIP, "", responseBody)
		_, _ = MongoDbUtil.InsertOne(mongoLog)
	}
}
