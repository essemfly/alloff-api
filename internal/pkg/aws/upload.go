package aws

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lessbutter/alloff-api/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetCdnURL(key string) string {
	return "https://d3n8j02f1nzq5t.cloudfront.net/" + key
}

func getProductFilePath(filename, ext string) string {
	return "products/" + filename + "." + ext
}

func UploadImage(file io.Reader, filekey, extension string) (string, error) {
	uploader := manager.NewUploader(config.S3Client)
	bucketname := viper.GetString("AWS_STORAGE_BUCKET_NAME")
	filepath := getProductFilePath(filekey, extension)
	objectInput := s3.PutObjectInput{
		Bucket: &bucketname,
		Key:    &filepath,
		Body:   file,
	}

	_, err := uploader.Upload(context.TODO(), &objectInput)

	if err != nil {
		config.Logger.Error("upload image failed", zap.Error(err))
		return "", err
	}

	return GetCdnURL(filekey + "." + extension), nil
}
