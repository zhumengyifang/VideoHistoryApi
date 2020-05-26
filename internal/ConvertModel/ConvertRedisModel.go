package ConvertModel

import (
	"gindemo/internal/Model/MysqlModel"
	"gindemo/internal/Model/RedisModel"
	"gindemo/internal/Model/ServiceModel"
	"strings"
	"time"
)

func ConvertSubmitHistoryRedisModel(parameter *ServiceModel.SubmitHistoryParameter) *RedisModel.HistoryInfo {
	if parameter == nil {
		return nil
	}

	return &RedisModel.HistoryInfo{
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

func MysqlConvertInfoRedisModel(info *MysqlModel.HistoryInfo) []*RedisModel.HistoryInfo {
	var result []*RedisModel.HistoryInfo

	for _, history := range info.VideosHistories {
		title := strings.Split(history.Title, ",")
		result = append(result, &RedisModel.HistoryInfo{
			OpenId:     info.OpenId,
			AuthorName: info.AuthorName,

			Title:      &title,
			VideoId:    history.VideoId,
			UseTime:    int(history.UseTime),
			CoverUrl:   history.CoverUrl,
			SubmitDate: history.SubmitDateTime,
			IsDelete:   history.IsDel,
		})
	}
	return result
}
