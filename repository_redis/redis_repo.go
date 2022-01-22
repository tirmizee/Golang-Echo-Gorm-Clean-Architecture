package repositoryredis

import "github.com/go-redis/redis"

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(c *redis.Client) *RedisRepository {
	return &RedisRepository{client: c}
}
