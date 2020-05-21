package RedisUtil

import (
	"encoding/json"
	"fmt"
	"gindemo/api/ServiceModel"
	"gindemo/internal/RedisUtil/RedisModel"
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	MaxIdle      = 10
	IdleTimeout  = 240 * time.Second
	ipAddress    = "192.168.170.137:6379"
	protocolType = "tcp"
	passWord     = "myredis123"
)

var pool *redis.Pool

func init() {
	pool = newPool()
}

/**
创建线程池
*/
func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     MaxIdle,
		IdleTimeout: IdleTimeout,
		Dial:        func() (redis.Conn, error) { return redis.Dial(protocolType, ipAddress, redis.DialPassword(passWord)) },
	}
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

func SubmitInfo(parameter *RedisModel.HistoryInfoParameter) error {
	conn := pool.Get()
	defer conn.Close()

	bytes, err := json.Marshal(parameter)
	if err != nil {
		return err
	}

	if _, err := conn.Do("HSet", parameter.OpenId, parameter.VideoId, bytes); err != nil {
		return err
	}

	return nil
}

func GetInfo(parameter *ServiceModel.InfoHistoryParameter) (*RedisModel.HistoryInfoParameter, error) {
	conn := pool.Get()
	defer conn.Close()

	v, err := redis.Bytes(conn.Do("HGet", parameter.OpenId, parameter.VideoId))
	if err != nil {
		return nil, err
	}

	result := new(RedisModel.HistoryInfoParameter)
	if err = json.Unmarshal(v, result); err != nil {
		return nil, err
	}

	return result, nil
}

//func GetALl(parameter *ServiceModel.DeleteHistoryParameter) ([]*RedisModel.HistoryInfoParameter, error) {
//	conn := pool.Get()
//	defer conn.Close()
//
//	values, err := redis.Values(conn.Do("HVALS", parameter.OpenId))
//	if err != nil {
//		return nil, err
//	}
//
//	var infos []*RedisModel.HistoryInfoParameter
//	for _, v := range values {
//		result := new(RedisModel.HistoryInfoParameter)
//		if err = json.Unmarshal(v.([]uint8), result); err != nil {
//			fmt.Println(err)
//		}
//		infos = append(infos, result)
//	}
//
//	return infos, nil
//}

func GetALl(key string) (map[string]*RedisModel.HistoryInfoParameter, error) {
	conn := pool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HVALS", key))
	if err != nil {
		return nil, err
	}

	infos := make(map[string]*RedisModel.HistoryInfoParameter)
	for _, v := range values {
		result := new(RedisModel.HistoryInfoParameter)
		if err = json.Unmarshal(v.([]uint8), result); err != nil {
			fmt.Println(err)
		}
		infos[result.VideoId] = result
	}

	return infos, nil
}

func Del(key string, isDel map[string][]byte) error {
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HMSet", redis.Args{}.Add(key).AddFlat(isDel)...)
	if err != nil {
		return err
	}

	return nil
}
