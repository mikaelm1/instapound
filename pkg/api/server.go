package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	router *chi.Mux
}

func NewServer(
	router *chi.Mux) *Server {
	return &Server{
		router: router,
	}
}

func (s *Server) Run() error {
	s.Routes()
	fmt.Println("Starting server on port 8080")

	return http.ListenAndServe(":8080", s.router)
}

func (s *Server) Routes() {
	router := s.router

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to InstaPound!"))
	})

	v1ApiRouter := chi.NewRouter()
	router.Mount("/api", v1ApiRouter)

	v1ApiRouter.Group(func(r chi.Router) {
		r.Route("/auth", func(r chi.Router) {
			r.Post("/login", s.loginHandler)
			r.Post("/register", s.registerHandler)
		})
	})
}
