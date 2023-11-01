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
	Region      string
	AccessKeyID string
	SecretKey   string
	Token       string
	Endpoint    string
}

type Client struct {
	region      string
	accessKeyID string
	secretKey   string
	token       string
	endpoint    string
}

func New(o Opts) *Client {
	return &Client{
		region:      o.Region,
		accessKeyID: o.AccessKeyID,
		secretKey:   o.SecretKey,
		token:       o.Token,
		endpoint:    o.Endpoint,
	}
}

func (a *Client) NewSession() (*session.Session, error) {
	var (
		reg, ep *string
		cred    *credentials.Credentials
	)
	if a.region != "" {
		reg = aws.String(a.region)
	}

	if a.accessKeyID != "" {
		cred = credentials.NewStaticCredentials(
			a.accessKeyID,
			a.secretKey,
			a.token,
		)
	}

	if a.endpoint != "" {
		ep = aws.String(a.endpoint)
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      reg,
		Credentials: cred,
		Endpoint:    ep,
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}
