package api

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Akshit8/url-shortner/serializer/json"
	"github.com/Akshit8/url-shortner/serializer/msgpack"
	"github.com/Akshit8/url-shortner/url"
	"github.com/go-chi/chi"
	errs "github.com/pkg/errors"
)

// URLHandler defines handlers for resource url
type URLHandler interface {
	RedirectURL(http.ResponseWriter, *http.Request)
	CreateShortURL(http.ResponseWriter, *http.Request)
}

type handler struct {
	redirectService url.RedirectService
}

// NewHandler creates instance of handler
func NewHandler(redirectService url.RedirectService) URLHandler {
	return &handler{redirectService: redirectService}
}

func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

func (h *handler) serializer(contentType string) url.RedirectSerializer {
	if contentType == "application/x-msgpack" {
		return &msgpack.Redirect{}
	}
	return &json.Redirect{}
}

func (h *handler) RedirectURL(w http.ResponseWriter, req *http.Request) {
	code := chi.URLParam(req, "code")
	redirect, err := h.redirectService.Find(code)
	if err != nil {
		if errs.Cause(err) == url.ErrRedirectNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, req, redirect.URL, http.StatusMovedPermanently)
}

func (h *handler) CreateShortURL(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	redirect, err := h.serializer(contentType).Decode(requestBody)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = h.redirectService.Store(redirect)
	if err != nil {
		if errs.Cause(err) == url.ErrRedirectInvalid {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, err := h.serializer(contentType).Encode(redirect)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	setupResponse(w, contentType, responseBody, http.StatusCreated)
}
