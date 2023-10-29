package helpers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func ConnectAWS() (*s3.S3, error) {

    sess := session.Must(session.NewSession(&aws.Config{
        Region: aws.String(GetConfig("AWS_REGION")),
        Credentials: credentials.NewStaticCredentials(GetConfig("AWS_ACCESS_KEY"), GetConfig("AWS_SECRET_ACCESS_KEY"), ""),
    }))
    return s3.New(sess), nil
}
