package mongo

import (
	"github.com/Akshit8/url-shortner/pkg/url"
	"go.mongodb.org/mongo-driver/mongo"
)

type urlRepository struct {
	client     *mongo.Client
	database   string
	collection string
}

// NewURLRepository creates instance of urlRepository
func NewURLRepository(client *mongo.Client, database string, collection string) url.Repository {
	return &urlRepository{client: client, database: database, collection: collection}
}

func (u *urlRepository) Save(url *url.URL) (*url.URL, error) {
	return nil, nil
}

func (u *urlRepository) Get(code string) (*url.URL, error) {
	return nil, nil
}

func (u *urlRepository) GetAll() ([]*url.URL, error) {
	return nil, nil
}

func (u *urlRepository) Update(code string, url *url.URL) (*url.URL, error) {
	return nil, nil
}

func (u *urlRepository) Delete(code string) error {
	return nil
}
