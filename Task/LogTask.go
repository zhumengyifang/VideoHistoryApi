package Task

import (
	"gindemo/internal/Model/MongodbModel"
	"gindemo/internal/MongoDbUtil"
	"github.com/garyburd/redigo/redis"
)

var LogForApi chan MongodbModel.MongoApiLog
var LogForTask chan MongodbModel.MongoTaskLog

func init() {
	LogForApi = make(chan MongodbModel.MongoApiLog, 50)
	LogForTask = make(chan MongodbModel.MongoTaskLog, 50)

	go insertApiLog()
	go insertTaskLog()
}

func insertApiLog() {
	defer func() {
		if err := recover(); err != nil && err != redis.ErrNil {
			taskLog := MongodbModel.MongoTaskLog{TaskType: "System Error"}
			taskLog.TaskError = (err.(error)).Error()
			MongoDbUtil.InsertTaskLog(taskLog)
			insertApiLog()
		}
	}()

	for {
		select {
		case t := <-LogForApi:
			go MongoDbUtil.InsertApiLog(t)
		}
	}
}

func insertTaskLog() {
	defer func() {
		if err := recover(); err != nil && err != redis.ErrNil {
			taskLog := MongodbModel.MongoTaskLog{TaskType: "System Error"}
			taskLog.TaskError = (err.(error)).Error()
			MongoDbUtil.InsertTaskLog(taskLog)
			insertApiLog()
		}
	}()

	for {
		select {
		case t := <-LogForTask:
			go MongoDbUtil.InsertTaskLog(t)
		}
	}
}
