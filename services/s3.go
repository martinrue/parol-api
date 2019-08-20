package services

import (
	"io"

	"github.com/martinrue/parol-api/token"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// UploadData uploads data to the configured bucket and returns the location.
func (s *Services) UploadData(reader io.ReadCloser, contentType string) (string, error) {
	svc := s3manager.NewUploader(s.createSession())

	input := &s3manager.UploadInput{
		Bucket:      aws.String(s.Config.AWSBucket),
		Key:         aws.String(token.New()),
		ContentType: aws.String(contentType),
		Body:        reader,
	}

	result, err := svc.Upload(input)
	if err != nil {
		return "", err
	}

	return result.Location, nil
}
