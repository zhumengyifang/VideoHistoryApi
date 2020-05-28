package api

import (
	"bytes"
	"gindemo/internal/Model/MongodbModel"
	"gindemo/internal/MongoDbUtil"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 请求request
		requestBody, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		//记录响应参数
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

		//响应body
		responseBody := blw.body.String()

		mongoLog := MongodbModel.BuildMongoApiLog(startTime, endTime, latencyTime.Milliseconds(), reqMethod, reqUri, statusCode, clientIP, string(requestBody), responseBody)
		_, _ = MongoDbUtil.InsertApiLog(mongoLog)
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
