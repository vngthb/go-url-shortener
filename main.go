package main

import (
	"hash/fnv"
	"strconv"
	"time"

	"go-url-shortener/http"
	"go-url-shortener/mem_repo"
	"go-url-shortener/url_shortener"
)

func main() {

	genFunc := func(link string) string {
		h := fnv.New32a()
		h.Write([]byte(link))
		a := h.Sum32()
		return strconv.FormatInt(int64(a), 10)
	}

	timeFunc := func() int64 {
		return time.Now().Unix()
	}

	service := url_shortener.NewService(
		url_shortener.WithRepo(mem_repo.NewRepo()),
		url_shortener.WithGeneratorFunc(genFunc),
		url_shortener.WithTimestampFunc(timeFunc),
	)

	handler := http.NewHandler(service, nil)

	http.NewServer(":8080").
		WithHandler("/shorten", handler.Shorten).
		WithHandler("/", handler.Redirect).
		Start()
}
