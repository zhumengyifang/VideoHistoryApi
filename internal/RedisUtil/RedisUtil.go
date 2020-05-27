package RedisUtil

import (
	"encoding/json"
	"fmt"
	"gindemo/internal/Config"
	"gindemo/internal/Model/RedisModel"
	"gindemo/internal/Model/ServiceModel"
	"github.com/garyburd/redigo/redis"
	"time"
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
		MaxActive:   Config.GetRedis().MaxActive,
		MaxIdle:     Config.GetRedis().MaxIdle,
		IdleTimeout: time.Duration(Config.GetRedis().IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial(
				Config.GetRedis().ProtocolType,
				Config.GetRedis().Host,
				redis.DialPassword(Config.GetRedis().Password))
		},
		Wait: true,
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

func SubmitInfo(parameter *RedisModel.HistoryInfo) error {
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

func GetInfo(parameter *ServiceModel.InfoHistoryParameter) (*RedisModel.HistoryInfo, error) {
	conn := pool.Get()
	defer conn.Close()

	v, err := redis.Bytes(conn.Do("HGet", parameter.OpenId, parameter.VideoId))
	if err != nil {
		return nil, err
	}

	result := new(RedisModel.HistoryInfo)
	if err = json.Unmarshal(v, result); err != nil {
		return nil, err
	}

	return result, nil
}

func GetAllMap(key string) (map[string]*RedisModel.HistoryInfo, error) {
	conn := pool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HVALS", key))
	if err != nil {
		return nil, err
	}

	infos := make(map[string]*RedisModel.HistoryInfo)
	for _, v := range values {
		result := new(RedisModel.HistoryInfo)
		if err = json.Unmarshal(v.([]uint8), result); err != nil {
			fmt.Println(err)
		}

		infos[result.VideoId] = result
	}

	return infos, nil
}

func GetAllSlice(key string) ([]*RedisModel.HistoryInfo, error) {
	conn := pool.Get()
	defer conn.Close()

	values, err := redis.Values(conn.Do("HVALS", key))
	if err != nil {
		return nil, err
	}

	var infos []*RedisModel.HistoryInfo
	for _, v := range values {
		result := new(RedisModel.HistoryInfo)
		if err = json.Unmarshal(v.([]uint8), result); err != nil {
			fmt.Println(err)
		}

		if !result.IsDelete {
			infos = append(infos, result)
		}
	}

	return infos, nil
}

func Del(key string, isDel map[string][]byte) error {
	if len(isDel) == 0 {
		return nil
	}

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HMSet", redis.Args{}.Add(key).AddFlat(isDel)...)
	if err != nil {
		return err
	}

	return nil
}

func SaveInfos(key string, isSave map[string][]byte) error {
	if len(isSave) == 0 {
		return nil
	}

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("HSet", redis.Args{}.Add(key).AddFlat(isSave)...)
	if err != nil {
		return err
	}
	return nil
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

func Clear(removes string) error {
	if removes == "" {
		return nil
	}

	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("Del", removes)
	if err != nil {
		return err
	}

	return nil
}
