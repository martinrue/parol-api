package services

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/polly"
)

// SynthesiseSpeech synthesises speech via AWS Polly.
func (s *Services) SynthesiseSpeech(text string, voice string) (io.ReadCloser, string, error) {
	voiceID := "Ewa"

	if voice == "male" {
		voiceID = "Jacek"
	}

	svc := polly.New(s.createSession())

	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String("mp3"),
		SampleRate:   aws.String("22050"),
		VoiceId:      aws.String(voiceID),
		TextType:     aws.String("text"),
		Text:         aws.String(text),
	}

	result, err := svc.SynthesizeSpeech(input)
	if err != nil {
		return nil, "", err
	}

	return result.AudioStream, *result.ContentType, nil
}
