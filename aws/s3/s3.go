package s3

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 interface {
	CreateBucket(bucketName string) error
	ListBuckets() (*s3.ListBucketsOutput, error)
	UploadFile(ctx context.Context, o *UploadFileOpts) (*s3manager.UploadOutput, error)
	DownloadFile(bucketName string, key string) (io.Reader, error)
}

type S3Opts struct {
	Session *session.Session
}

type S3Cilent struct {
	svc      *s3.S3
	uploader *s3manager.Uploader
}

type UploadFileOpts struct {
	BucketName string
	FileName   string
	File       io.Reader
}

func NewUploadFileOpts() *UploadFileOpts {
	return &UploadFileOpts{}
}

func NewS3(o S3Opts) S3 {
	svc := s3.New(o.Session)
	uploader := s3manager.NewUploader(o.Session)

	return &S3Cilent{
		svc:      svc,
		uploader: uploader,
	}
}

func (s *S3Cilent) CreateBucket(bucketName string) error {
	_, err := s.svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	return err
}

func (s *S3Cilent) ListBuckets() (*s3.ListBucketsOutput, error) {
	res, err := s.svc.ListBuckets(nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *S3Cilent) UploadFile(ctx context.Context, o *UploadFileOpts) (*s3manager.UploadOutput, error) {
	return s.uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String(o.BucketName),
		Key:    aws.String(o.FileName),
		Body:   o.File,
	})
}

func (s *S3Cilent) DownloadFile(bucketName string, key string) (io.Reader, error) {
	// Get the file data from S3
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	}

	result, err := s.svc.GetObject(input)
	if err != nil {
		return nil, err
	}

	return result.Body, err
}
