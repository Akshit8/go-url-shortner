package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// client struct defines client for redis db
type client struct{}

// NewClient creates instance of client and returns a redis client
func NewClient(redisURI string) (*redis.Client, error) {
	options, err := redis.ParseURL(redisURI)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(options)
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
