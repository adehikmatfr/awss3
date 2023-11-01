# AWS Client With S3

This Go package provides a simple AWS client for managing AWS sessions with configurable options. It's designed to work with the AWS SDK for Go (a.k.a. "aws-sdk-go").

Package provides a flexible and easy-to-use client for interacting with Amazon S3 (Simple Storage Service). It allows you to create buckets, list buckets, upload files with customizable metadata, and download files from S3. This package is built on top of the AWS SDK for Go and provides a higher-level interface for common S3 operations.

## Features

- Easily create and manage AWS sessions with custom options.
- Configuration of AWS region, credentials, and endpoint.
- Support for setting `S3ForcePathStyle` for working with AWS S3.
- Create and manage S3 buckets.
- List available S3 buckets.
- Upload files to S3 with customizable content disposition and content type.
- Download files from S3.

## Installation

You can include this package in your Go project using the `go get` command:

```shell
go get github.com/adehikmatfr/awss3
```

## Usage

```go
import "github.com/adehikmatfr/awss3/aws"

// Initialize an AWS client with your desired configuration.
opts := awss3.Opts{
    Region:           "us-east-1",
    AccessKeyID:      "YOUR_ACCESS_KEY_ID",
    SecretKey:        "YOUR_SECRET_KEY",
    Token:            "YOUR_SECURITY_TOKEN", // Optional, set to empty string if not used.
    Endpoint:         "AWS_ENDPOINT_URL",     // Optional, set to empty string if not used.
    S3ForcePathStyle: true,                  // Set to true if needed.
}

awsClient := awss3.New(opts)

// Create a new AWS session with the configured options.
session, err := awsClient.NewSession()
if err != nil {
    // Handle the error
}

// Now you can use the AWS session for various AWS services.
```

## Configuration

- `Region`: The AWS region to use.
- `AccessKeyID`: Your AWS access key ID.
- `SecretKey`: Your AWS secret key.
- `Token`: Your AWS security token (optional, set to an empty string if not used).
- `Endpoint`: The AWS endpoint URL (optional, set to an empty string if not used).
- `S3ForcePathStyle`: Set to `true` if using S3 and you need path-style access.

```go
import "github.com/adehikmatfr/awss3/aws/s3"

// Initialize an S3 client with your AWS session.
opts := s3.S3Opts{
    Session: yourAWSSession, // Provide your configured AWS session.
}

s3Client := s3.New(opts)

// Create an S3 bucket.
bucketName := "your-bucket-name"
err := s3Client.CreateBucket(bucketName)
if err != nil {
    // Handle the error
}

// List S3 buckets.
buckets, err := s3Client.ListBuckets()
if err != nil {
    // Handle the error
}
for _, bucket := range buckets.Buckets {
    // Process each bucket
}

// Upload a file to S3.
uploadOpts := s3.NewUploadFileOpts()
uploadOpts.BucketName = bucketName
uploadOpts.ContentDisposition = "inline"
uploadOpts.FileName = "example.jpg"
uploadOpts.File = yourFileReader // Provide an io.Reader for the file content.

url, err := s3Client.UploadFile(ctx, uploadOpts)
if err != nil {
    // Handle the error
}

// Download a file from S3.
key := "path/to/your/4-download.jpg"
reader, err := s3Client.DownloadFile(bucketName, key)
if err != nil {
    // Handle the error
}

// Process the downloaded file content using the reader.
```
