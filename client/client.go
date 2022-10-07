package client

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ClientInterface interface {
	CreateSession() (*session.Session, error)
	CreateBucket(client *s3.S3, bucketName string) error
	ListBuckets(client *s3.S3) (*s3.ListBucketsOutput, error)
	ListObjects(client *s3.S3, bucketName string, prefix string) (*s3.ListObjectsV2Output, error)
}

type Client struct {
	Profile string
	Region  string
}

func NewClient(profile string, region string) ClientInterface {
	return &Client{Profile: profile, Region: region}
}

func (cl *Client) CreateSession() (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: cl.Profile,
		Config: aws.Config{
			Region: aws.String(cl.Region),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return nil, err
	}
	return sess, nil
}

func (cl *Client) CreateBucket(client *s3.S3, bucketName string) error {
	_, err := client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	return err
}

func (cl *Client) ListBuckets(client *s3.S3) (*s3.ListBucketsOutput, error) {
	res, err := client.ListBuckets(nil)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (cl *Client) ListObjects(client *s3.S3, bucketName string, prefix string) (*s3.ListObjectsV2Output, error) {
	res, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, err
	}
	
	return res, nil
}
