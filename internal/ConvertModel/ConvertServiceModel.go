package ConvertModel

import (
	"gindemo/internal/Model/RedisModel"
	"gindemo/internal/Model/ServiceModel"
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

func ConvertGetInfosServiceModel(parameter []*RedisModel.HistoryInfoParameter) []*ServiceModel.InfoHistoryResponse {
	if parameter == nil {
		return nil
	}

	var infos []*ServiceModel.InfoHistoryResponse
	for _, v := range parameter {
		infos = append(infos, ConvertGetInfoServiceModel(v))
	}
	return infos
}
