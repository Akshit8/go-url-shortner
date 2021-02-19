package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Akshit8/url-shortner/api"
	"github.com/Akshit8/url-shortner/repository/mongo"
	"github.com/Akshit8/url-shortner/repository/redis"
	"github.com/Akshit8/url-shortner/url"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	repo := repoSelector()
	service := url.NewRedirectService(repo)
	handler := api.NewHandler(service)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{code}", handler.RedirectURL)
	r.Post("/", handler.CreateShortURL)

	errs := make(chan error, 2)
	go func() {
		address := httpAddress()
		fmt.Printf("server listening on %s\n", address)
		errs <- http.ListenAndServe(address, r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("interupted: %s", <-c)
	}()

	fmt.Printf("Terminated %s\n", <-errs)
}

func httpAddress() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf("0.0.0.0:%s", port)
}

func repoSelector() url.RedirectRepository {
	switch os.Getenv("DB") {
	case "redis":
		redisURI := os.Getenv("REDIS_URI")
		repo, err := redis.NewRedisRepository(redisURI)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	case "mongo":
		mongoURI := os.Getenv("MONGO_URI")
		database := os.Getenv("DB_NAME")
		timeout := 10
		repo, err := mongo.NewMongoRepository(mongoURI, database, timeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	default:
		log.Println("please select any available databse")
	}
	return nil
}
