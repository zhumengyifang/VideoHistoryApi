package Task

import (
	"errors"
	"fmt"
	"gindemo/internal/Config"
	"gindemo/internal/Model/MongodbModel"
	"gindemo/internal/Model/RedisModel"
	"gindemo/internal/Model/ServiceModel"
	"gindemo/internal/MysqlUtil"
	"gindemo/internal/RedisUtil"
	"github.com/garyburd/redigo/redis"
	"github.com/json-iterator/go"
	"time"
)

func init() {
	for i := 0; i < Config.SubmitTaskCount(); i++ {
		go StartTask("Submit")
	}

	for i := 0; i < Config.DelTaskCount(); i++ {
		go StartTask("Del")
	}

	for i := 0; i < Config.ClearTaskCount(); i++ {
		go StartTask("Clear")
	}
}

func StartTask(taskType string) {
	for {
		HandleTask(taskType)
	}
}

func HandleTask(taskType string) {
	taskLog := MongodbModel.MongoTaskLog{TaskType: taskType}

	defer func() {
		if err := recover(); err != nil && err != redis.ErrNil {
			taskLog.TaskError = (err.(error)).Error()
			LogForTask <- taskLog
		}
	}()

	message, err := RedisUtil.TaskBRPOP(taskType)
	if err != nil {
		panic(err)
	}

	taskLog.StartTime = time.Now()

	var jsonIterator = jsoniter.ConfigCompatibleWithStandardLibrary
	bytes, err := jsonIterator.Marshal(message)
	if err != nil {
		panic(err)
	}
	taskLog.TaskMessage = string(bytes)

	err = HandleMessage(message)
	if err != nil {
		panic(err)
	}
	taskLog.EndTime = time.Now()
	taskLog.LatencyTime = taskLog.EndTime.Sub(taskLog.StartTime).Milliseconds()
	LogForTask <- taskLog
}

func HandleMessage(message *RedisModel.Task) error {
	switch message.TaskType {
	case "Submit":
		if err := submit(message); err != nil {
			panic(err)
		}
		break
	case "Del":
		if err := del(message); err != nil {
			panic(err)
		}
		break
	case "Clear":
		if err := clear(message); err != nil {
			panic(err)
		}
		break
	}
	return nil
}

func submit(message *RedisModel.Task) error {
	body, ok := (message.TaskMessage).(*RedisModel.HistoryInfo)
	if !ok {
		return errors.New("ConversionFailure")
	}

	err := MysqlUtil.Submit(body)
	if err != nil {
		return err
	}
	return nil
}

func del(message *RedisModel.Task) error {
	body, ok := (message.TaskMessage).(*ServiceModel.DelHistoryParameter)
	if !ok {
		return errors.New("ConversionFailure")
	}

	err := MysqlUtil.Del(body)
	if err != nil {
		return err
	}
	return nil
}

func clear(message *RedisModel.Task) error {
	body, ok := (message.TaskMessage).(*ServiceModel.ClearHistoryParameter)
	if !ok {
		return errors.New("ConversionFailure")
	}

	err := MysqlUtil.Clear(body)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
