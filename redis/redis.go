package redis

import (
	"github.com/go-redis/redis"
)

func Connect(db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: db,
	})
}
