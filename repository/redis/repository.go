package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/Akshit8/url-shortner/url"
	"github.com/go-redis/redis/v8"
	errs "github.com/pkg/errors"
)

type redisRepository struct {
	client *redis.Client
}

// NewRedisRepository creates a new instance of redisRepository
func NewRedisRepository(redisURI string) (url.RedirectRepository, error) {
	repo := &redisRepository{}
	client, err := newRedisClient(redisURI)
	if err != nil {
		return nil, errs.Wrap(err, "repository.NewRedisRepository")
	}
	repo.client = client
	return repo, nil
}

func newRedisClient(redisURI string) (*redis.Client, error) {
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

func (r *redisRepository) generateKey(code string) string {
	return fmt.Sprintf("redirect:%s", code)
}

func (r *redisRepository) Find(code string) (*url.Redirect, error) {
	redirect := &url.Redirect{}
	key := r.generateKey(code)
	data, err := r.client.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, errs.Wrap(err, "repository.Redirect.Find")
	}
	if len(data) == 0 {
		return nil, errs.Wrap(url.ErrRedirectNotFound, "repository.Redirect.Find")
	}
	createdAt, err := time.Parse(time.UnixDate, data["created_at"])
	if err != nil {
		return nil, errs.Wrap(err, "repository.Redirect.Find")
	}
	redirect.Code = data["code"]
	redirect.URL = data["url"]
	redirect.CreatedAt = createdAt
	return redirect, nil
}

func (r *redisRepository) Store(redirect *url.Redirect) error {
	key := r.generateKey(redirect.Code)
	data := map[string]interface{}{
		"code":       redirect.Code,
		"url":        redirect.URL,
		"created_at": redirect.CreatedAt,
	}
	_, err := r.client.HMSet(context.Background(), key, data).Result()
	if err != nil {
		return errs.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}
