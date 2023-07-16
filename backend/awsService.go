package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"os"
)

//type awsService struct {
//	awsClient *awsClient
//}

var BUCKET_NAME = "ignite-go-blog-server-bucket"

var BUCKTET_FILE_PATH = "https://ignite-go-blog-server-bucket.s3.amazonaws.com/"

type awsS3Client struct {
	S3Client *s3.Client
}

var AwsS3client awsS3Client

func InitAws() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("ignite"), config.WithRegion("us-east-1"))
	if err != nil {
		log.Printf("Couldn't load config. Here's why: %v\n", err)
	}
	// Create an Amazon S3 service client
	c := s3.NewFromConfig(cfg)

	AwsS3client = awsS3Client{
		S3Client: c,
	}

	//path, err := client.UploadFile(BUCKET_NAME, "cat.png", "/Users/aaronjanes/GolandProjects/goSever/backend/images/cat.png")
	//if err != nil {
	//	log.Printf("Couldn't upload file. Here's why: %v\n", err)
	//}
	//
	//fmt.Printf("Uploaded file to %v\n", path)

}

func (c awsS3Client) UploadFile(bucketName string, objectKey string, fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Couldn't open file %v to upload. Here's why: %v\n", fileName, err)
	} else {
		defer file.Close()
		_, err := c.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(objectKey),
			Body:   file,
		})

		if err != nil {
			log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
				fileName, bucketName, objectKey, err)
		}

		return BUCKTET_FILE_PATH + objectKey, err
	}
	return "", err
}
