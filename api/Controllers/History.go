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
	if err := c.ShouldBindJSON(postBody); err != nil {
		handleError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("please_check_parameter")))
		return
	}

	body, ok := (postBody.Body).(*ServiceModel.InfoHistoryParameter)
	if !ok {
		handleError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
	}

	responseBody := Domain.Info(body)
	handleError(c, *responseBody)
}

/**
提交历史信息
*/
func submit(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.SubmitHistoryParameter{}}
	if err := c.ShouldBindJSON(postBody); err != nil {
		handleError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("please_check_parameter")))
		return
	}

	body, ok := (postBody.Body).(*ServiceModel.SubmitHistoryParameter)
	if !ok {
		handleError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
	}

	responseBody := Domain.Submit(body)
	handleError(c, *responseBody)
}

/**
历史分页
*/
func list(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.ListHistoryParameter{}}
	if err := c.ShouldBindJSON(postBody); err != nil {
		handleError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("please_check_parameter")))
		return
	}

	body, ok := (postBody.Body).(*ServiceModel.ListHistoryParameter)
	if !ok {
		handleError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
	}

	responseBody := Domain.List(body)
	handleError(c, *responseBody)
}

/**
删除具体的历史信息
*/
func del(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.DelHistoryParameter{}}
	if err := c.ShouldBindJSON(postBody); err != nil {
		handleError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("please_check_parameter")))
		return
	}

	body, ok := (postBody.Body).(*ServiceModel.DelHistoryParameter)
	if !ok {
		handleError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
	}

	responseBody := Domain.Del(body)
	handleError(c, *responseBody)
}

/**
清空
*/
func clear(c *gin.Context) {
	postBody := &ServiceModel.PostBody{Body: &ServiceModel.ClearHistoryParameter{}}
	if err := c.ShouldBindJSON(postBody); err != nil {
		handleError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("please_check_parameter")))
		return
	}

	body, ok := (postBody.Body).(*ServiceModel.ClearHistoryParameter)
	if !ok {
		handleError(c, *ApiUtil.BuildErrorApiResponse(500, errors.New("ConvertBodyErr")))
	}

	responseBody := Domain.Clear(body)
	handleError(c, *responseBody)
}
