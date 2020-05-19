package Controllers

import (
	"errors"
	"fmt"
	"gindemo/api/ApiUtil"
	"gindemo/api/ServiceModel"
	"gindemo/internal/Domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func History(router *gin.Engine) {
	historyGroup := router.Group("/history")
	{
		historyGroup.POST("/info", info)
		historyGroup.POST("/submit", submit)
		historyGroup.POST("/list", list)
		historyGroup.POST("/clear", clear)
		historyGroup.POST("/del", del)
	}
}

/**
历史详细信息
*/
func info(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.InfoHistoryParameter{}}
	err := c.Bind(&postBody)
	if err != nil {
		HandelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertPostBodyErr")))
		return
	}

	body, ok := (postBody.Body).(*ServiceModel.InfoHistoryParameter)
	if !ok {
		HandelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
		return
	}

	responseBody := Domain.Info(body)
	HandelError(c, *responseBody)
}

/**
提交历史信息
*/
func submit(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.SubmitHistoryParameter{}}
	err := c.Bind(&postBody)
	if err != nil {
		HandelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertPostBodyErr")))
		return
	}

	body, ok := (postBody.Body).(*ServiceModel.SubmitHistoryParameter)
	if !ok {
		HandelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
		return
	}

	responseBody := Domain.Submit(body)
	HandelError(c, *responseBody)
}

/**
历史分页
*/
func list(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.ListHistoryParameter{}}
	err := c.Bind(&postBody)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, postBody)
}

/**
清空
*/
func clear(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.ClearHistoryParameter{}}
	err := c.Bind(&postBody)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, postBody)
}

/**
删除具体的历史信息
*/
func del(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.DeleteHistoryParameter{}}
	err := c.Bind(&postBody)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, postBody)
}
