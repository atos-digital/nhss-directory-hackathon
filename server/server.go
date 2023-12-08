package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
	// TODO: delcare other resources the server may need
}

func New() *Server {
	router := chi.NewRouter()
	return &Server{
		router: router,
	}
}

func (s *Server) Run() error {
	s.router.Get("/example", auth(s.handleExample))
	return http.ListenAndServe(":8080", s.router)
}

func (s *Server) handleExample(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("example"))
}
