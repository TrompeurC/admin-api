package utils

import (
	"admin-api/constant"
	"admin-api/pkg/redis"
	"log"
	"time"
)

type RedisStore struct {
}

func (r RedisStore) Set(id, val string) error {
	key := constant.LOGIN_CODE + "-" + id
	err := redis.RedisDB.Set(key, val, 5*time.Minute).Err()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (r RedisStore) Get(id string, clear bool) string {
	key := constant.LOGIN_CODE + "-" + id
	result, err := redis.RedisDB.Get(key).Result()
	if err != nil {
		return ""
	}
	return result
}

func (r RedisStore) Verify(id, answer string, clear bool) bool {
	val := (&RedisStore{}).Get(id, clear)
	return val == answer
}
