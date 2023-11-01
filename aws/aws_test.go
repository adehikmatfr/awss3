package aws

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Mock implementation of session.New for testing purposes.
type mockSession struct {
	cfg *aws.Config
}

func (m *mockSession) New(*aws.Config) (*session.Session, error) {
	return &session.Session{Config: m.cfg}, nil
}

var (
	region           = "ap-southeast-1"
	accessKeyID      = "test-access-key"
	secretKey        = "test-secret-key"
	token            = "test-token"
	endpoint         = "test-endpoint"
	s3ForcePathStyle = true
)

func TestNewClient(t *testing.T) {
	opts := Opts{
		Region:           region,
		AccessKeyID:      accessKeyID,
		SecretKey:        secretKey,
		Token:            token,
		Endpoint:         endpoint,
		S3ForcePathStyle: s3ForcePathStyle,
	}

	client := New(opts)

	if client.region != region {
		t.Error("Expected region to be 'us-east-1'")
	}
	if client.accessKeyID != "test-access-key" {
		t.Error("Expected accessKeyID to be 'test-access-key'")
	}
	if client.secretKey != "test-secret-key" {
		t.Error("Expected secretKey to be 'test-secret-key'")
	}
	if client.token != "test-token" {
		t.Error("Expected token to be 'test-token'")
	}
	if client.endpoint != "test-endpoint" {
		t.Error("Expected endpoint to be 'test-endpoint'")
	}
	if !client.s3ForcePathStyle {
		t.Error("Expected s3ForcePathStyle to be true")
	}
}

func TestNewSession(t *testing.T) {
	opts := Opts{
		Region:           region,
		AccessKeyID:      accessKeyID,
		SecretKey:        secretKey,
		Token:            token,
		Endpoint:         endpoint,
		S3ForcePathStyle: s3ForcePathStyle,
	}

	client := New(opts)

	session, err := client.NewSession()

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if session == nil {
		t.Error("Expected session to be non-nil")
	}
}
