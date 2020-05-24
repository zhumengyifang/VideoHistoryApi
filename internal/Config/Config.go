package Config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type ApiConfig struct {
	APIBearerToken string      `json:"apiBearerToken"`
	APIPort        int         `json:"apiPort"`
	Redis          RedisConfig `json:"redis"`
	Mongo          MongoConfig `json:"mongo"`
}

type RedisConfig struct {
	IPAddress    string `json:"ipAddress"`
	Password     string `json:"password"`
	ProtocolType string `json:"protocolType"`
	MaxIdle      int    `json:"maxIdle"`
	IdleTimeout  int    `json:"idleTimeout"`
}

type MongoConfig struct {
	MongoAddr string `json:"mongoAddr"`
	MaxIdle   uint64 `json:"maxIdle"`
}

var apiConfig ApiConfig

func init() {
	file, err := os.Open("./Config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&apiConfig)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func GetAPIBearerToken() string {
	return apiConfig.APIBearerToken
}

func GetAPIPort() string {
	return ":" + strconv.Itoa(apiConfig.APIPort)
}

func GetRedis() RedisConfig {
	return apiConfig.Redis
}

func GetMongo() MongoConfig {
	return apiConfig.Mongo
}
