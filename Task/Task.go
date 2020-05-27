package Task

import (
	"errors"
	"fmt"
	"gindemo/internal/Model/RedisModel"
	"gindemo/internal/Model/ServiceModel"
	"gindemo/internal/MysqlUtil"
	"gindemo/internal/RedisUtil"
)

func init() {
	go StartTask("Submit")
	go StartTask("Del")
	go StartTask("Clear")
}

func StartTask(taskType string) {
	for {
		HandelMessage(taskType)
	}
}

func HandelMessage(taskType string) {
	message, err := RedisUtil.TaskRPop(taskType)
	if err != nil {
		return
	}

	switch message.TaskType {
	case "Submit":
		if err := submit(message); err != nil {

		}
		break
	case "Del":
		if err := del(message); err != nil {

		}
		break
	case "Clear":
		if err := clear(message); err != nil {

		}
		break
	}
}

func submit(message *RedisModel.Task) error {
	body, ok := (message.TaskMessage).(*ServiceModel.SubmitHistoryParameter)
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
