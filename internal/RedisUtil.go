package internal

import (
	"encoding/json"
	"gindemo/api/ServiceModel"
	"github.com/garyburd/redigo/redis"
)

const ipAddress = "192.168.170.137:6379"
const protocolType = "tcp"
const passWord = "myredis123"

func SubmitInfo(parameter *ServiceModel.SubmitHistoryParameter) error {
	conn, err := createConn(protocolType, ipAddress, passWord)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(parameter)
	if err != nil {
		return err
	}

	if _, err := conn.Do("Hset", parameter.OpenId, parameter.VideoId, bytes); err != nil {
		return err
	}

	return nil
}

func GetInfo(parameter *ServiceModel.InfoHistoryParameter) (*ServiceModel.InfoHistoryResponse, error) {
	conn, err := createConn(protocolType, ipAddress, passWord)
	if err != nil {
		return nil, err
	}

	v, err := redis.Bytes(conn.Do("Hget", parameter.OpenId, parameter.VideoId))
	if err != nil {
		return nil, err
	}

	result := new(ServiceModel.InfoHistoryResponse)
	if err = json.Unmarshal(v, result); err != nil {
		return nil, err
	}

	return result, nil
}

/**
创建redis连接
*/
func createConn(protocolType string, redisAddress string, password string) (redis.Conn, error) {
	var conn redis.Conn
	var err error
	if password == "" {
		conn, err = redis.Dial(protocolType, redisAddress)
	} else {
		conn, err = redis.Dial(protocolType, redisAddress, redis.DialPassword(password))
	}
	if err != nil {
		return nil, err
	}

	return conn, nil
}
