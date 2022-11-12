package aws

import (
	"context"
	"log"
	"mime/multipart"
	"net/http"

	"example.com/m/internal/helpers"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) (err error) {
	client := uploader()

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		return err
	}

	bucket := helpers.Env("AWS_BUCKET")

	fileType, err := getFileType(file)
	if err != nil {
		return err
	}

	_, err = client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &fileHeader.Filename,
		Body:        file,
		ContentType: &fileType,
	})

	//TODO: Retrieve the object URL from the bucket
	return err
}

type AwsCfg struct {
	AccessKeyID     string
	SecretAccessKey string
}

// https://aws.github.io/aws-sdk-go-v2/docs/configuring-sdk/
func uploader() *s3.Client {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				helpers.Env("AWS_ACCESS_KEY_ID"),
				helpers.Env("AWS_SECRET_ACCESS_KEY"),
				""),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	return client
}

// https://freshman.tech/file-upload-golang/
func getFileType(file multipart.File) (fileType string, err error) {
	buff := make([]byte, 512)
	_, err = file.Read(buff)

	fileType = http.DetectContentType(buff)
	return fileType, err
}
