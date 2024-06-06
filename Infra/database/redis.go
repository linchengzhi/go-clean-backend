package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func NewRedis(host, password string, db int) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password, // no password set
		DB:       db,       // use default Db
	})

	err := rdb.Ping(context.TODO()).Err()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
