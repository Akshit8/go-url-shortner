package mongo

import (
	"context"
	"time"

	"github.com/Akshit8/url-shortner/url"
	errs "github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

// NewMongoRepository creates a new mongoRepository
func NewMongoRepository(mongoURI, database string, timeout int) (url.RedirectRepository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(timeout) * time.Second,
		database: database,
	}

	client, err := newMongoClient(mongoURI, timeout)
	if err != nil {
		return nil, errs.Wrap(err, "repository.NewMongoRepo")
	}
	repo.client = client

	return repo, nil
}

func newMongoClient(mongoURI string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (m *mongoRepository) Find(code string) (*url.Redirect, error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	redirect := &url.Redirect{}
	collection := m.client.Database(m.database).Collection("redirects")
	filter := bson.M{"code": code}
	err := collection.FindOne(ctx, filter).Decode(&redirect)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.Wrap(url.ErrRedirectNotFound, "repository.Redirect.Find")
		}
		return nil, errs.Wrap(err, "repository.Redirect.Find")
	}
	return redirect, nil
}

func (m *mongoRepository) Store(redirect *url.Redirect) error {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	collection := m.client.Database(m.database).Collection("redirects")
	_, err := collection.InsertOne(ctx, redirect)
	if err != nil {
		return errs.Wrap(err, "repository.Redirect.Store")
	}
	return nil
}
