package Domain

import (
	"gindemo/api/ApiUtil"
	"gindemo/internal/ConvertModel"
	"gindemo/internal/Model/RedisModel"
	"gindemo/internal/Model/ServiceModel"
	"gindemo/internal/MysqlUtil"
	"gindemo/internal/RedisUtil"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"sort"
	"time"
)

func Info(body *ServiceModel.InfoHistoryParameter) *ServiceModel.ResponseBody {
	result, err := RedisUtil.GetInfo(body)
	if err != nil && err != redis.ErrNil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	if result != nil && !result.IsDelete {
		return ApiUtil.BuildApiResponse(ConvertModel.ConvertInfoServiceModel(result))
	}

	result1, err := MysqlUtil.Info(body)
	if err != nil && err != gorm.ErrRecordNotFound {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	return ApiUtil.BuildApiResponse(ConvertModel.ConvertInfoServiceModel(result1))
}

func Submit(body *ServiceModel.SubmitHistoryParameter) *ServiceModel.ResponseBody {
	redisInfo := ConvertModel.ConvertSubmitHistoryRedisModel(body)
	err := RedisUtil.SubmitInfo(redisInfo)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	task := RedisModel.Task{TaskType: "Submit", TaskMessage: redisInfo}
	if err = RedisUtil.TaskLPush(task); err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	return ApiUtil.BuildApiResponse(nil)
}

func List(body *ServiceModel.ListHistoryParameter) *ServiceModel.ResponseBody {
	result := ServiceModel.ListHistoryResponse{PageCount: body.PageCount}

	redisInfos, err := RedisUtil.GetAllSlice(body.OpenId)
	if err != nil && err != redis.ErrNil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	sort.Sort(RedisModel.HistoryInfos(redisInfos))

	excludeCount := (body.PageCount - 1) * body.PageSize
	if len(redisInfos) > excludeCount {
		if len(redisInfos) >= excludeCount+body.PageSize {
			redisInfos = redisInfos[excludeCount : excludeCount+body.PageSize]
		} else {
			redisInfos = redisInfos[excludeCount:]
		}
	} else {
		redisInfos = []*RedisModel.HistoryInfo{}
	}

	if len(redisInfos) < body.PageSize {
		mysqlInfos, err := MysqlUtil.List(body)
		if err != nil && err != gorm.ErrRecordNotFound {
			return ApiUtil.BuildErrorApiResponse(500, err)
		}
		//合并
		redisInfos = append(redisInfos, ConvertModel.MysqlConvertInfoRedisModel(mysqlInfos)...)
	}

	//去重
	redisInfos = distinctVideoId(redisInfos)

	if len(redisInfos) >= body.PageSize {
		redisInfos = redisInfos[0:body.PageSize]
	} else {
		redisInfos = redisInfos[0:]
	}

	//转换
	result.Videos = ConvertModel.RedisConvertInfosServiceModel(redisInfos)
	result.PageSize = len(result.Videos)
	return ApiUtil.BuildApiResponse(result)
}

func distinctVideoId(infos []*RedisModel.HistoryInfo) []*RedisModel.HistoryInfo {
	m := map[string]string{}
	var result []*RedisModel.HistoryInfo
	for _, info := range infos {
		if _, ok := m[info.VideoId]; !ok {
			m[info.VideoId] = info.VideoId
			result = append(result, info)
		}
	}
	return result
}

func Clear(body *ServiceModel.ClearHistoryParameter) *ServiceModel.ResponseBody {
	infos, err := RedisUtil.GetAllMap(body.OpenId)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	isDel := make(map[string][]byte)
	for k, v := range infos {
		if v.IsDelete {
			continue
		}
		v.IsDelete = true
		var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
		bytes, err := jsonIterator.Marshal(v)
		if err != nil {
			continue
		}
		isDel[k] = bytes
	}

	err = RedisUtil.Del(body.OpenId, isDel)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	task := RedisModel.Task{TaskType: "Clear", TaskMessage: body}
	if err = RedisUtil.TaskLPush(task); err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	return ApiUtil.BuildApiResponse(nil)
}

func Del(body *ServiceModel.DelHistoryParameter) *ServiceModel.ResponseBody {
	infos, err := RedisUtil.GetAllMap(body.OpenId)
	if err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	result := ServiceModel.DelHistoryResponse{OpenId: body.OpenId, DeleteInfo: make(map[string]bool)}

	isDel := make(map[string][]byte)
	isSave := make(map[string][]byte)

	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	for _, v := range body.VideoIds {
		result.DeleteInfo[v] = true

		if _, ok := infos[v]; !ok {
			save := RedisModel.HistoryInfo{OpenId: body.OpenId, VideoId: v, IsDelete: true, SubmitDate: time.Now()}
			bytes, err := jsonIterator.Marshal(save)
			if err != nil {
				continue
			}
			isSave[v] = bytes
			continue
		}

		del := infos[v]
		if del.IsDelete {
			continue
		}
		del.IsDelete = true
		bytes, err := jsonIterator.Marshal(del)
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

	task := RedisModel.Task{TaskType: "Del", TaskMessage: body}
	if err = RedisUtil.TaskLPush(task); err != nil {
		return ApiUtil.BuildErrorApiResponse(500, err)
	}

	return ApiUtil.BuildApiResponse(result)
}
