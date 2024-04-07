package redis

import (
	"admin-api/common/config"
	"github.com/go-redis/redis"
)

var RedisDB *redis.Client

func SetupRedisDB() error {
	redisConfig := config.Config.Redis
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       0,
	})
	_, err := RedisDB.Ping().Result()
	if err != nil {
		panic(err)
	}
	return nil
}
