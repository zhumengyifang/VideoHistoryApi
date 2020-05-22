package Domain

import (
	"encoding/json"
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
		return ApiUtil.BuildErrorApiResponse(200, errors.New("TheVideoHasBeenDeleted"))
	}

	return ApiUtil.BuildApiResponse(ConvertModel.ConvertGetInfoServiceModel(result))
}

func Submit(body *ServiceModel.SubmitHistoryParameter) *ServiceModel.ResponseBody {
	err := RedisUtil.SubmitInfo(ConvertModel.ConvertSubmitHistoryRedisModel(body))
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, errors.New("GetInfoErr"))
	}
	return ApiUtil.BuildApiResponse(nil)
}

func List(body *ServiceModel.ListHistoryParameter) *ServiceModel.ResponseBody {
	return ApiUtil.BuildApiResponse(nil)
}

func Clear(body *ServiceModel.ClearHistoryParameter) *ServiceModel.ResponseBody {
	infos, err := RedisUtil.GetALl(body.OpenId)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	result := ServiceModel.ClearHistoryResponse{Count: 0}
	isDel := make(map[string][]byte)
	for k, v := range infos {
		if v.IsDelete {
			continue
		}
		v.IsDelete = true
		bytes, err := json.Marshal(v)
		if err != nil {
			continue
		}
		result.Count++
		isDel[k] = bytes
	}

	err = RedisUtil.Del(body.OpenId, isDel)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	return ApiUtil.BuildApiResponse(result)
}

func Del(body *ServiceModel.DelHistoryParameter) *ServiceModel.ResponseBody {
	infos, err := RedisUtil.GetALl(body.OpenId)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	result := ServiceModel.DelHistoryResponse{OpenId: body.OpenId, DeleteInfo: make(map[string]bool)}
	isDel := make(map[string][]byte)
	for _, v := range *body.VideoIds {
		result.DeleteInfo[v] = false
		if _, ok := infos[v]; !ok {
			continue
		}

		value := infos[v]
		if value.IsDelete {
			result.DeleteInfo[v] = true
			continue
		}

		value.IsDelete = true
		bytes, err := json.Marshal(value)
		if err != nil {
			continue
		}
		isDel[v] = bytes
		result.DeleteInfo[v] = true
	}

	err = RedisUtil.Del(body.OpenId, isDel)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	return ApiUtil.BuildApiResponse(result)
}
