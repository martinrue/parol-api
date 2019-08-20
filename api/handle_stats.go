package api

import (
	"net/http"
	"time"
)

// Commit is the current Git SHA, injected at build-time.
var Commit string

func (s *Server) handleStats() http.HandlerFunc {
	type response struct {
		Uptime        string         `json:"uptime"`
		Commit        string         `json:"commit"`
		TotalRequests int            `json:"totalRequests"`
		HourlyUsage   map[string]int `json:"hourlyUsage"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.addCORSHeaders(w, r)

		s.writeJSON(w, &response{
			Uptime:        time.Since(s.Services.Usage.Started).String(),
			Commit:        Commit,
			TotalRequests: s.Services.Usage.TotalRequests,
			HourlyUsage:   s.Services.Usage.HourlyUsage,
		})
	}
}
