package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func HistoryConroller(group *gin.Engine) {
	historyGroup := group.Group("/history")
	{
		historyGroup.POST("/info", info)
		historyGroup.POST("/submit", submit)
		historyGroup.POST("/list", list)
		historyGroup.POST("/clear", clear)
		historyGroup.POST("/delete", delete)
	}
}

/**
历史详细信息
*/
func info(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(body))

	c.JSON(200, gin.H{
		"message": "info",
	})
}

/**
提交历史信息
*/
func submit(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "submit",
	})
}

/**
历史分页
*/
func list(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}

/**
清空
*/
func clear(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "clear",
	})
}

/**
删除具体的历史信息
*/
func delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete",
	})
}
