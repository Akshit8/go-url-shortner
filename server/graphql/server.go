package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Akshit8/url-shortner/pkg/url"
	"github.com/Akshit8/url-shortner/server/graphql/generated"
	"github.com/Akshit8/url-shortner/server/graphql/resolver"
)

// StartGraphqlServer starts a Graphql server on given address and service
func StartGraphqlServer(urlService url.Service, address string) {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		UrlService: urlService,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://%s/ for GraphQL playground", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
