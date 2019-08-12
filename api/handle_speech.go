package api

import (
	"errors"
	"net/http"
	"strings"
)

type speechRequest struct {
	Config string `json:"config"`
	Text   string `json:"text"`
	Voice  string `json:"voice"`
}

func (r *speechRequest) validate() error {
	if strings.TrimSpace(r.Config) == "" {
		return errors.New("missing config")
	}

	if strings.TrimSpace(r.Text) == "" {
		return errors.New("missing text")
	}

	voice := strings.TrimSpace(r.Voice)

	if voice != "male" && voice != "female" {
		return errors.New("voice must be male or female")
	}

	return nil
}

type speechResponse struct {
	URL string `json:"url"`
}

func (s *Server) handleSpeech() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rid := s.rid()

		s.addCORSHeaders(w, r)

		var data speechRequest
		if ok := s.readJSON(w, r, &data); !ok {
			return
		}

		s.Logger.Print(rid, "handling speech request")

		if err := data.validate(); err != nil {
			s.Logger.Print(rid, "invalid request → %s", err)
			s.writeError(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.writeJSON(w, &speechResponse{"/sample.mp3"})
	}
}
