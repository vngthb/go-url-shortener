package main

import (
	"fmt"
	"go-url-shortener/memrepo"
	"go-url-shortener/shortener"
	"hash/fnv"
	"strconv"
	"time"
)

func main() {
	svc := shortener.New(
		shortener.WithRepo(memrepo.New()),
		shortener.WithGeneratorFunc(id),
		shortener.WithTimestampFunc(now),
	)
	
	entry, _ := svc.Shorten("https://www.google.com/")
	fmt.Printf("url: " + entry.Url() + "\n")
	fmt.Printf("shortUrl: " + entry.Path() + "\n")
	dateadded := strconv.FormatInt(int64(entry.DateAdded()), 10)
	fmt.Printf("dateAdded: "+dateadded + "\n")

	s1, _ := svc.Redirect("1508094990")
	_, err2 := svc.Redirect("1508094991")

	fmt.Println("url for shortUrl: 1508094990: " + s1)
	fmt.Println("url for shortUrl: 1508094991: " + err2.Error())
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