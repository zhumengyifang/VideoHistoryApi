package MongodbModel

import (
	"time"
)

type MongoLog struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	//单位(毫秒)ms
	LatencyTime  int64  `json:"latency_time"`
	ReqMethod    string `json:"req_method"`
	ReqUri       string `json:"req_uri"`
	StatusCode   int    `json:"status_code"`
	ClientIP     string `json:"client_ip"`
	RequestBody  string `json:"request_body"`
	ResponseBody string `json:"response_body"`
}

func BuildMongoLog(startTime time.Time, endTime time.Time, latencyTime int64, reqMethod string, reqUrl string, statusCode int, clientIP string, requestBody string, responseBody string) *MongoLog {
	return &MongoLog{
		StartTime:    startTime,
		EndTime:      endTime,
		LatencyTime:  latencyTime,
		ReqMethod:    reqMethod,
		ReqUri:       reqUrl,
		StatusCode:   statusCode,
		ClientIP:     clientIP,
		RequestBody:  requestBody,
		ResponseBody: responseBody,
	}
}
