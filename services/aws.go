package services

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func (s *Services) createSession() *session.Session {
	return session.New(&aws.Config{
		Region:      aws.String(s.Config.AWSRegion),
		Credentials: credentials.NewStaticCredentials(s.Config.AWSKey, s.Config.AWSSecret, ""),
	})
}
