package api

import (
	"net/http"

	"github.com/martinrue/parol-api/token"
)

func (s *Server) rid() string {
	return token.NewShort()
}

func (s *Server) addCORSHeaders(w http.ResponseWriter, r *http.Request) {
	origin := "roboto.martinrue.com"

	if s.Development == true {
		origin = "http://localhost:1234"
	}

	w.Header().Add("Access-Control-Allow-Origin", origin)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
}
