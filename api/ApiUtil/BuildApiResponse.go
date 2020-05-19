package ApiUtil

import "gindemo/api/ServiceModel"

func BuildErrorApiResponse(code int, err error) *ServiceModel.ResponseBody {
	return &ServiceModel.ResponseBody{
		Header: ServiceModel.ResponseHeader{
			Version: 1, IsSuccess: false, Error: &ServiceModel.ResponseError{
				Code: code, Message: err.Error(),
			},
		},
	}
}

func BuildApiResponse(result interface{}) *ServiceModel.ResponseBody {
	return &ServiceModel.ResponseBody{
		Header: ServiceModel.ResponseHeader{
			Version: 1, IsSuccess: true, Error: nil,
		},
		Body: result,
	}
}
