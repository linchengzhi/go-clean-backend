package redis

import (
	"github.com/go-redis/redis/v8"
)

type UserRedis struct {
	db *redis.Client
}
