package ConvertModel

import (
	"gindemo/internal/Model/MysqlModel"
	"gindemo/internal/Model/RedisModel"
	"gindemo/internal/Model/ServiceModel"
	. "strings"
)

func RedisConvertInfosServiceModel(parameter []*RedisModel.HistoryInfo) []*ServiceModel.InfoHistoryResponse {
	if parameter == nil {
		return nil
	}

	infos := make([]*ServiceModel.InfoHistoryResponse, len(parameter))
	for _, v := range parameter {
		infos = append(infos, ConvertInfoServiceModel(v))
	}
	return infos
}

func MysqlConvertInfosServiceModel(parameter []*RedisModel.HistoryInfo) []*ServiceModel.InfoHistoryResponse {
	if parameter == nil {
		return nil
	}

	var infos []*ServiceModel.InfoHistoryResponse
	for _, v := range parameter {
		infos = append(infos, ConvertInfoServiceModel(v))
	}
	return infos
}

func ConvertInfoServiceModel(parameter interface{}) *ServiceModel.InfoHistoryResponse {
	switch parameter.(type) {
	case *RedisModel.HistoryInfo:
		value := parameter.(*RedisModel.HistoryInfo)
		return redisInfoConvertServiceModel(value)
	case *MysqlModel.HistoryInfo:
		value := parameter.(*MysqlModel.HistoryInfo)
		return mysqlInfoConvertServiceModel(value)
	}
	return nil
}

func redisInfoConvertServiceModel(parameter *RedisModel.HistoryInfo) *ServiceModel.InfoHistoryResponse {
	if parameter == nil {
		return nil
	}

	return &ServiceModel.InfoHistoryResponse{
		OpenId:     parameter.OpenId,
		VideoId:    parameter.VideoId,
		UseTime:    parameter.UseTime,
		AuthorName: parameter.AuthorName,
		Title:      parameter.Title,
		CoverUrl:   parameter.CoverUrl,
	}
}

func mysqlInfoConvertServiceModel(parameter *MysqlModel.HistoryInfo) *ServiceModel.InfoHistoryResponse {
	if parameter == nil {
		return nil
	}

	title := Split(parameter.VideoHistories.Title, ",")

	return &ServiceModel.InfoHistoryResponse{
		OpenId:     parameter.OpenId,
		VideoId:    parameter.VideoHistories.VideoId,
		UseTime:    int(parameter.VideoHistories.UseTime),
		AuthorName: parameter.AuthorName,
		Title:      &title,
		CoverUrl:   parameter.VideoHistories.CoverUrl,
	}
}
