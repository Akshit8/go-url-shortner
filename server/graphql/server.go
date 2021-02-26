package graphql

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Akshit8/url-shortner/pkg/url"
	"github.com/Akshit8/url-shortner/server/graphql/generated"
	"github.com/Akshit8/url-shortner/server/graphql/resolver"
)

// StartGraphqlServer starts a Graphql server on given address and service
func NewGraphqlServer(urlService urls.Service) *http.ServeMux {

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		UrlService: urlService,
	}}))

	r := http.NewServeMux()

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	return r
}
