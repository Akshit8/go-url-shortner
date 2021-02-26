package redis

import (
	"github.com/Akshit8/url-shortner/pkg/url"
	"github.com/go-redis/redis/v8"
)

type urlRepository struct {
	client *redis.Client
}

// NewURLRepository creates new instance of urlRepository
func NewURLRepository(client *redis.Client) urls.Repository {
	return &urlRepository{client: client}
}

func (u *urlRepository) Save(url *urls.URL) (*urls.URL, error) {
	return nil, nil
}

func (u *urlRepository) Get(code string) (*urls.URL, error) {
	return nil, nil
}
