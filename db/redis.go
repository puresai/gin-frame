package db

import (
	"fmt"
	"context"
	"time"
	
	"github.com/spf13/viper"
	redis "github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password: viper.GetString("redis.auth"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.MaxActive"),DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolTimeout:  30 * time.Second,
	})

	pong, err := RedisClient.Ping(Ctx).Result()
	fmt.Println(pong, err)
}