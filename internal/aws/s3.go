package aws

import (
	"bytes"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Upload(path string) error {
	session, _ := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("AWS_S3_REGION"))})
	uploadedFile, err := os.Open(path)

	if err != nil {
		return err
	}

	defer uploadedFile.Close()

	uploadedFileInfo, _ := uploadedFile.Stat()
	var fileSize int64 = uploadedFileInfo.Size()
	fileBuffer := make([]byte, fileSize)
	uploadedFile.Read(fileBuffer)

	_, err = s3.New(session).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(os.Getenv("AWS_S3_BUCKET")),
		Key:                  aws.String(path),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(fileBuffer),
		ContentLength:        aws.Int64(fileSize),
		ContentType:          aws.String(http.DetectContentType(fileBuffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}
