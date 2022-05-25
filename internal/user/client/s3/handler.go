package s3

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"mime/multipart"
)

type AWSConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	BucketName      string
	UploadTimeout   int
	BaseURL         string
}

type Handler struct {
	AWSConfig *AWSConfig
	sess      *session.Session
	s3Sess    *s3.S3
}

func MakeHandler(config *AWSConfig) *Handler {
	handler := &Handler{
		AWSConfig: config,
	}
	handler.CreateSession()
	handler.CreateS3Session()

	return handler
}

func (a *Handler) CreateSession() {
	a.sess = session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(a.AWSConfig.Region),
			Credentials: credentials.NewStaticCredentials(
				a.AWSConfig.AccessKeyID,
				a.AWSConfig.AccessKeySecret,
				"",
			),
		},
	))
}

func (a *Handler) CreateS3Session() {
	a.s3Sess = s3.New(a.sess)
}

func (a *Handler) UploadObject(file multipart.File, bucket string, fileName string) error {
	uploader := s3manager.NewUploader(a.sess)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		return errors.New("error while upload file to s3")
	}

	return nil
}
