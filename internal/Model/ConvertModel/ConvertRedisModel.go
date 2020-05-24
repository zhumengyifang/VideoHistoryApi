package ConvertModel

import (
	"gindemo/internal/Model/RedisModel"
	"gindemo/internal/Model/ServiceModel"
	"time"
)

func ConvertSubmitHistoryRedisModel(parameter *ServiceModel.SubmitHistoryParameter) *RedisModel.HistoryInfoParameter {
	if parameter == nil {
		return nil
	}

	return &RedisModel.HistoryInfoParameter{
		OpenId:     parameter.OpenId,
		VideoId:    parameter.VideoId,
		UseTime:    parameter.UseTime,
		AuthorName: parameter.AuthorName,
		Title:      parameter.Title,
		CoverUrl:   parameter.CoverUrl,

		SubmitDate: time.Now(),
		IsDelete:   false,
	}
}
