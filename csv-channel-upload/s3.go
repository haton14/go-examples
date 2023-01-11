package main

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3 interface {
	Upload(ctx context.Context, in io.Reader, fileName, conentType string) error
}

type s3Client struct {
	client *s3.Client
	bucket string
}

func NewS3(cfg aws.Config, bucket string) S3 {
	return s3Client{
		client: s3.NewFromConfig(cfg),
		bucket: bucket,
	}
}

func (s s3Client) Upload(ctx context.Context, in io.Reader, fileName, conentType string) error {
	uploader := manager.NewUploader(s.client)
	_, err := uploader.Upload(ctx,
		&s3.PutObjectInput{
			Bucket:      aws.String(s.bucket),
			Key:         aws.String(fileName),
			Body:        in,
			ContentType: aws.String(conentType),
		},
	)
	if err != nil {
		return fmt.Errorf("(manager.Uploader).Upload(): err %w", err)
	}
	return nil
}
