package Controllers

import (
	"errors"
	"gindemo/api/ApiUtil"
	"gindemo/internal/Domain"
	"gindemo/internal/Model/ServiceModel"
	"github.com/gin-gonic/gin"
)

func History(engine *gin.Engine) {
	historyGroup := engine.Group("/history")
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
		handelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertPostBodyErr")))
	}

	body, ok := (postBody.Body).(*ServiceModel.InfoHistoryParameter)
	if !ok {
		handelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
	}

	responseBody := Domain.Info(body)
	handelError(c, *responseBody)
}

/**
提交历史信息
*/
func submit(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.SubmitHistoryParameter{}}
	err := c.Bind(&postBody)
	if err != nil {
		handelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertPostBodyErr")))
	}

	body, ok := (postBody.Body).(*ServiceModel.SubmitHistoryParameter)
	if !ok {
		handelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
	}

	responseBody := Domain.Submit(body)
	handelError(c, *responseBody)
}

/**
历史分页
*/
func list(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.ListHistoryParameter{}}
	err := c.Bind(&postBody)
	if err != nil {
		handelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertPostBodyErr")))
	}

	body, ok := (postBody.Body).(*ServiceModel.ListHistoryParameter)
	if !ok {
		handelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
	}

	responseBody := Domain.List(body)
	handelError(c, *responseBody)
}

/**
删除具体的历史信息
*/
func del(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.DelHistoryParameter{}}
	err := c.Bind(&postBody)
	if err != nil {
		handelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertPostBodyErr")))
	}

	body, ok := (postBody.Body).(*ServiceModel.DelHistoryParameter)
	if !ok {
		handelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
	}

	responseBody := Domain.Del(body)
	handelError(c, *responseBody)
}

/**
清空
*/
func clear(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.ClearHistoryParameter{}}
	err := c.Bind(&postBody)
	if err != nil {
		handelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertPostBodyErr")))
	}

	body, ok := (postBody.Body).(*ServiceModel.ClearHistoryParameter)
	if !ok {
		handelError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
	}

	responseBody := Domain.Clear(body)
	handelError(c, *responseBody)
}
