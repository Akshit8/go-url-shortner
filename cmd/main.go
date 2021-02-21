package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Akshit8/url-shortner/pkg/redirect"
	"github.com/Akshit8/url-shortner/pkg/repository/mongo"
	"github.com/Akshit8/url-shortner/pkg/repository/redis"
	"github.com/Akshit8/url-shortner/pkg/url"
	"github.com/Akshit8/url-shortner/server/graphql"
	"github.com/Akshit8/url-shortner/server/rest"
)

func main() {
	urlRepository, redirectRepository := repoInitializer()
	urlService := url.NewURLService(urlRepository)
	redirectService := redirect.NewRedirectService(redirectRepository)

	go rest.StartRestServer(urlService, redirectService, "0.0.0.0:8080")

	go graphql.StartGraphqlServer(urlService, "0.0.0.0:8081")

	fmt.Scanln()
}

func repoInitializer() (url.Repository, redirect.Repository) {
	switch os.Getenv("DB") {
	case "redis":
		redisURI := os.Getenv("REDIS_URI")
		client, err := redis.NewClient(redisURI)
		if err != nil {
			log.Fatal(err)
		}
		urlRepository := redis.NewURLRepository(client)
		redirectRepository := redis.NewRedirectRepository(client)
		return urlRepository, redirectRepository
	case "mongo":
		mongoURI := os.Getenv("MONGO_URI")
		database := os.Getenv("DB_NAME")
		urlCollection := os.Getenv("URL_COLLECTION")
		timeout := 10
		client, err := mongo.NewClient(mongoURI, timeout)
		if err != nil {
			log.Fatal(err)
		}
		urlRepository := mongo.NewURLRepository(client, database, urlCollection)
		redirectRepository := mongo.NewRedirectRepository(client, database, urlCollection)
		return urlRepository, redirectRepository
	default:
		log.Println("please select any available database")
	}
	return nil, nil
}
