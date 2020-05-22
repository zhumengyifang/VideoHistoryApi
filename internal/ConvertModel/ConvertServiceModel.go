package ConvertModel

import (
	"gindemo/api/ServiceModel"
	"gindemo/internal/RedisUtil/RedisModel"
)

func ConvertGetInfoServiceModel(parameter *RedisModel.HistoryInfoParameter) *ServiceModel.InfoHistoryResponse {
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
