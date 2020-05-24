package Domain

import (
	"encoding/json"
	"errors"
	"gindemo/api/ApiUtil"
	"gindemo/internal/ConvertModel"
	"gindemo/internal/InternalUtil"
	"gindemo/internal/Model/RedisModel"
	"gindemo/internal/Model/ServiceModel"
	"gindemo/internal/RedisUtil"
	"time"
)

func Info(body *ServiceModel.InfoHistoryParameter) *ServiceModel.ResponseBody {
	result, err := RedisUtil.GetInfo(body)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, errors.New("GetInfoErr"))
	}

	if result != nil && result.IsDelete {
		return ApiUtil.BuildErrorApiResponse(500, errors.New("TheVideoHasBeenDeleted"))
	}

	return ApiUtil.BuildApiResponse(ConvertModel.ConvertGetInfoServiceModel(&RedisModel.HistoryInfoParameter{OpenId: body.OpenId}))
}

func Submit(body *ServiceModel.SubmitHistoryParameter) *ServiceModel.ResponseBody {
	err := RedisUtil.SubmitInfo(ConvertModel.ConvertSubmitHistoryRedisModel(body))
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, errors.New("GetInfoErr"))
	}
	return ApiUtil.BuildApiResponse(nil)
}

func List(body *ServiceModel.ListHistoryParameter) *ServiceModel.ResponseBody {
	result := ServiceModel.ListHistoryResponse{PageSize: body.PageSize, PageCount: body.PageCount}

	infos, err := RedisUtil.GetALl1(body.OpenId)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	videoCount := body.PageCount * body.PageSize
	if len(infos) <= videoCount {
		InternalUtil.SortBySubmitDate(infos)
	} else {

	}

	result.Videos = ConvertModel.ConvertGetInfosServiceModel(infos)
	return ApiUtil.BuildApiResponse(result)
}

func Clear(body *ServiceModel.ClearHistoryParameter) *ServiceModel.ResponseBody {
	infos, err := RedisUtil.GetALl(body.OpenId)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	result := ServiceModel.ClearHistoryResponse{Count: 0}
	isDel := make(map[string][]byte)
	for k, v := range infos {
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
	isSave := make(map[string][]byte)
	for _, v := range *body.VideoIds {
		result.DeleteInfo[v] = true

		if _, ok := infos[v]; !ok {
			save := RedisModel.HistoryInfoParameter{OpenId: body.OpenId, VideoId: v, IsDelete: true, SubmitDate: time.Now()}
			bytes, err := json.Marshal(save)
			if err != nil {
				continue
			}
			isSave[v] = bytes
			continue
		}

		del := infos[v]
		del.IsDelete = true

		bytes, err := json.Marshal(del)
		if err != nil {
			continue
		}
		isDel[v] = bytes
	}

	if err = RedisUtil.Del(body.OpenId, isDel); err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	if err = RedisUtil.SaveInfos(body.OpenId, isSave); err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	return ApiUtil.BuildApiResponse(result)
}
