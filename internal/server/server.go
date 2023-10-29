package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	server *http.Server
}

// constructor para retornar un servidor configurado
func New(port string) (*Server, error) {
	// inicialiamos chi para route
	r := chi.NewRouter()

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

func (serv *Server) Close() error {
	// TODO: finish the method
	return nil
}

func (serv *Server) Start() {
	log.Printf("Server running on localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
