package config

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/spf13/viper"
)

var S3Client *s3.Client

func InitAws() {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(viper.GetString("AWS_ACCESS_KEY_ID"), viper.GetString("AWS_SECRET_ACCESS_KEY"), "")),
		config.WithRegion(viper.GetString("AWS_S3_REGION_NAME")),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	S3Client = s3.NewFromConfig(cfg)
}
