package api

import (
	"gindemo/api/ServiceModel"
	"github.com/gin-gonic/gin"
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
	postBody, err := ServiceModel.Convert(c, &ServiceModel.PostBody{Body: &ServiceModel.InfoHistoryParameter{}})
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"body": postBody,
	})
}

/**
提交历史信息
*/
func submit(c *gin.Context) {
	postBody, err := ServiceModel.Convert(c, &ServiceModel.PostBody{Body: &ServiceModel.SubmitHistoryParameter{}})
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"body": postBody,
	})
}

/**
历史分页
*/
func list(c *gin.Context) {
	postBody, err := ServiceModel.Convert(c, &ServiceModel.PostBody{Body: &ServiceModel.ListHistoryParameter{}})
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"body": postBody,
	})
}

/**
清空
*/
func clear(c *gin.Context) {
	postBody, err := ServiceModel.Convert(c, &ServiceModel.PostBody{Body: &ServiceModel.ClearHistoryParameter{}})
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"body": postBody,
	})
}

/**
删除具体的历史信息
*/
func delete(c *gin.Context) {
	postBody, err := ServiceModel.Convert(c, &ServiceModel.PostBody{Body: &ServiceModel.DeleteHistoryParameter{}})
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"body": postBody,
	})
}
