package api

import (
	"encoding/json"
	"errors"
	"net/http"
)

func (s *Server) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) bool {
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		s.writeError(w, http.StatusUnprocessableEntity, errors.New("nevalida peta objekto"))
		return false
	}

	return true
}

func (s *Server) writeJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) writeStatus(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func (s *Server) writeError(w http.ResponseWriter, code int, err ...error) {
	type response struct {
		Error string `json:"error"`
	}

	w.WriteHeader(code)

	if code == http.StatusInternalServerError || len(err) == 0 {
		s.writeJSON(w, response{"la peto ne sukcesis, bonvolu reprovi"})
		return
	}

	s.writeJSON(w, response{err[0].Error()})
}
