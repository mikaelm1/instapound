package api

import "net/http"

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

func (s *Server) registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Register"))
}
