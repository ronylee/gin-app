package redis

import (
	"gin-app/config"
	"gin-app/pkg/logging"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() {
	conf := config.Bases.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port,
		Password: conf.Password,
		DB:       conf.DB,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		logging.Error("redis connection error: ", err)
		panic(err)
	}
	logging.Info("redis connect ping response:", pong)
	Redis = client
}
