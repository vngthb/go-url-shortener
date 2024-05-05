package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"go-url-shortener/url_shortener"
)

type Handler struct {
	service    url_shortener.IService
	serializer url_shortener.ISerializer
}

func NewHandler(service url_shortener.IService, serializer url_shortener.ISerializer) *Handler {
	return &Handler{
		service: service,
		serializer: serializer,
	}
}

func (handler *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var jap map[string]string
		json.NewDecoder(r.Body).Decode(&jap)
		entry, _ := handler.service.Shorten(jap["url"])
		jsonRespone, _ := json.Marshal(entry)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRespone)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (handler *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		path := strings.Split(r.URL.Path, "/")
		url, err := handler.service.Redirect(path[1])
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
