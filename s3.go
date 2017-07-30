package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func initS3() *s3.S3 {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)
	return svc
}

func createRedirObject(c *s3.S3, URL string, S3Bucket string, S3Key string) (*RedirObject, error) {
	input := &s3.PutObjectInput{
		Bucket: aws.String(S3Bucket),
		Key:    aws.String(S3Key + "/index.html"),
		WebsiteRedirectLocation: aws.String(URL),
		ContentType:             aws.String("text/html"),
	}
	if _, err := c.PutObject(input); err != nil {
		return nil, err
	}
	redir := RedirObject{
		ShortURL: "https://" + S3Bucket + "/" + S3Key,
	}
	return &redir, nil
}
