package dependencies

import (
	"github.com/go-redis/redis"
)

func NewRedisClient(config *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
