package aws

import (
	"context"
	"log"
	"mime/multipart"
	"net/http"

	"example.com/m/internal/helpers"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) (err error) {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	//TODO: Config S3 configuration
	client := s3.NewFromConfig(cfg)

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

	return err
}

//func uploader() *s3manager.Uploader {
// s3Config := &aws.Config{
// 	Region:      aws.String(os.Getenv("AWS_DEFAULT_REGION")),
// 	Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
// }

// s3Session := session.New(s3Config)

// uploader := s3manager.NewUploader(s3Session)
// return uploader
//}

func getFileType(file multipart.File) (fileType string, err error) {
	buff := make([]byte, 512)
	_, err = file.Read(buff)

	fileType = http.DetectContentType(buff)
	return fileType, err
}
