package redis

import (
	"github.com/Akshit8/url-shortner/pkg/redirect"
	"github.com/go-redis/redis/v8"
)

type redirectRepository struct {
	Client *redis.Client
}

// NewRedirectRepository creates a new instance of redirectRepository
func NewRedirectRepository(client *redis.Client) redirect.Repository {
	return &redirectRepository{Client: client}
}

func (r *redirectRepository) Find(code string) (string, error) {
	return "", nil
}
