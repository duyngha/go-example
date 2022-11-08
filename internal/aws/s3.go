package aws

func Upload() {
	// input := &s3manager.UploadInput{
	// 	Bucket:      aws.String(os.Getenv("AWS_BUCKET")),
	// 	Key:         aws.String(path),
	// 	Body:        file,
	// 	ContentType: aws.String("image/jpg"),
	// }

	// output, err := uploader().UploadWithContext(context.Background(), input)

	// log.Printf("res %v\n", output)
	// log.Printf("err %v\n", err)
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
