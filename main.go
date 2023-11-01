package main

import (
	"awss3/aws"
	"awss3/aws/s3"
	"fmt"
)

func main() {
	c := aws.New(aws.Opts{
		AccessKeyID: "test",
		SecretKey:   "test",
		Token:       "",
		Region:      "ap-southeast-1",
		Endpoint:    "http://103.31.39.7:4566",
	})

	sess, err := c.NewSession()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	sClient := s3.NewS3(s3.S3Opts{
		Session: sess,
	})

	list, err := sClient.ListBuckets()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	for _, bucket := range list.Buckets {
		fmt.Printf("Found bucket: %s, created at: %s\n", *bucket.Name, *bucket.CreationDate)
	}
}
