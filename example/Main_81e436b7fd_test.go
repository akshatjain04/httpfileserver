package main

import (
	"net/http"
	"github.com/schollz/httpfileserver"
	"testing"
)

type Server struct {
	port string
}

func NewServer(port string) *Server {
	return &Server{port}
}

func (s *Server) Start() error {
	http.Handle("/new/", httpfileserver.New("/new/", "."))
	http.Handle("/", http.FileServer(http.Dir(".")))
	return http.ListenAndServe(s.port, nil)
}

func TestMain(t *testing.T) {
	server := NewServer(":1113")
	err := server.Start()
	if err != nil {
		t.Errorf("Failed to start server: %v", err)
	}
}
