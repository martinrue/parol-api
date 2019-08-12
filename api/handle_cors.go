package api

import (
	"net/http"
)

func (s *Server) handleCORS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.addCORSHeaders(w, r)
		w.WriteHeader(http.StatusOK)
	}
}
