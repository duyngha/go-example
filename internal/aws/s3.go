package aws

import (
	"context"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"example.com/m/internal/helpers"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context, path string, key string) (pathFile string, err error) {
	client := uploader()

	file, fileHeader, err := c.Request.FormFile(key)
	if err != nil {
		return "can not read the file.", err
	}

	bucket := helpers.Env("AWS_BUCKET")

	fileType, err := getFileType(file)
	if err != nil {
		return "can not get the file extension", err
	}

	pathFile = strings.Trim(path, "/") + "/" + fileHeader.Filename

	uploader := manager.NewUploader(client)
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      &bucket,
		Key:         &pathFile,
		Body:        file,
		ContentType: &fileType,
	})

	return
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

// func getFileURL(output *manager.UploadOutput) (url string) {
// 	return strings.Replace(output.Location, "https://s3."+helpers.Env("AWS_DEFAULT_REGION")+".amazonaws.com/"+helpers.Env("AWS_BUCKET")+"/", helpers.Env("AWS_CDN_URL")+"/", -1)
// }

func GetFileURLWithCDN(path string) string {
	return helpers.Env("AWS_CDN_URL") + "/" + path
}
