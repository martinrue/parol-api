package api

import (
	"net/http"
)

// Commit is the current Git SHA, injected at build-time.
var Commit = ""

func (s *Server) handleHealthz() http.HandlerFunc {
	type response struct {
		Commit string `json:"commit"`
		Passed bool   `json:"passed"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		rid := s.rid()

		s.Logger.Print(rid, "creating health report")

		s.addCORSHeaders(w, r)

		s.writeJSON(w, &response{
			Commit: Commit,
			Passed: true,
		})
	}
}
