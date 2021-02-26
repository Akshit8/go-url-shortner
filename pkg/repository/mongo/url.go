package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Akshit8/url-shortner/pkg/urls"
	"go.mongodb.org/mongo-driver/mongo"
)

type URL struct {
	ID        primitive.ObjectID `bson:"_id"`
	Code      string             `bson:"code"`
	URL       string             `bson:"url"`
	ShortURL  string             `bson:"shortUrl"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type urlRepository struct {
	client     *mongo.Client
	database   string
	collection string
	timeout    time.Duration
}

// NewURLRepository creates instance of urlRepository
func NewURLRepository(client *mongo.Client, database string, collection string, timeout int) urls.Repository {
	timeoutDuration := time.Duration(timeout) * time.Second
	return &urlRepository{
		client:     client,
		database:   database,
		collection: collection,
		timeout:    timeoutDuration,
	}
}

func (u *urlRepository) Save(url *urls.URL) (*urls.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	newURL := URL{
		ID:        primitive.NewObjectID(),
		Code:      url.Code,
		URL:       url.URL,
		ShortURL:  url.ShortURL,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	collection := u.client.Database(u.database).Collection(u.collection)
	_, err := collection.InsertOne(ctx, newURL)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (u *urlRepository) Get(code string) (*urls.URL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	url := URL{}

	collection := u.client.Database(u.database).Collection(u.collection)
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&url)
	if err != nil {
		return nil, err
	}

	result := urls.URL{
		Code:     url.Code,
		URL:      url.URL,
		ShortURL: url.ShortURL,
	}

	return &result, nil
}
