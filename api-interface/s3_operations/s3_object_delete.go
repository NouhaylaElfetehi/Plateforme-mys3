package s3_operations

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3ObjectDeleter struct {
	client *s3.Client
}

func NewS3ObjectDeleter(client *s3.Client) *S3ObjectDeleter {
	return &S3ObjectDeleter{
		client: client,
	}
}

func (deleter *S3ObjectDeleter) DeleteObject(ctx context.Context, bucketName string, objectKey string) error {
	_, err := deleter.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	if err != nil {
		log.Printf("Error deleting object %s from bucket %s: %v", objectKey, bucketName, err)
		return err
	}
	return nil
}
