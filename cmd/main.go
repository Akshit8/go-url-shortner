package main

import (
	"fmt"
	"github.com/Akshit8/url-shortner/pkg/urls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Akshit8/url-shortner/server/graphql"

	"github.com/Akshit8/url-shortner/cmd/config"
	"github.com/Akshit8/url-shortner/pkg/repository/cassandra"

	"github.com/Akshit8/url-shortner/pkg/redirect"
	"github.com/Akshit8/url-shortner/pkg/repository/mongo"
	"github.com/Akshit8/url-shortner/pkg/repository/redis"
	"github.com/Akshit8/url-shortner/pkg/url"
	"github.com/Akshit8/url-shortner/server/rest"
)

func main() {
	appConfig, err := config.LoadConfig("cmd/config")
	if err != nil {
		log.Fatalln("error loading config: ", err)
	}

	urlRepository, redirectRepository := repoSelector(appConfig)

	urlService := urls.NewURLService(urlRepository)
	redirectService := redirect.NewRedirectService(redirectRepository)

	restServer := rest.NewRestServer(urlService, redirectService)
	graphqlServer := graphql.NewGraphqlServer(urlService)

	errs := make(chan error, 3)

	go func() {
		restAddress := getListeningAddress(appConfig.Host, appConfig.RestPort)
		log.Println("starting rest server on address:", restAddress)
		errs <- http.ListenAndServe(restAddress, restServer)
	}()

	go func() {
		graphqlAddress := getListeningAddress(appConfig.Host, appConfig.GraphqlPort)
		log.Println("starting graphql server on address:", graphqlAddress)
		errs <- http.ListenAndServe(graphqlAddress, graphqlServer)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated: %s", <-errs)
}

func getListeningAddress(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}

func repoSelector(config config.AppConfig) (urls.Repository, redirect.Repository) {
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
