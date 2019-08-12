package api

import (
	"net/http"

	"github.com/matryer/way"

	"github.com/martinrue/parol-api/logger"
	"github.com/martinrue/parol-api/services"
)

// Server defines the API server and its dependencies.
type Server struct {
	Development bool
	Logger      *logger.Logger
	Router      *way.Router
	Services    *services.Services
}

// Start brings up the server.
func (s *Server) Start(bind string) error {
	s.routes()
	return http.ListenAndServe(bind, s.Router)
}
