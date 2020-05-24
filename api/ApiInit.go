package api

import (
	"context"
	"fmt"
	"gindemo/api/Controllers"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const bearerToken = "Bearer welcome"

/**
api初始化
*/
func init() {
	engine := gin.Default()

	engine.Use(validate())
	engine.Use(logger())

	Router(engine)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if bearerToken != c.Request.Header.Get("Authorization") {
			c.Abort()
			c.JSON(http.StatusUnauthorized, nil)
		}
	}
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

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

		requestBody := c.Request.Body
		body, _ := ioutil.ReadAll(c.Request.Body)

		fmt.Println(latencyTime, reqMethod, reqUri, statusCode, clientIP, requestBody, body)
	}
}

func Router(engine *gin.Engine) {
	Controllers.History(engine)
}
