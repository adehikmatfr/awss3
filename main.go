package main

import (
	"awss3/aws"
	"awss3/aws/s3"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func exampleUpload(sClient s3.S3) {
	fileName := "4.jpg"
	filePath := filepath.Join("assets", fileName)

	// Open the file for reading.
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	url, err := sClient.UploadFile(context.Background(), &s3.UploadFileOpts{
		ContentDisposition: "inline",
		BucketName:         "career-bridge",
		FileName:           fileName,
		File:               file,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(url)
}

func exampleDownload(sClient s3.S3) {
	reader, err := sClient.DownloadFile("career-bridge", "4.jpg")
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	localFilePath := "4-download.jpg"

	file, err := os.Create(localFilePath)
	if err != nil {
		fmt.Println("Error creating local file:", err)
		return
	}
	defer file.Close()

	// Copy the content from the S3 reader to the local file.
	_, err = io.Copy(file, reader)
	if err != nil {
		fmt.Println("Error copying S3 content to local file:", err)
		return
	}
}

func main() {
	c := aws.New(aws.Opts{
		AccessKeyID:      "test",
		SecretKey:        "test",
		Token:            "",
		Region:           "ap-southeast-1",
		Endpoint:         "http://103.31.39.7:4566/",
		S3ForcePathStyle: true,
	})

	sess, err := c.NewSession()
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}

	sClient := s3.NewS3(s3.S3Opts{
		Session: sess,
	})

	exampleUpload(sClient)
	exampleDownload(sClient)
}
