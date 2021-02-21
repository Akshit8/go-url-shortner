package mongo

import (
	"github.com/Akshit8/url-shortner/pkg/redirect"
	"go.mongodb.org/mongo-driver/mongo"
)

type redirectRepository struct {
	Client     *mongo.Client
	database   string
	collection string
}

// NewRedirectRepository creates new instance of redirectRepository
func NewRedirectRepository(client *mongo.Client, database string, collection string) redirect.Repository {
	return &redirectRepository{Client: client, database: database, collection: collection}
}

func (r *redirectRepository) Find(code string) (string, error) {
	return "", nil
}
