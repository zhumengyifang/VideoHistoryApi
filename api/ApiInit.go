package api

import (
	"context"
	"gindemo/api/Controllers"
	"github.com/gin-gonic/gin"
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
	router := gin.Default()

	router.Use(Validate())
	Router(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
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

func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if bearerToken != c.Request.Header.Get("Authorization") {
			c.Abort()
			c.JSON(http.StatusUnauthorized, nil)
		}
	}
}

func Router(router *gin.Engine) {
	Controllers.History(router)
}
