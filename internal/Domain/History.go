package Domain

import (
	"errors"
	"gindemo/api/ApiUtil"
	"gindemo/api/ServiceModel"
	"gindemo/internal/ConvertModel"
	"gindemo/internal/RedisUtil"
)

func Info(body *ServiceModel.InfoHistoryParameter) *ServiceModel.ResponseBody {
	result, err := RedisUtil.GetInfo(body)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, errors.New("GetInfoErr"))
	}

	if result != nil && result.IsDelete {
		return ApiUtil.BuildErrorApiResponse(500, errors.New("TheVideoHasBeenDeleted"))
	}

	return ApiUtil.BuildApiResponse(ConvertModel.ConverGetInfoServiceModel(result))
}

func Submit(body *ServiceModel.SubmitHistoryParameter) *ServiceModel.ResponseBody {
	err := RedisUtil.SubmitInfo(ConvertModel.ConverSubmitHistoryRedisModel(body))
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, errors.New("GetInfoErr"))
	}
	return ApiUtil.BuildApiResponse(nil)
}

func List(body *ServiceModel.ListHistoryParameter) *ServiceModel.ResponseBody {
	return ApiUtil.BuildApiResponse(nil)
}

func Clear(body *ServiceModel.ClearHistoryParameter) *ServiceModel.ResponseBody {
	return ApiUtil.BuildApiResponse(nil)
}

func Del(body *ServiceModel.DeleteHistoryParameter) *ServiceModel.ResponseBody {
	_, err := RedisUtil.GetInfos(body)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}
	return ApiUtil.BuildApiResponse(nil)
}
