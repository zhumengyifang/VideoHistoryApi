package Controllers

import (
	"errors"
	"gindemo/api/ApiUtil"
	"gindemo/api/Middleware"
	"gindemo/internal/Domain"
	"gindemo/internal/Model/ServiceModel"
	"github.com/gin-gonic/gin"
)

func History(engine *gin.Engine) {
	historyGroup := engine.Group("/history", Middleware.Validate(), Middleware.Logger())
	{
		historyGroup.POST("/info", info)
		historyGroup.POST("/submit", submit)
		historyGroup.POST("/list", list)
		historyGroup.POST("/clear", clear)
		historyGroup.POST("/del", del)
	}
}

// @Summary info
// @Description get Video history Info
// @Tags history
// @Security Bearer
// @Accept json
// @Produce  json
// @Param body body ServiceModel.InfoHistoryParameter true "body"
// @Success 200 {object} ServiceModel.InfoHistoryResponse
// @Router /history/info [post]
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

// @Summary submit
// @Description submit Video history Info
// @Tags history
// @Security Bearer
// @Accept json
// @Produce  json
// @Param body body ServiceModel.SubmitHistoryParameter true "body"
// @Success 200
// @Router /history/submit [post]
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

// @Summary list
// @Description get Video history Infos
// @Tags history
// @Security Bearer
// @Accept json
// @Produce  json
// @Param body body ServiceModel.ListHistoryParameter true "body"
// @Success 200 {object} ServiceModel.ListHistoryResponse
// @Router /history/list [post]
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

// @Summary del
// @Description del Video history Info
// @Tags history
// @Security Bearer
// @Accept json
// @Produce  json
// @Param body body ServiceModel.DelHistoryParameter true "body"
// @Success 200 {object} ServiceModel.DelHistoryResponse
// @Router /history/del [post]
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

// @Summary clear
// @Description clear Video history Info
// @Tags history
// @Security Bearer
// @Accept json
// @Produce  json
// @Param body body ServiceModel.ClearHistoryParameter true "body"
// @Success 200
// @Router /history/clear [post]
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
