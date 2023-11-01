package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AWS interface {
	NewSession() (*session.Session, error)
}

type Opts struct {
	Region           string
	AccessKeyID      string
	SecretKey        string
	Token            string
	Endpoint         string
	S3ForcePathStyle bool
}

type Client struct {
	region           string
	accessKeyID      string
	secretKey        string
	token            string
	endpoint         string
	s3ForcePathStyle bool
}

func New(o Opts) *Client {
	return &Client{
		region:           o.Region,
		accessKeyID:      o.AccessKeyID,
		secretKey:        o.SecretKey,
		token:            o.Token,
		endpoint:         o.Endpoint,
		s3ForcePathStyle: o.S3ForcePathStyle,
	}
}

func (a *Client) NewSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(a.region),
		Credentials: credentials.NewStaticCredentials(
			a.accessKeyID,
			a.secretKey,
			a.token,
		),
		Endpoint:         aws.String(a.endpoint),
		S3ForcePathStyle: aws.Bool(a.s3ForcePathStyle),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}
