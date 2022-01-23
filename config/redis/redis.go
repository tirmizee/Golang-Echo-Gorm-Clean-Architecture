package redis

import (
	"clean-architect/config"

	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client {

	redisConfigProp := config.GetRdConfigProp()

	client := redis.NewClient(&redis.Options{
		Addr:     redisConfigProp.Host + ":" + redisConfigProp.Port,
		Password: redisConfigProp.Pass,
		DB:       redisConfigProp.Index,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	return client
}
