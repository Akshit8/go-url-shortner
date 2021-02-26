package handler

import (
	"net/http"

	"github.com/Akshit8/url-shortner/pkg/url"
)

// URLHandler defines available actions for resource url
type URLHandler interface {
	CreateURL(http.ResponseWriter, *http.Request)
	GetURLByID(http.ResponseWriter, *http.Request)
}

type urlHandler struct {
	urlService urls.Service
}

// NewURLHandler creates new instance of urlHandler
func NewURLHandler(urlService urls.Service) URLHandler {
	return &urlHandler{urlService: urlService}
}

func (u *urlHandler) CreateURL(http.ResponseWriter, *http.Request) {}

func (u *urlHandler) GetURLByID(http.ResponseWriter, *http.Request) {}
