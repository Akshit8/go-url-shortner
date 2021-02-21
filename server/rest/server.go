package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Akshit8/url-shortner/pkg/redirect"
	"github.com/Akshit8/url-shortner/pkg/url"
	"github.com/Akshit8/url-shortner/server/rest/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// StartRestServer starts a http server(REST) on given address and service
func StartRestServer(urlService url.Service, redirectService redirect.Service, address string) {
	urlHandler := handler.NewURLHandler(urlService)
	redirectHandler := handler.NewRedirectHandler(redirectService)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", healthHandler)

	r.Get("/{code}", redirectHandler.RedirectURL)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/url", urlHandler.CreateURL)
		r.Get("/url/{code}", urlHandler.GetURLById)
		r.Get("/url", urlHandler.GetAllURL)
		r.Put("/url/{code}", urlHandler.UpdateURL)
		r.Delete("/url/{code}", urlHandler.DeleteURL)
	})

	log.Printf("starting rest server on address: %s", address)
	log.Fatal(http.ListenAndServe(address, r))
}

type healthResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func healthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	response := &healthResponse{
		StatusCode: 200,
		Message:    "Rest service is working",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println("healthHandler.Encode: ", err)
	}
}
