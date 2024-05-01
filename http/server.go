package http

import (
	"net/http"
)

type Server struct {
	http *http.Server
}

func NewServer(port string) *Server {
	server := http.Server{
		Addr: port,
	}

	return &Server{
		http: &server,
	}
}

func (server *Server) WithHandler(path string, fun func(http.ResponseWriter, *http.Request)) *Server {
	http.HandleFunc(path, fun)
	return server
}

func (server *Server) Start() {
	server.http.ListenAndServe()
}
