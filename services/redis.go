package services

import (
	"encoding/json"
	"fmt"
	"time"
	"errors"
	
	"github.com/spf13/viper"
	"github.com/go-redis/redis"

	"github.com/13sai/gin-frame/logging"
)

var RedisClient *redis.Client

func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: viper.GetString("redis.auth"),
		DB:       0,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
}


func LPush(key string, value interface{}) (len int64, err error){
	ret := RedisClient.LPush(key, value)
	if ret.Err() != nil {
		fmt.Println(ret.Err())
		err = errors.New("error")
		logging.Error(value)
	} else {
		len = ret.Val()
	}
	return 
}

func BrPopLPush(key string, bakKey string) (val interface{}, err error){
	val, err = RedisClient.BRPopLPush(key, bakKey, 2*time.Second).Result()
	if err != nil {
		return
	}

	return val, err
}

func SetCache(key string, val interface{}) {
	result, _ := json.Marshal(val)
	_, err := RedisClient.Set(key, result, 0).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func SetCacheWT(key string, val interface{}, ttl int) {
	result, _ := json.Marshal(val)
	_, err := RedisClient.Set(key, result, time.Second*time.Duration(ttl)).Result()
	if err != nil {
		fmt.Println(err)
	}
}

func GetCache(key string) (ret interface{}) {
	if RedisClient.Exists(key).Val() > 0 {
		by, err := RedisClient.Get(key).Bytes();
		if err !=nil {
			fmt.Println("error")
		}

		json.Unmarshal(by, &ret)
		return ret
	}

	return nil
}
