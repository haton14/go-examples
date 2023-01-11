package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/gocarina/gocsv"
)

func main() {
	datas := []DTO{
		{
			ID:   3,
			Name: "name3",
		},
		{
			ID:   44,
			Name: "name44",
		},
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(
		aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:               "http://localhost:4566",
				HostnameImmutable: true,
			}, nil
		}),
	))
	if err != nil {
		log.Fatalf("config.LoadDefaultConfig(): err %s", err)
	}
	s3 := NewS3(cfg, "local-bucket")

	pr, pw := io.Pipe()

	ch := make(chan any)
	go func() {
		defer pw.Close()
		go func() {
			defer close(ch)
			for _, v := range datas {
				ch <- v
			}
		}()
		if err := gocsv.MarshalChan(ch, gocsv.DefaultCSVWriter(pw)); err != nil {
			log.Fatalf("gocsv.MarshalChan(): err %s", err)
		}
		log.Println("csv writed")
	}()
	fileName := fmt.Sprint(time.Now().Unix())

	if err := s3.Upload(context.TODO(), pr, fileName, "text/csv"); err != nil {
		log.Fatalf("(S3).Upload(): err %s", err)
	}
	log.Println("success")
}

type DTO struct {
	ID   int    `csv:"id"`
	Name string `csv:"name"`
}
