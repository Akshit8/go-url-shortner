package handler

import (
	"net/http"

	"github.com/Akshit8/url-shortner/pkg/redirect"
)

// RedirectHandler defines available methods on resource redirect
type RedirectHandler interface {
	RedirectURL(http.ResponseWriter, *http.Request)
}

type redirectHandler struct {
	redirectService redirect.Service
}

// NewRedirectHandler creates new instance of redirectHandler
func NewRedirectHandler(redirectService redirect.Service) RedirectHandler {
	return &redirectHandler{redirectService: redirectService}
}

func (r *redirectHandler) RedirectURL(http.ResponseWriter, *http.Request) {}
