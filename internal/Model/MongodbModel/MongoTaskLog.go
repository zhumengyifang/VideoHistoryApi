package MongodbModel

import "time"

type MongoTaskLog struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	//单位(毫秒)ms
	LatencyTime int64  `json:"latency_time"`
	TaskType    string `json:"task_type"`
	TaskMessage string `json:"task_message"`
	TaskError   string `json:"task_error"`
}

func BuildMongoTaskLog(startTime time.Time, endTime time.Time, latencyTime int64, taskType string, taskMessage string, taskError string) *MongoTaskLog {
	return &MongoTaskLog{
		StartTime:   startTime,
		EndTime:     endTime,
		LatencyTime: latencyTime,
		TaskType:    taskType,
		TaskMessage: taskMessage,
		TaskError:   taskError,
	}
}
