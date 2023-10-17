package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

// var (
// 	bucket    = "bucketname"
// 	objectKey = "object/key"
// )

type Bucket struct {
	//todo

	Client
}

func NewBucket() *Bucket {
	var key, secret, token, endpoint string
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, token)),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: endpoint,
			}, nil
		})),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	client := s3.NewFromConfig(cfg)

	return &Bucket{}
}

func CreateBucket() {
	_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucket,
	})

	if err != nil {
		log.Fatalf("unable CreateBucket %s, %v", bucket, err)
	}

	f, err := os.Open("/tmp/dat")
	info, err := f.Stat()
	if err != nil {
		log.Fatalf("unable stat file,  %v", err)
	}
	// putobject
	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:        &bucket,
		Key:           &objectKey,
		Body:          f,
		ContentLength: info.Size(),
	})
	if err != nil {
		log.Fatalf("unable put object: , %v", err)
	}

	ff, err := os.Open("/tmp/dat")
	finfo, err := ff.Stat()
	if err != nil {
		log.Fatalf("unable stat file, %v", err)
	}
	// multipart upload
	uploader := manager.NewUploader(client, func(u *manager.Uploader) {
		u.PartSize = 16 * 1024 * 1024 // 16MB per part
	})

	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:        &bucket,
		Key:           &objectKey,
		Body:          ff,
		ContentLength: finfo.Size(),
	})
}

func ListBuckets() {
	listBucketOutPut, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})

	if err != nil {
		log.Fatalf("unable ListBuckets: , %v", err)
	}
	//listObjectV2
	listObjectsV2Output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucket,
	})

	if err != nil {
		log.Fatalf("unable listObjectsV2: , %v", err)
	}
	fmt.Println(listObjectsV2Output)
}

func GetHash() {
	var (
		bucket    = "bucketname"
		objectKey = "objectkey"
	)
	object, err := client.HeadObject(context.Background(), &s3.HeadObjectInput{
		Bucket: &bucket,
		Key:    &objectKey,
	})
	if err != nil {
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("ipfs cid:", object.Metadata["ipfs-hash"])
	// if synced to arweave
	fmt.Println("arweave hash:", object.Metadata["arweave-hash"])
}
