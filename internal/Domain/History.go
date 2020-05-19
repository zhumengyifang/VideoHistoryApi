package Domain

import (
	"errors"
	"gindemo/api/ApiUtil"
	"gindemo/api/ServiceModel"
	"gindemo/internal/InternalUtil"
)

func Info(body *ServiceModel.InfoHistoryParameter) *ServiceModel.ResponseBody {
	result, err := InternalUtil.GetInfo(body)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, errors.New("GetInfoErr"))
	}

	return ApiUtil.BuildApiResponse(result)
}

func Submit(body *ServiceModel.SubmitHistoryParameter) *ServiceModel.ResponseBody {
	err := InternalUtil.SubmitInfo(body)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, errors.New("GetInfoErr"))
	}
	return ApiUtil.BuildApiResponse(nil)
}
