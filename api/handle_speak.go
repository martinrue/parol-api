package api

import (
	"errors"
	"net/http"
	"strings"
	"time"
)

type speakRequest struct {
	Text   string `json:"text"`
	Voice  string `json:"voice"`
	Config string `json:"config"`
}

func (r *speakRequest) validate() error {
	if strings.TrimSpace(r.Text) == "" {
		return errors.New("mankas teksto")
	}

	voice := strings.TrimSpace(r.Voice)

	if voice != "male" && voice != "female" {
		return errors.New("mankas voĉo (kiu devus esti aŭ 'male' aŭ 'female')")
	}

	if strings.TrimSpace(r.Config) == "" {
		return errors.New("mankas agordo")
	}

	return nil
}

type speakResponse struct {
	URL string `json:"url"`
}

func (s *Server) handleSpeak() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rid := s.rid()

		s.addCORSHeaders(w, r)

		var data speakRequest
		if ok := s.readJSON(w, r, &data); !ok {
			return
		}

		time.Sleep(1 * time.Second)

		s.Logger.Print(rid, "handling speak request (%d chars)", len(data.Text))

		if err := data.validate(); err != nil {
			s.Logger.Print(rid, "invalid request → %s", err)
			s.writeError(w, http.StatusUnprocessableEntity, err)
			return
		}

		s.writeJSON(w, &speakResponse{"/sample.mp3"})
	}
}
