package main

import (
	"fmt"
	"log"

	"github.com/Akshit8/url-shortner/cmd/config"
	"github.com/Akshit8/url-shortner/pkg/repository/cassandra"

	"github.com/Akshit8/url-shortner/pkg/redirect"
	"github.com/Akshit8/url-shortner/pkg/repository/mongo"
	"github.com/Akshit8/url-shortner/pkg/repository/redis"
	"github.com/Akshit8/url-shortner/pkg/url"
	"github.com/Akshit8/url-shortner/server/graphql"
	"github.com/Akshit8/url-shortner/server/rest"
)

func main() {
	appConfig, err := config.LoadConfig("cmd/config")
	if err != nil {
		log.Fatalln("error loading config: ", err)
	}
	fmt.Println(appConfig)
	fmt.Println("a")
	urlRepository, redirectRepository := repoSelector(appConfig)
	fmt.Println("b")
	urlService := url.NewURLService(urlRepository)
	redirectService := redirect.NewRedirectService(redirectRepository)

	go rest.StartRestServer(urlService, redirectService, getListeningAddress(appConfig.Host, appConfig.RestPort))

	go graphql.StartGraphqlServer(urlService, getListeningAddress(appConfig.Host, appConfig.GraphqlPort))

	fmt.Scanln()
}

func getListeningAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func repoSelector(config config.AppConfig) (url.Repository, redirect.Repository) {
	urlTable := config.URLTable
	switch config.RepoType {
	case "redis":
		client, err := redis.NewClient(config.RedisURI)
		if err != nil {
			log.Fatalln("error creating redis client: ", err)
		}
		urlRepository := redis.NewURLRepository(client)
		redirectRepository := redis.NewRedirectRepository(client)
		return urlRepository, redirectRepository

	case "mongodb":
		timeout := 10
		client, err := mongo.NewClient(config.MongoURI, timeout)
		if err != nil {
			log.Fatalln("error creating mongo client: ", err)
		}
		urlRepository := mongo.NewURLRepository(client, config.DbName, urlTable)
		redirectRepository := mongo.NewRedirectRepository(client, config.DbName, urlTable)
		return urlRepository, redirectRepository

	case "cassandra":
		session, err := cassandra.NewClient(config.CassandraHost, config.CassandraPort, urlTable)
		if err != nil {
			log.Fatalln("error creating cassandra client: ", err)
		}
		urlRepository := cassandra.NewURLRepository(session, urlTable)
		redirectRepository := cassandra.NewRedirectRepository(session, urlTable)
		return urlRepository, redirectRepository

	default:
		log.Fatalln("please select a valid repostory")
	}

	return nil, nil
}
