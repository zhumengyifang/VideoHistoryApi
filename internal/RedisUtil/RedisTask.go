package RedisUtil

import (
	"encoding/json"
	"gindemo/internal/Model/RedisModel"
	"gindemo/internal/Model/ServiceModel"
	"github.com/garyburd/redigo/redis"
)

func TaskBRPOP(taskType string) (*RedisModel.Task, error) {
	conn := pool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("BRPop", taskType+"Task", 2))
	if err != nil {
		return nil, err
	}

	var key string
	var value []byte

	if _, err := redis.Scan(values, &key, &value); err != nil {
		return nil, err
	}

	result := RedisModel.Task{}
	switch taskType {
	case "Submit":
		result.TaskMessage = new(RedisModel.HistoryInfo)
		break
	case "Del":
		result.TaskMessage = new(ServiceModel.DelHistoryParameter)
		break
	case "Clear":
		result.TaskMessage = new(ServiceModel.ClearHistoryParameter)
		break
	}

	if err = json.Unmarshal(value, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func DelCommand(key string, removes []string) error {
	if removes == nil {
		return nil
	}

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HDel", redis.Args{}.Add(key).AddFlat(removes)...)
	if err != nil {
		return err
	}

	return nil
}
