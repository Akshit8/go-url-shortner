package handler

import (
	"net/http"

	"github.com/Akshit8/url-shortner/pkg/url"
)

// URLHandler defines available actions for resource url
type URLHandler interface {
	CreateURL(http.ResponseWriter, *http.Request)
	GetURLById(http.ResponseWriter, *http.Request)
	GetAllURL(http.ResponseWriter, *http.Request)
	UpdateURL(http.ResponseWriter, *http.Request)
	DeleteURL(http.ResponseWriter, *http.Request)
}

type urlHandler struct {
	urlService url.Service
}

// NewURLHandler creates new instance of urlHandler
func NewURLHandler(urlService url.Service) URLHandler {
	return &urlHandler{urlService: urlService}
}

func (u *urlHandler) CreateURL(http.ResponseWriter, *http.Request) {}

func (u *urlHandler) GetURLById(http.ResponseWriter, *http.Request) {}

func (u *urlHandler) GetAllURL(http.ResponseWriter, *http.Request) {}

func (u *urlHandler) UpdateURL(http.ResponseWriter, *http.Request) {}

func (u *urlHandler) DeleteURL(http.ResponseWriter, *http.Request) {}
