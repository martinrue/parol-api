package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/martinrue/vocx"
)

const (
	maxTextLength      = 300
	maxTextLengthExtra = 3000
)

var (
	errMissingText      = errors.New("mankas teksto")
	errTextTooLong      = fmt.Errorf("la teksto devas esti %d literoj aŭ malpli", maxTextLength)
	errInvalidVoice     = errors.New("mankas voĉo (kiu devus esti aŭ 'male' aŭ 'female')")
	errMissingRules     = errors.New("mankas reguloj")
	errInvalidRules     = errors.New("la reguloj devas esti valida JSON-a dosiero")
	errGlobalUsageLimit = errors.New("la ĉiuhora uzadlimo estis atingita – reprovu baldaŭ")
)

type speakRequest struct {
	Text  string `json:"text"`
	Voice string `json:"voice"`
	Rules string `json:"rules"`
	Key   string `json:"key"`
}

func (r *speakRequest) validate(hasOverride bool) error {
	if strings.TrimSpace(r.Text) == "" {
		return errMissingText
	}

	textLimit := maxTextLength

	if hasOverride {
		textLimit = maxTextLengthExtra
	}

	if len(strings.TrimSpace(r.Text)) > textLimit {
		return errTextTooLong
	}

	voice := strings.TrimSpace(r.Voice)

	if voice != "male" && voice != "female" {
		return errInvalidVoice
	}

	if strings.TrimSpace(r.Rules) == "" {
		return errMissingRules
	}

	return nil
}

type speakResponse struct {
	URL         string `json:"url"`
	Transcribed string `json:"transcribed"`
}

func (s *Server) handleSpeak() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rid := s.rid()

		s.addCORSHeaders(w, r)

		var req speakRequest
		if ok := s.readJSON(w, r, &req); !ok {
			return
		}

		s.Services.Logger.Print(rid, "handling speak request (%d chars)", len(req.Text))

		hasOverride := s.Services.Config.ValidKey(req.Key)

		if err := req.validate(hasOverride); err != nil {
			s.Services.Logger.Print(rid, "invalid request → %s", err)
			s.writeError(w, http.StatusUnprocessableEntity, err)
			return
		}

		maxRequests := s.Services.Config.MaxRequests

		if maxRequests != 0 && s.Services.Usage.LimitExceeded(maxRequests) {
			if hasOverride {
				s.Services.Logger.Print(rid, "ignoring hourly limit, key: %s", req.Key)
			} else {
				s.Services.Logger.Print(rid, "global hourly limit reached, request rejected")
				s.writeError(w, http.StatusTooManyRequests, errGlobalUsageLimit)
				return
			}
		}

		transcriber := vocx.NewTranscriber()

		if err := transcriber.LoadRules(req.Rules); err != nil {
			s.Services.Logger.Print(rid, "failed to load rules")
			s.writeError(w, http.StatusUnprocessableEntity, errInvalidRules)
			return
		}

		transcribed := transcriber.Transcribe(req.Text)

		reader, contentType, err := s.Services.SynthesiseSpeech(transcribed, req.Voice)
		if err != nil {
			s.Services.Logger.Print(rid, "failed to synthesise speech → %s", err)
			s.writeError(w, http.StatusInternalServerError)
			return
		}

		defer reader.Close()

		location, err := s.Services.UploadData(reader, contentType)
		if err != nil {
			s.Services.Logger.Print(rid, "failed to upload data → %s", err)
			s.writeError(w, http.StatusInternalServerError)
			return
		}

		s.Services.Usage.TrackRequest()

		s.writeJSON(w, &speakResponse{location, transcribed})
	}
}
