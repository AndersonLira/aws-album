package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/andersonlira/album/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const ALBUM_BUCKET = "ALBUM_BUCKET"

type AwsS3 struct {
	svc         *s3.S3
	albumBucket string
	initialized bool
}

func (awsS3 *AwsS3) Init() {
	awsS3.albumBucket = os.Getenv(ALBUM_BUCKET)
	if len(awsS3.albumBucket) == 0 {
		errorMsg := fmt.Sprintf("%s env variable not declared", ALBUM_BUCKET)
		panic(errorMsg)
	}

	sess := session.Must(session.NewSession())

	awsS3.svc = s3.New(sess)
	awsS3.initialized = true
}

func (awsS3 *AwsS3) List() []model.File {
	list := []model.File{}
	if !awsS3.initialized {
		fmt.Print(list)
		panic("AWS S3 not initialized")
	}
	i := 0
	err := awsS3.svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket: &awsS3.albumBucket,
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		fmt.Println("Page,", i)
		i++

		for _, obj := range p.Contents {
			list = AddByKey(*obj.Key)
		}
		return true
	})
	if err != nil {
		log.Println("failed to list objects", err)
	}
	return list
}

func (awsS3 *AwsS3) GetPreSignedUrls(keys []string) []string {
	urls := []string{}

	for _, key := range keys {
		req, _ := awsS3.svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: &awsS3.albumBucket,
			Key:    aws.String(key),
		})
		urlStr, err := req.Presign(15 * time.Minute)

		if err != nil {
			log.Println("Failed to sign request", err)
		}
		urls = append(urls, urlStr)
	}

	return urls
}
