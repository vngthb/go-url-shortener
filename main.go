package main

import (
	"encoding/json"
	"hash/fnv"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go-url-shortener/memrepo"
	"go-url-shortener/shortener"
)

type HttpHandler struct {
	Service shortener.ShortenerService
}

func main() {
	svc := shortener.New(
		shortener.WithRepo(memrepo.New()),
		shortener.WithGeneratorFunc(id),
		shortener.WithTimestampFunc(now),
	)

	handler := &HttpHandler{
		Service: svc,
	}

	srv := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/shorten", handler.shorten)
	http.HandleFunc("/", handler.redirect)

	srv.ListenAndServe()
}

func (handler *HttpHandler) shorten(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var jap map[string]string
		json.NewDecoder(r.Body).Decode(&jap)
		entry, _ := handler.Service.Shorten(jap["url"])
		jsonRespone, _ := json.Marshal(entry)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonRespone)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (handler *HttpHandler) redirect(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		path := strings.Split(r.URL.Path, "/")
		url, err := handler.Service.Redirect(path[1])
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

func id(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	a := h.Sum32()
	return strconv.FormatInt(int64(a), 10)
}

func now() int64 {
	return time.Now().Unix()
}

// curl --header "Content-Type: application/json" --request POST -data '{\"url\":\"https\:\/\/google\.com\"}' http://localhost:8080/
